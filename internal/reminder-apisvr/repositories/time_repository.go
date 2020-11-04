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
	GetByUserId(uint,int,int) ([]*models.Time,int,error)
}