/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 13:52
 */
package repositories

import "github.com/myxy99/reminder/internal/reminder-apisvr/models"

type TimeRepository interface {
	Add(*models.Time) error
	Update(*models.Time) error
	Del(int) error
	GetByUserId(userId uint, start int, size int) ([]*models.Time, int, error)
	GetByTime(timeType int, month int, day int, status int, reminderTime int) ([]*models.Time, error)
}
