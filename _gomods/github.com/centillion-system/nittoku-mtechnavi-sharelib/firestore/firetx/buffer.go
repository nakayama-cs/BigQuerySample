package firetx

import (
	"cloud.google.com/go/firestore"
)

type Buffer interface {
	Doc(string) *firestore.DocumentRef

	Collection(string) *firestore.CollectionRef

	Get(*firestore.DocumentRef) (*firestore.DocumentSnapshot, error)

	GetAll([]*firestore.DocumentRef) ([]*firestore.DocumentSnapshot, error)

	DocumentRefs(*firestore.CollectionRef) *firestore.DocumentRefIterator

	Documents(firestore.Queryer) *firestore.DocumentIterator

	Create(*firestore.DocumentRef, any)

	Update(*firestore.DocumentRef, []firestore.Update, ...firestore.Precondition)

	Delete(*firestore.DocumentRef, ...firestore.Precondition)

	Flush() error
}
