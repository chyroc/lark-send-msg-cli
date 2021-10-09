package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:   "lark-send-msg-cli",
		Usage:  "send lark message",
		Action: runAction,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

func runAction(c *cli.Context) error {
	// webhook := c.String("webhook")
	// secret := c.String("secret")

	return nil
}

func appFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "webhook",
			Usage: "webhook url",
		},
		&cli.StringFlag{
			Name:  "secret",
			Usage: "webhook secret or app_secret",
		},
	}
}
