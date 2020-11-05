/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 11:28
 */
package server

type Options struct {
	Addr string `json:"addr" yaml:"addr"`
	Locale string `json:"locale" yaml:"locale"`
}

func NewServerOptions() *Options {
	return &Options{
		Addr: "",
		Locale:"zh",
	}
}
