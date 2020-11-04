/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/5 0:00
 */
package main

import (
	"github.com/myxy99/reminder/cmd/reminder-agent/app"
	"github.com/myxy99/reminder/pkg/signals"
	"log"
	"os"
)

func main() {
	err := app.Run(signals.SetupSignalHandler())
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
