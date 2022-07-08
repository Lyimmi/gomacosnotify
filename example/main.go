package main

import (
	"fmt"

	notify "github.com/willdot/gomacosnotify"
)

func main() {
	n, err := notify.New()
	if err != nil {
		panic(err)
	}

	notification := notify.Notification{
		Message:      "YO",
		Title:        "hello",
		SubTitle:     "world",
		ContentImage: "../RandomImage.png",
	}

	_ = notification.SetTimeout(5)

	resp, err := n.Send(notification)
	if err != nil {
		panic(err)
	}

	fmt.Printf("action: %s\n", resp.ActivationValue)
	fmt.Printf("action: %s\n", resp.ActivationType)

}
