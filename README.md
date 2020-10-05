# Telegram

This module wraps Telegram Bot API using only standard library without any external dependencies.

## Usage

Add this module to the project:

```shell script
go get github.com/trickstersio/telegram
```

Create client instance and start calling methods:

```golang
package main

import (
    "context"
    "github.com/trickstersio/telegram"
)

func main()  {
    client := telegram.NewClient()

    msg, err := client.SendMessage(context.Background(), SendMessageArgs{
        ChatID: 1,
        Text: "Hello, World!",
    })

    if err != nil {
        log.Fatal(err)
    }

    log.Println("Sent message", msg.ID)
}
```

## Test

You can run tests using following command:

```shell script
go test github.com/trickstersio/telegram
```
