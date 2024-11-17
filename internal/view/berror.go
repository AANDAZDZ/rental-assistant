package view

import (
	"errors"
	"fmt"
)

type BError struct {
	ErrNo  int
	ErrMsg string
	innErr error
}

func (b *BError) Error() string {
	return b.ErrMsg
}

func NewBError(errno int, errMsg string) *BError {
	return &BError{ErrNo: errno, ErrMsg: errMsg}
}

func AsBError(err error) (*BError, bool) {
	be := &BError{}
	isRight := errors.As(err, &be)
	return be, isRight
}

func (crp *CommonResp) AddErr(err error) {
	if bErr, ok := AsBError(err); ok {
		crp.ErrNum = bErr.ErrNo
		crp.ErrMsg = bErr.ErrMsg
	}
	return
}

func (b *BError) WithCustomMsg(msgFmt string) *BError {
	n := NewBError(b.ErrNo, b.ErrMsg)
	n.ErrMsg = fmt.Sprintf(msgFmt)
	return n
}

const (
	ErrorNumParamVerifyFailed = 10001
	ErrorNumAuthVerifyFailed  = 10002
	ErrorNumRegionNotExist    = 10003
)

var (
	ErrMsgParamVerifyFailed    = NewBError(ErrorNumParamVerifyFailed, "参数校验失败")
	ErrMsgAuthVerifyFailed     = NewBError(ErrorNumAuthVerifyFailed, "权限校验失败")
	ErrMsgCenterRegionNotExist = NewBError(ErrorNumRegionNotExist, "中心地点不存在")
)
