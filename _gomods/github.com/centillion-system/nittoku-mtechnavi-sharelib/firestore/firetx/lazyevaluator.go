package firetx

import (
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/hashicorp/go-multierror"
)

type lazyFunc func() error

type lazyError struct {
	Op string

	Ref *firestore.DocumentRef

	Err error
}

func (e *lazyError) Error() string {
	return fmt.Sprintf("%v: %v: %v", e.Op, e.Err, e.Ref.Path)
}

type LazyEvaluator struct {
	Client firestoreClient

	Tx *firestore.Transaction

	delayed []lazyFunc
}

func (le *LazyEvaluator) Doc(p string) *firestore.DocumentRef {
	return le.Client.Doc(p)
}

func (le *LazyEvaluator) Collection(p string) *firestore.CollectionRef {
	return le.Client.Collection(p)
}

func (le *LazyEvaluator) Get(dr *firestore.DocumentRef) (*firestore.DocumentSnapshot, error) {
	return le.Tx.Get(dr)
}

func (le *LazyEvaluator) GetAll(drs []*firestore.DocumentRef) ([]*firestore.DocumentSnapshot, error) {
	return le.Tx.GetAll(drs)
}

func (le *LazyEvaluator) DocumentRefs(cr *firestore.CollectionRef) *firestore.DocumentRefIterator {
	return le.Tx.DocumentRefs(cr)
}

func (le *LazyEvaluator) Documents(q firestore.Queryer) *firestore.DocumentIterator {
	return le.Tx.Documents(q)
}

func (le *LazyEvaluator) Create(dr *firestore.DocumentRef, data any) {
	le.delayed = append(le.delayed, func() error {
		if err := le.Tx.Create(dr, data); err != nil {
			return &lazyError{
				Op:  "Create",
				Ref: dr,
				Err: err,
			}
		}
		return nil
	})
}

func (le *LazyEvaluator) Update(dr *firestore.DocumentRef, updates []firestore.Update, opts ...firestore.Precondition) {
	le.delayed = append(le.delayed, func() error {
		if err := le.Tx.Update(dr, updates, opts...); err != nil {
			return &lazyError{
				Op:  "Update",
				Ref: dr,
				Err: err,
			}
		}
		return nil
	})
}

func (le *LazyEvaluator) Delete(dr *firestore.DocumentRef, opts ...firestore.Precondition) {
	le.delayed = append(le.delayed, func() error {
		if err := le.Tx.Delete(dr, opts...); err != nil {
			return &lazyError{
				Op:  "Delete",
				Ref: dr,
				Err: err,
			}
		}
		return nil
	})
}

func (le *LazyEvaluator) Flush() error {
	var err error
	for _, fn := range le.delayed {
		if e := fn(); e != nil {
			err = multierror.Append(err, e)
		}
	}
	return err
}
