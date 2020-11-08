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

func (t *timeRepository) GetByTime(month int, day int) (times []*models.Time, err error) {
	err = t.db.Model(&models.Time{}).
		Where("time_month=?", month).
		Where("time_day=?", day).
		Preload("Remind").
		Preload("User").
		Find(&times).Error
	return
}

// ((pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize)
func (t *timeRepository) GetByUserId(userId uint, start int, size int) (times []*models.Time, total int, err error) {
	err = t.db.Model(&models.Time{}).Where("user_id=?", userId).Limit(size).Offset(start).Find(&times).Error
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
	err = t.db.Delete(&models.Time{}, id).Error
	return
}

//func (t *timeRepository) GetByTime(timeType int, month int, day int, status int, reminderTime int) (times []*models.Time, err error) {
//	err = t.db.Model(&models.Time{}).
//		Where("time_type=?", timeType).
//		Where("reminder_time=?", reminderTime).
//		Where("status=?", status).
//		Where("month=?", month).
//		Where("day=?", day).
//		Preload("User").
//		Find(&times).Error
//	return
//}
