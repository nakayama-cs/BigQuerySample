package csv

import (
	"fmt"
	"mtechnavi/sharelib/constants"
	pb "mtechnavi/sharelib/protobuf"
	"mtechnavi/sharelib/text/message"
)

type Error struct {
	Line int

	Column int

	Causes []error

	// エラー種別
	ErrorLevel pb.ErrorLevel

	// システムで定義されたメッセージコード
	MessageName constants.MessageName

	// メッセージパラメータ
	MessageArgs []any
}

func (e *Error) Error() string {
	return fmt.Sprintf("CSVError Line:%d,Column:%d,Cause:%v,ErrorLevel:%v,MessageName:%s,MessageArgs:%v", e.Line, e.Column, e.Causes, e.ErrorLevel, e.MessageName.String(), e.MessageArgs)
}

func (e *Error) GetLine() int {
	return e.Line
}

func (e *Error) GetColumn() int {
	return e.Column
}

func (e *Error) GetErrorContent(lang string) string {
	T := message.T(lang)

	return T(e.MessageName.String(), e.MessageArgs...)
}

func (e *Error) GetErrorLevel() pb.ErrorLevel {
	return e.ErrorLevel
}
