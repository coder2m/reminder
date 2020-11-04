/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 13:43
 */
package models

import "github.com/jinzhu/gorm"

type Time struct {
	*gorm.Model
	TimeType     int `gorm:"default:1"`
	Month        int
	Day          int
	Message      string
	UserId       uint
	User         User
	ReminderTime int `gorm:"default:1"`
	Status       int `gorm:"default:1"`
}
