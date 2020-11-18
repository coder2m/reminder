/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/8 16:09
 */
package _map

type UserLoginService struct {
	Code string `form:"code" json:"code" binding:"required" label:"code"`
}

type SetUserService struct {
	OpenId string `json:"open_id"`
	Email  string `form:"email" json:"email" binding:"required,email" label:"邮箱"`
}

type GetUserReminder struct {
	OpenId   string `json:"open_id"`
	PageSize int    `json:"size" form:"size" binding:"required,numeric"`
	Page     int    `json:"page" form:"page" binding:"required,numeric"`
}