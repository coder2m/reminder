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

func (r remindRepository) Add(remind *models.Remind) error {
	panic("implement me")
}

func (r remindRepository) Update(remind *models.Remind) error {
	panic("implement me")
}

func (r remindRepository) Del(id int) error {
	panic("implement me")
}
