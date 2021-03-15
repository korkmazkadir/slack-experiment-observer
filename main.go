package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

const channel = "experiment-status"

func main() {
	api := slack.New()

	result1, result2, err := api.PostMessage(channel, slack.MsgOptionAsUser(true), slack.MsgOptionText("*[Batch:1, Macroblock-size:1000000, CC:8]*\t_Experiment Started_", true))

	if err != nil {
		panic(err)
	}

	fmt.Println(result1)
	fmt.Println(result2)

}
