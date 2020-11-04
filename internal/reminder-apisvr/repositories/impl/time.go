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

func NewTimeRepository(db *gorm.DB) repositories.TimeRepository {
	return &timeRepository{
		db: db,
	}
}

type timeRepository struct {
	db *gorm.DB
}

// ((pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize)
func (t *timeRepository) GetByUserId(userId uint, start int, size int) (users []*models.Time,total int, err error) {
	err = t.db.Model(&models.Time{}).Where("user_id=?", userId).Limit(size).Offset(start).Find(&users).Error
	err = t.db.Model(&models.Time{}).Count(&total).Error
	return
}

func (t *timeRepository) Add(time *models.Time) (err error) {
	err = t.db.Create(time).Error
	return
}

func (t *timeRepository) Update(time *models.Time) (err error) {
	err = t.db.Model(&models.Time{}).Update(time).Error
	return
}

func (t *timeRepository) Del(id int) (err error) {
	err = t.db.Delete(&models.User{}, id).Error
	return
}
