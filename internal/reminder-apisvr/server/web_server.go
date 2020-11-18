/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/8 15:50
 */
package server

import (
	"errors"
	"github.com/myxy99/reminder/internal/reminder-apisvr/jwt"
	_map "github.com/myxy99/reminder/internal/reminder-apisvr/map"
	"github.com/myxy99/reminder/internal/reminder-apisvr/models"
	"github.com/myxy99/reminder/internal/reminder-apisvr/repositories"
	"github.com/myxy99/reminder/pkg/log"
	R "github.com/myxy99/reminder/pkg/response"
)

type WebService struct {
	user   repositories.UserRepository
	remind repositories.RemindRepository
}

func NewWebService(user repositories.UserRepository, remind repositories.RemindRepository) *WebService {
	return &WebService{user, remind}
}

func (uh *WebService) Login(server *_map.UserLoginService) (token, email string, err error) {
	wxc, err := GetWx()
	if err != nil {
		log.Error(err.Error())
		err = errors.New("系统错误")
		return
	}
	openId, _, _, err := wxc.Login(server.Code)
	if err != nil {
		log.Error(err.Error())
		err = errors.New("登录失败")
		return
	}
	user, err := uh.user.GetOrCreateByOpenId(openId)
	if err != nil {
		log.Error(err.Error())
		err = errors.New("存储用户信息错误，登录失败")
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
	if err != nil {
		log.Error(err.Error())
		return errors.New("设置失败")
	}
	return
}

func (uh *WebService) GetUserReminder(data *_map.GetUserReminder) (pageData R.PageData, err error) {
	userData, err := uh.user.GetOrCreateByOpenId(data.OpenId)
	if err != nil {
		log.Error(err.Error())
		return pageData, errors.New("获取用户信息错误")
	}
	remindData, total, err := uh.remind.GetByUserId(userData.ID, (data.Page-1)*data.PageSize, data.PageSize)
	if err != nil {
		log.Error(err.Error())
		return pageData, errors.New("获取用户提醒列表信息错误")
	}
	return R.Page(total, data.Page, data.PageSize, remindData), nil
}
