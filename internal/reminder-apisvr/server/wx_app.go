/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/5 18:25
 */
package server

import (
	"errors"
	"github.com/medivhzhan/weapp/v2"
)

type WxApp struct {
	Secret string
	AppId  string
}

var wx *WxApp

func NewWxAPP(Secret, AppId string) {
	wx = &WxApp{Secret, AppId}
}

func (w *WxApp) Login(code string) (OpenID, SessionKey, UnionID string, err error) {
	res, err := weapp.Login(w.AppId, w.Secret, code)
	if err != nil {
		return
	}

	if err = res.GetResponseError(); err != nil {
		// 处理微信返回错误信息
		return
	}
	return res.OpenID, res.SessionKey, res.UnionID, err
}

func GetWx() (*WxApp, error) {
	if wx == nil {
		return wx, errors.New("wx Uninitialized")
	}
	return wx, nil
}
