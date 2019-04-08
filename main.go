package main

import (
    "fmt"
    "github.com/borderstech/logmatic"
    "github.com/gregdel/pushover"
    "os"
    "time"
)

func main() {
    c := GetConfig()
    l := logmatic.NewLogger()
    l.SetLevel(logmatic.DEBUG)
    l.Trace(fmt.Sprintf("%+v\n", c))

    // Ensure tokens exist for use
    if c.AppToken == "" {
        l.Error("Application Token was not supplied. Use -a or PUSHOVER_APP_TOKEN environment variable.")
        os.Exit(1)
    }

    if c.UserToken == "" {
        l.Error("User Token was not supplied. Use -u or PUSHOVER_USER_TOKEN environment variable.")
        os.Exit(1)
    }

    // Make sure message is defined
    if c.Message == "" {
        l.Error("Message was not supplied")
        os.Exit(1)
    }

    // All good from here onwards

    app := pushover.New(c.AppToken)
    recipient := pushover.NewRecipient(c.UserToken)

    message := &pushover.Message{
        Message:    c.Message,
        Title:      c.Title,
        Priority:   c.Priority,
        URL:        c.Url,
        URLTitle:   c.UrlTitle,
        Timestamp:  time.Now().Unix(),
        Retry:      60 * time.Second,
        Expire:     time.Hour,
        DeviceName: c.Device,
        Sound:      c.Sound,
    }

    // Send the message to the recipient
    _, err := app.SendMessage(message, recipient)
    if err != nil {
        l.Error(fmt.Sprintf("%v", err))
        os.Exit(1)
    } else {
        l.Info("Sent OK")
    }
}
