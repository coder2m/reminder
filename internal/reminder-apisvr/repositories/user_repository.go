/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 13:52
 */
package repositories

import "github.com/myxy99/reminder/internal/reminder-apisvr/models"

type UserRepository interface {
	Add(*models.User) error
	Update(*models.User) error
	Del(id int) error
	GetOrCreateByOpenId(openId string) (user *models.User,err error)
}
