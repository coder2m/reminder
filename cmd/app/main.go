package main

import (
	"fmt"
	"github.com/coder2m/component/pkg/xcolor"
	"github.com/coder2m/component/pkg/xflag"
	"github.com/coder2m/component/pkg/xsignals"
	"github.com/coder2m/reminder/cmd/app/app"
	"os"
)

func main() {
	xflag.NewRootCommand(&xflag.Command{
		Use:                "app",
		DisableSuggestions: false,
	})
	xflag.Register(
		xflag.CommandNode{
			Name: "run",
			Command: &xflag.Command{
				Use:   "run",
				Short: "run your app",
				Run: func(cmd *xflag.Command, args []string) {
					err := app.Run(xsignals.SetupSignalHandler())
					if err != nil {
						fmt.Println(xcolor.Red(err.Error()))
						os.Exit(1)
					}
				},
			},
			Flags: func(c *xflag.Command) {
				c.Flags().StringP("xcfg", "c", "", "配置文件")
				_ = c.MarkFlagRequired("xcfg")
			},
		},
	)
	_ = xflag.Parse()
}
