package firetx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"cloud.google.com/go/firestore"
	firestorepb "google.golang.org/genproto/googleapis/firestore/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Tracer struct {
	Base Buffer

	Printf func(string, ...any)
}

func formatQuery(q firestore.Query) string {
	b, err := q.Serialize()
	if err != nil {
		return err.Error()
	}
	var runQueryReq firestorepb.RunQueryRequest
	if err := proto.Unmarshal(b, &runQueryReq); err != nil {
		return err.Error()
	}
	data := json.RawMessage(protojson.Format(&runQueryReq))
	var buffer bytes.Buffer
	if err := json.Compact(&buffer, data); err != nil {
		return string(data)
	}
	return buffer.String()
}

func (t *Tracer) trace(op string, arg any) {
	if t.Printf == nil {
		return
	}

	var suffix string
	switch v := arg.(type) {
	case nil:
	case *firestore.DocumentRef:
		suffix = v.Path
	case []*firestore.DocumentRef:
		l := make([]string, len(v))
		for i := range v {
			l[i] = v[i].Path
		}
		suffix = strings.Join(l, ", ")
	case *firestore.CollectionRef:
		suffix = v.Path
	case firestore.Query:
		suffix = formatQuery(v)
	case firestore.Queryer:
		suffix = fmt.Sprint(v)
	default:
		suffix = fmt.Sprintf("unhandled %T", arg)
	}
	if suffix != "" {
		suffix = ": " + suffix
	}
	t.Printf("%T.%s%s", t.Base, op, suffix)
}

func (t *Tracer) Doc(p string) *firestore.DocumentRef {
	// no trace
	return t.Base.Doc(p)
}

func (t *Tracer) Collection(p string) *firestore.CollectionRef {
	// no trace
	return t.Base.Collection(p)
}

func (t *Tracer) Get(dr *firestore.DocumentRef) (*firestore.DocumentSnapshot, error) {
	t.trace("Get", dr)
	return t.Base.Get(dr)
}

func (t *Tracer) GetAll(drs []*firestore.DocumentRef) ([]*firestore.DocumentSnapshot, error) {
	t.trace("GetAll", drs)
	return t.Base.GetAll(drs)
}

func (t *Tracer) DocumentRefs(cr *firestore.CollectionRef) *firestore.DocumentRefIterator {
	t.trace("DocumentRefs", cr)
	return t.Base.DocumentRefs(cr)
}

func (t *Tracer) Documents(q firestore.Queryer) *firestore.DocumentIterator {
	t.trace("Documents", q)
	return t.Base.Documents(q)
}

func (t *Tracer) Create(dr *firestore.DocumentRef, data any) {
	t.trace("Create", dr)
	t.Base.Create(dr, data)
}

func (t *Tracer) Update(dr *firestore.DocumentRef, updates []firestore.Update, opts ...firestore.Precondition) {
	t.trace("Update", dr)
	t.Base.Update(dr, updates, opts...)
}

func (t *Tracer) Delete(dr *firestore.DocumentRef, opts ...firestore.Precondition) {
	t.trace("Delete", dr)
	t.Base.Delete(dr, opts...)
}

func (t *Tracer) Flush() error {
	t.trace("Flush", nil)
	return t.Base.Flush()
}
