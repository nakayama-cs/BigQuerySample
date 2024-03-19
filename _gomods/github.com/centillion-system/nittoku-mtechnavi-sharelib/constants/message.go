package constants

//go:generate generate-message-constants -o message_gen.go message.tsv
//go:generate gofmt -w message_gen.go

type MessageName string

func (m MessageName) String() string {
	return string(m)
}
