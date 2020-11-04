/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 11:18
 */
package reminder

type Options struct {
	Spec []string `json:"spec,omitempty" yaml:"spec"`
}

func NewReminderOptions() *Options {
	return &Options{
		Spec: []string{"* * * * * *"},
	}
}
