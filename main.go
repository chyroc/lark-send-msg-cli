package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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
	timeout := c.Int64("timeout")

	opts := []lark.ClientOptionFunc{lark.WithCustomBot(webhook, secret)}
	if timeout > 0 {
		opts = append(opts, lark.WithTimeout(time.Second*time.Duration(timeout)))
	}

	larkCli := lark.New(opts...)

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
		&cli.Int64Flag{
			Name:  "timeout",
			Usage: "timeout second for send lark request",
		},
		&cli.StringFlag{
			Name:     "message",
			Usage:    "send message",
			Required: true,
		},
	}
}
