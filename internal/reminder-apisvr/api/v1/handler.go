/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 11:54
 */
package v1

import (
	"github.com/gin-gonic/gin"
	_map "github.com/myxy99/reminder/internal/reminder-apisvr/map"
	"github.com/myxy99/reminder/internal/reminder-apisvr/server"
	R "github.com/myxy99/reminder/pkg/response"
	"github.com/myxy99/reminder/pkg/validator"
	"net/http"
)

func NewReminderHandler(webServer *server.WebService, validator *validator.Validator) *ReminderHandler {
	return &ReminderHandler{webServer, validator}
}

type ReminderHandler struct {
	webServer *server.WebService
	validator *validator.Validator
}

//获取用户的提醒 分页
func (rh *ReminderHandler) GetUserReminder(ctx *gin.Context) {
	var getUserReminderService _map.GetUserReminder
	if err := ctx.ShouldBind(&getUserReminderService); err != nil {
		R.Response(ctx,
			http.StatusUnprocessableEntity,
			R.MSG_422, rh.validator.GetParamError(err),
			http.StatusUnprocessableEntity)
		return
	}
	if data, err := rh.webServer.GetUserReminder(&getUserReminderService); err == nil {
		R.Ok(ctx, R.MSG_OK, data)
	} else {
		R.Error(ctx, err.Error(), nil)
	}
	return
}

//

func NewUserHandler(webServer *server.WebService, validator *validator.Validator) *UserHandler {
	return &UserHandler{webServer, validator}
}

type UserHandler struct {
	webServer *server.WebService
	validator *validator.Validator
}

//用户登录,返回	token+email
func (uh *UserHandler) Login(ctx *gin.Context) {
	var loginService _map.UserLoginService
	if err := ctx.ShouldBind(&loginService); err != nil {
		R.Response(ctx,
			http.StatusUnprocessableEntity,
			R.MSG_422, uh.validator.GetParamError(err),
			http.StatusUnprocessableEntity)
		return
	}
	if token, email, err := uh.webServer.Login(&loginService); err == nil {
		R.Ok(ctx, R.MSG_OK, gin.H{"token": token, "email": email})
	} else {
		R.Error(ctx, err.Error(), nil)
	}
	return
}

//用户设置email
func (uh *UserHandler) SetUser(ctx *gin.Context) {
	var setUserService _map.SetUserService
	if err := ctx.ShouldBind(&setUserService); err != nil {
		R.Response(ctx,
			http.StatusUnprocessableEntity,
			R.MSG_422, uh.validator.GetParamError(err),
			http.StatusUnprocessableEntity)
		return
	}
	openid, ok := ctx.Get("openid")
	if !ok {
		R.Error(ctx, R.MSG_ERR, nil)
		return
	}
	setUserService.OpenId = openid.(string)
	if err := uh.webServer.SetUser(&setUserService); err == nil {
		R.Ok(ctx, R.MSG_OK, nil)
	} else {
		R.Error(ctx, err.Error(), nil)
	}
	return
}
