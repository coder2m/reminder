/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 11:18
 */
package email

type Options struct {
	Host     string `json:"host,omitempty" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Username string `json:"username,omitempty" yaml:"username"`
	Password string `json:"-" yaml:"password"`
}

func NewEmailOptions() *Options {
	return &Options{
		Host:     "127.0.0.1",
		Port:     5672,
		Username: "root",
		Password: "root",
	}
}
