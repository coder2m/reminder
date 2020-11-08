/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/8 14:55
 */
package models

import "github.com/jinzhu/gorm"

type Remind struct {
	*gorm.Model
	TimeType     int    `gorm:"default:1"` //1为每年为一个周期的 2为指定时间提醒一次
	TimeDataType int    `gorm:"default:1"` //1为阳 2为农
	RemindMonth  int    // 提醒的月
	RemindDay    int    // 提醒的日
	Message      string //提醒的具体信息设置
	UserId       uint   //用户id
	User         User
	Status       int `gorm:"default:1"` //是否提醒
}
