package migrate

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// データの差分有無に関わらず、マイグレート処理を行う。
	MigrateForce = "force"

	// データの差分有無に従い、マイグレート処理を実施するか判断する。
	MigrateDefault = "default"
)

type Migrator struct {
	Marshal func(interface{}) (map[string]interface{}, error)

	Unmarshal func(map[string]interface{}) (interface{}, error)

	Migrate func(context.Context, DataContainer) (DataContainer, string, error)
}

type ShowRequest struct {
	PlanFile string
}

func (v *ShowRequest) IsValid() bool {
	for _, s := range []string{
		v.PlanFile,
	} {
		if s == "" {
			return false
		}
	}
	return true
}

func (m *Migrator) Show(ctx context.Context, req ShowRequest) error {
	if !req.IsValid() {
		return errors.New("invalid argument")
	}

	rc, err := os.Open(req.PlanFile)
	if err != nil {
		return err
	}
	defer rc.Close()

	jd := json.NewDecoder(rc)
	jd.DisallowUnknownFields()
	jd.UseNumber()

	for jd.More() {
		var item planItem
		if err := jd.Decode(&item); err != nil {
			return err
		}

		differ := diffopts[item.DiffOption](item.From, item.To)
		if differ == "" {
			continue
		}
		fmt.Fprintln(os.Stdout, differ)
	}
	return nil
}

type PlanRequest struct {
	ProjectID string

	CollectionPath string

	PlanFile string
}

func (v *PlanRequest) IsValid() bool {
	for _, s := range []string{
		v.ProjectID,
		v.CollectionPath,
		v.PlanFile,
	} {
		if s == "" {
			return false
		}
	}
	return true
}

func (m *Migrator) Plan(ctx context.Context, req PlanRequest) error {
	if !req.IsValid() {
		return errors.New("invalid argument")
	}

	wc, err := os.Create(req.PlanFile)
	if err != nil {
		return err
	}
	defer wc.Close()

	// jsonlinesで出力する
	je := json.NewEncoder(wc)
	je.SetIndent("", "")

	cli, err := firestore.NewClient(ctx, req.ProjectID, option.WithGRPCDialOption(grpc.WithBlock()))
	if err != nil {
		return err
	}
	defer cli.Close()

	collection := cli.Collection(req.CollectionPath)
	snapshots, err := collection.Documents(ctx).GetAll()
	if err != nil {
		return err
	}
	for _, fromSnap := range snapshots {
		fromValue, err := m.Unmarshal(fromSnap.Data())
		if err != nil {
			return err
		}
		from := DataContainer{
			CollectionPath: fromSnap.Ref.Parent.ID,
			DocumentID:     fromSnap.Ref.ID,
			Value:          fromValue,
			CreateTime:     fromSnap.CreateTime,
			UpdateTime:     fromSnap.UpdateTime,
		}
		to, diffopt, err := m.Migrate(ctx, from.Clone())
		if err != nil {
			return err
		}

		differ := diffopts[diffopt](from, to)
		if differ == "" {
			continue
		}
		fmt.Fprintln(os.Stdout, differ)

		item := &planItem{
			From:       from,
			To:         to,
			DiffOption: diffopt,
		}
		if err := je.Encode(item); err != nil {
			return err
		}
	}

	return nil
}

type ApplyRequest struct {
	ProjectID string

	PlanFile string
}

func (v *ApplyRequest) IsValid() bool {
	for _, s := range []string{
		v.ProjectID,
		v.PlanFile,
	} {
		if s == "" {
			return false
		}
	}
	return true
}

func (m *Migrator) Apply(ctx context.Context, req ApplyRequest) error {
	if !req.IsValid() {
		return errors.New("invalid argument")
	}

	rc, err := os.Open(req.PlanFile)
	if err != nil {
		return err
	}
	defer rc.Close()

	jd := json.NewDecoder(rc)
	jd.DisallowUnknownFields()
	jd.UseNumber()

	cli, err := firestore.NewClient(ctx, req.ProjectID, option.WithGRPCDialOption(grpc.WithBlock()))
	if err != nil {
		return err
	}
	defer cli.Close()

	var planItems []*planItem
	for jd.More() {
		var item planItem
		if err := jd.Decode(&item); err != nil {
			return err
		}
		planItems = append(planItems, &item)
	}

	var (
		okCount int
	)
	for i, item := range planItems {
		to := item.To
		p := path.Join(to.CollectionPath, to.DocumentID)
		data, err := m.Marshal(to.Value)
		if err != nil {
			log.Printf("%d: marshaling error: path=%v: %v", i, p, err)
			continue
		}
		doc := cli.Doc(p)
		var updates []firestore.Update
		for k, v := range data {
			updates = append(updates, firestore.Update{
				Path:  k,
				Value: v,
			})
		}
		_, err = doc.Update(ctx, updates, firestore.LastUpdateTime(to.UpdateTime))
		switch status.Code(err) {
		case codes.OK:
			okCount++
		default:
			log.Printf("%d: failed to migrate firestore document: path=%v: %v", i, p, err)
		}
	}
	fmt.Printf("(%d/%d) documents migrated!\n", okCount, len(planItems))
	return nil
}
