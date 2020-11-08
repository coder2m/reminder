/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/8 15:50
 */
package server

import (
	"github.com/myxy99/reminder/internal/reminder-apisvr/jwt"
	_map "github.com/myxy99/reminder/internal/reminder-apisvr/map"
	"github.com/myxy99/reminder/internal/reminder-apisvr/models"
	"github.com/myxy99/reminder/internal/reminder-apisvr/repositories"
)

type WebService struct {
	user repositories.UserRepository
}

func NewWebService(user repositories.UserRepository) *WebService {
	return &WebService{user}
}

func (uh *WebService) Login(server *_map.UserLoginService) (token, email string, err error) {
	wxc, err := GetWx()
	if err != nil {
		return
	}
	openId, _, _, err := wxc.Login(server.Code)
	if err != nil {
		return
	}
	user, err := uh.user.GetOrCreateByOpenId(openId)
	if err != nil {
		return
	}
	email = user.Email
	info := &jwt.Info{Openid: openId}
	token, err = info.GenerateToken()
	return
}

func (uh *WebService) SetUser(server *_map.SetUserService) (err error) {
	var user *models.User
	user = &models.User{
		Openid: server.OpenId,
		Email:  server.Email,
	}
	err = uh.user.Update(user)
	return
}
