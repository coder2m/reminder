/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 13:38
 */
package models

import "github.com/jinzhu/gorm"

type User struct {
	*gorm.Model
	Openid string `gorm:"uniqueIndex"`
	Email  string `gorm:"uniqueIndex"`
}
