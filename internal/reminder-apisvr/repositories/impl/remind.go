/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/8 15:06
 */
package impl

import (
	"github.com/jinzhu/gorm"
	"github.com/myxy99/reminder/internal/reminder-apisvr/models"
	"github.com/myxy99/reminder/internal/reminder-apisvr/repositories"
)

type remindRepository struct {
	db *gorm.DB
}

func NewRemindRepository(db *gorm.DB) repositories.RemindRepository {
	return &remindRepository{
		db: db,
	}
}

func (r *remindRepository) GetByUserId(userId uint, start int, size int) (reminds []*models.Remind, total int, err error) {
	err = r.db.Model(&models.Remind{}).
		Where("user_id=?", userId).
		Limit(size).Offset(start).
		Find(&reminds).Error
	err = r.db.Model(&models.Remind{}).Count(&total).Error
	return
}

func (r *remindRepository) Add(remind *models.Remind) error {
	panic("implement me")
}

func (r *remindRepository) Update(remind *models.Remind) error {
	panic("implement me")
}

func (r *remindRepository) Del(id int) error {
	panic("implement me")
}
