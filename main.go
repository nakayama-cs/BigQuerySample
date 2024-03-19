package main

import (
	"context"
	"fmt"
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
	q := client.Query("SELECT 'Hello' as name ")

	// クエリーの実行
	it, err := q.Read(ctx)
	if err != nil {
		panic(err)
	}

	// BigQuery->protobufの変換を行うLoaderを用意する
	messageLoader := &protobq.MessageLoader{
		Message: &pb.TestMessage{},
	}

	// クエリー結果をレコードごとに取得する
	for {
		err := it.Next(messageLoader)
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}

		// protobuf変換された結果を取得
		message := proto.Clone(messageLoader.Message)
		fmt.Println(message)
	}
}
