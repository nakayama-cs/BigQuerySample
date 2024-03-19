package message_test

import (
	"context"
	"encoding/json"
	"fmt"
	"mtechnavi/sharelib/text/message"
	"os"

	"golang.org/x/text/language"
)

func init() {
	message.Loader = func(ctx context.Context, lang language.Tag) (map[string]string, error) {
		var data map[string]string
		b, err := os.ReadFile("./testdata/dict.json")
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(b, &data); err != nil {
			panic(err)
		}
		return data, nil
	}
}

func Example() {
	T := message.T("ja-JP")

	fmt.Println(T("E01", 0, 1))
	fmt.Println(T("E02", 0, 1))
	fmt.Println(T("E03", 0, 1))
	fmt.Println(T("E04", 0, 1))

	// Output:
	// Hello
	// Hello 0
	// Hello 0 1
	// Hello 0 1 0
}
