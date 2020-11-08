/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 13:43
 */
package models

import "github.com/jinzhu/gorm"

type Time struct {
	*gorm.Model
	RemindId  uint // 提醒id
	Remind    Remind
	UserId    uint //用户id
	User      User
	TimeMonth int // 确定提醒的月
	TimeDay   int // 确定提醒的日
	TimeNum   int `gorm:"default:1"` //这一个是提前多久提醒的 默认是提前一天
}
