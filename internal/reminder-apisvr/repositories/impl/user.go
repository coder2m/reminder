/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 13:53
 */
package impl

import (
	"github.com/jinzhu/gorm"
	"github.com/myxy99/reminder/internal/reminder-apisvr/models"
	"github.com/myxy99/reminder/internal/reminder-apisvr/repositories"
)

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepository{
		db: db,
	}
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) GetOrCreateByOpenId(openId string) (user *models.User, err error) {
	user = &models.User{
		Openid: openId,
	}
	err = u.db.Model(&models.User{}).
		Where("openid=?", openId).
		FirstOrCreate(user).Error
	return
}

func (u *userRepository) Add(user *models.User) (err error) {
	err = u.db.Create(user).Error
	return
}

func (u *userRepository) Update(user *models.User) (err error) {
	err = u.db.Model(&models.User{}).Update(user).Error
	return
}

func (u *userRepository) Del(id int) (err error) {
	err = u.db.Delete(&models.User{}, id).Error
	return
}
