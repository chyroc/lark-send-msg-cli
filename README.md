# lark-send-msg-cli

[![codecov](https://codecov.io/gh/chyroc/lark-send-msg-cli/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/lark-send-msg-cli)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/lark-send-msg-cli "go report card")](https://goreportcard.com/report/github.com/chyroc/lark-send-msg-cli)
[![test status](https://github.com/chyroc/lark-send-msg-cli/actions/workflows/test.yml/badge.svg)](https://github.com/chyroc/lark-send-msg-cli/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/lark-send-msg-cli)
[![Go project version](https://badge.fury.io/go/github.com%2Fchyroc%2Flark-send-msg-cli.svg)](https://badge.fury.io/go/github.com%2Fchyroc%2Flark-send-msg-cli)

![](./header.png)

## Install

- By Go

```shell
go install github.com/chyroc/lark-send-msg-cli@latest
```

- By Brew

```shell
brew install chyroc/tap/lark-send-msg-cli
```

## Usage

- With Webhook

```shell
lark-send-msg-cli --webhook "<webhook-url>" --secret "<secret>" --message "<message>"
```

If you send your message contains `\n`, but the message sent does not wrap, this may be because `\n` is escaped, you can add the `--wrap` parameter to reverse this situation

```shell
lark-send-msg-cli --webhook "<webhook-url>" --secret "<secret>" --message "msg1\nmsg2\n" --wrap
```

If you use this tool in GitHub action and other places, timeout may occur. You can add the `--timeout` parameter to specify the timeout period in seconds.

```shell
lark-send-msg-cli --webhook "<webhook-url>" --secret "<secret>" --message "<message>" --timeout 10
```