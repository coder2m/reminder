/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 11:28
 */
package wx

type Options struct {
	AppID  string `json:"appId" yaml:"appId"`
	Secret string `json:"secret" yaml:"secret"`
}

func NewWXOptions() *Options {
	return &Options{
		AppID: "",
		Secret: "",
	}
}
