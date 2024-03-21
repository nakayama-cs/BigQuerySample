package main

import (
	"context"
	"log"
	"os"
	pb "test/protobuf"

	"cloud.google.com/go/bigquery"
	"go.einride.tech/protobuf-bigquery/encoding/protobq"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"
)

func main() {
	// 環境変数からGCPのプロジェクトIDを取得する
	projectId := os.Getenv("GOOGLE_CLOUD_PROJECT")

	// BigQueryのクライアントを取得
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectId)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// クエリーの作成
	q := client.Query("call nkym_test.pr_sample()")

	// クエリーの実行
	it, err := q.Read(ctx)
	if err != nil {
		panic(err)
	}

	previousToken := ""
	if 2 <= len(os.Args) {
		previousToken = os.Args[1]
	}

	const maxSize = 12
	it.PageInfo().MaxSize = maxSize
	it.PageInfo().Token = previousToken

	// BigQuery->protobufの変換を行うLoaderを用意する
	messageLoader := &protobq.MessageLoader{
		Message: &pb.TestMessage{},
	}

	// クエリー結果をレコードごとに取得する
	for i := 0; i < maxSize; i++ {
		err := it.Next(messageLoader)
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}

		// protobuf変換された結果を取得
		message := proto.Clone(messageLoader.Message)
		if testMessage, ok := message.(*pb.TestMessage); ok {
			log.Println(testMessage)
		}
	}

	log.Println("page token: ", it.PageInfo().Token)
}
