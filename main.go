package main

import (
	"context"
	"log"

	"test/bq"
	"test/protobuf"

	"cloud.google.com/go/bigquery"
)

func main() {
	ctx := context.Background()

	options := []*bq.Option{
		bq.WithQueryParameter([]bigquery.QueryParameter{
			{
				Value: "aaa",
			},
		}),
		bq.WitMaxSize(10),
		bq.WitPageToken("BFO277EWRYAQAAASA4EAAEEAQCAAKGQIBAFBB7X7777QOIFQVYKUVGYCBJWAUGIKBZWXI3RNORQWOLLMMF2GK43UCG53H3TVRYAAAAASFFPTOMRVMY3TCYRWGBQWEOBVG5STKZLGGI2TMOLCMQYGIZBUMQ3WCMJSMQYTANJQMIZRUJBTG42GEYZVMMZS2YRTME2C2NBQME3C2OJQMY2S2YJQGMZWIOJTMNSTGNLDCJCGC3TPNYZTEMBTHA3WCMZXMIZDANLGMVRWKMDDHFQWGMZXGQZWENDEG44WKZJZG43WCY3CG5RWEMBUGMYTSMZQGM2WGNTGMJTDEMBSMM3DSYRRDJSTGMRQGM4DOYJTG5RDEMBVMZSWGZJQMM4WCYZTG42DGYRUMQ3TSZLFHE3TOYLDMI3WGYRQGQZTCOJTGAZTKYZWMZRGMMRQGJRTMOLCGERTAOBTGUZTOYJSFU2TGYJSFU2DQYLGFVRGGM3GFVRDMNRTGYYGMOBTMM3GK==="),
	}

	it, err := bq.ListMessages[protobuf.TestMessage](ctx, "call nkym_test.pr_sample(?)", options...)
	if err != nil {
		panic(err)
	}

	array := []*protobuf.TestMessage{}
	for {
		m, err := it.Next()
		if err == bq.Done {
			break
		} else if err != nil {
			panic(err)
		}
		array = append(array, m)
	}
	log.Println(array)
	log.Println(it.PageToken())
}
