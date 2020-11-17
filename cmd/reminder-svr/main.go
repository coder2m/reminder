/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 11:03
 */
package main

import (
	"fmt"
	"github.com/myxy99/reminder/cmd/reminder-svr/app"
	"github.com/myxy99/reminder/pkg/signals"
	"os"
)

func main() {
	err := app.Run(signals.SetupSignalHandler())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
