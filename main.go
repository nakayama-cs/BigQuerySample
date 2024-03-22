package main

import (
	"context"
	"log"
	bq "test/bp"
	"test/protobuf"

	"google.golang.org/api/iterator"
)

func main() {

	ctx := context.Background()
	it, err := bq.ListMessages(ctx, "call nkym_test.pr_sample()", protobuf.TestMessage{})
	if err != nil {
		panic(err)
	}

	for {
		v, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}

		log.Println("name is ", v.Name)
	}

}
