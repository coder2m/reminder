/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/8 15:06
 */
package repositories

import "github.com/myxy99/reminder/internal/reminder-apisvr/models"

type RemindRepository interface {
	Add(remind *models.Remind) error
	Update(remind *models.Remind) error
	Del(id int) error
}
