package firetx_test

import (
	"context"
	"fmt"
	"mtechnavi/sharelib/firestore/firetx"

	"cloud.google.com/go/firestore"
)

//nolint:errcheck
func ExampleTracer() {
	ctx := context.Background()

	// XXX: 実際に接続させないため、プロジェクトIDは適当
	cli, err := firestore.NewClient(ctx, "(^-^)")
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	var buffer firetx.Buffer = &nilBuffer{}
	buffer = &firetx.Tracer{
		Base: buffer,
		Printf: func(format string, args ...any) {
			fmt.Printf(format+"\n", args...)
		},
	}

	cr := cli.Collection("col")
	dr := cr.Doc("xxx")
	q := cr.Query.Where("hoge", "==", 3)

	buffer.Get(dr)
	buffer.GetAll([]*firestore.DocumentRef{dr})
	buffer.DocumentRefs(cr)
	buffer.Documents(q)
	buffer.Create(dr, map[string]any{})
	buffer.Update(dr, []firestore.Update{})
	buffer.Delete(dr)
	buffer.Flush()

	// Output:
	// *firetx_test.nilBuffer.Get: projects/(^-^)/databases/(default)/documents/col/xxx
	// *firetx_test.nilBuffer.GetAll: projects/(^-^)/databases/(default)/documents/col/xxx
	// *firetx_test.nilBuffer.DocumentRefs: projects/(^-^)/databases/(default)/documents/col
	// *firetx_test.nilBuffer.Documents: {"parent":"projects/(^-^)/databases/(default)/documents","structuredQuery":{"from":[{"collectionId":"col"}],"where":{"fieldFilter":{"field":{"fieldPath":"hoge"},"op":"EQUAL","value":{"integerValue":"3"}}}}}
	// *firetx_test.nilBuffer.Create: projects/(^-^)/databases/(default)/documents/col/xxx
	// *firetx_test.nilBuffer.Update: projects/(^-^)/databases/(default)/documents/col/xxx
	// *firetx_test.nilBuffer.Delete: projects/(^-^)/databases/(default)/documents/col/xxx
	// *firetx_test.nilBuffer.Flush
}

type nilBuffer struct{}

func (*nilBuffer) Doc(string) *firestore.DocumentRef {
	return nil
}

func (*nilBuffer) Collection(string) *firestore.CollectionRef {
	return nil
}

func (*nilBuffer) Get(*firestore.DocumentRef) (*firestore.DocumentSnapshot, error) {
	return nil, nil
}

func (*nilBuffer) GetAll([]*firestore.DocumentRef) ([]*firestore.DocumentSnapshot, error) {
	return nil, nil
}

func (*nilBuffer) DocumentRefs(*firestore.CollectionRef) *firestore.DocumentRefIterator {
	return nil
}

func (*nilBuffer) Documents(firestore.Queryer) *firestore.DocumentIterator {
	return nil
}

func (*nilBuffer) Create(*firestore.DocumentRef, any) {
}

func (*nilBuffer) Update(*firestore.DocumentRef, []firestore.Update, ...firestore.Precondition) {
}

func (*nilBuffer) Delete(*firestore.DocumentRef, ...firestore.Precondition) {
}

func (*nilBuffer) Flush() error {
	return nil
}
