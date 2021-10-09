package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chyroc/lark"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:   "lark-send-msg-cli",
		Usage:  "send lark message",
		Action: runAction,
		Flags:  appFlags(),
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

func runAction(c *cli.Context) error {
	webhook := c.String("webhook")
	secret := c.String("secret")
	message := c.String("message")

	larkCli := lark.New(lark.WithCustomBot(webhook, secret))

	resp, response, err := larkCli.Message.Send().SendText(context.Background(), message)
	if response != nil {
		fmt.Println("request-id:", response.RequestID)
	}
	if err != nil {
		return err
	}

	if resp != nil {
		fmt.Println("message-id:", resp.MessageID)
	}

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
		&cli.StringFlag{
			Name:     "message",
			Usage:    "send message",
			Required: true,
		},
	}
}
