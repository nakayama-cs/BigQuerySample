// firetx はfirestoreのTransactionを利用しやすくするパッケージである。
//
// firestoreのTransactionは、1度でもWrite操作をすると、以降のRead操作がすべてエラーとなる。
// そのためRead操作を先にやらねばならず、ロジックの記述について、自由度が低くなってしまう。
// 本パッケージを用いることで、Write操作を遅延させることができるため、ロジック記述の自由度を担保することができる。
package firetx

import (
	"cloud.google.com/go/firestore"
)

var (
	NoRetry = firestore.MaxAttempts(1)
)

type firestoreClient interface {
	Doc(string) *firestore.DocumentRef
	Collection(string) *firestore.CollectionRef
}
