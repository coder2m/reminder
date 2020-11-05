/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 13:43
 */
package models

import "github.com/jinzhu/gorm"

type Time struct {
	*gorm.Model
	TimeType     int `gorm:"default:1"` //1为阳 2为农
	Month        int
	Day          int
	Message      string
	UserId       uint
	User         User
	ReminderTime int `gorm:"default:1"` //1为提前一天提醒，2为提前1天和3天提醒，3为提前1，3，5都提醒
	Status       int `gorm:"default:1"`
}
