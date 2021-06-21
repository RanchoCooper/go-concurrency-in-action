package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/RanchoCooper/go-concurrency-in-action/ch02/chatbot"
)

var chatbotName string

func init() {
	flag.StringVar(&chatbotName, "chatbot", "simple.cn", "the chatbot's name for dialogue")
	err := chatbot.Register(chatbot.NewSimpleCN("simple.cn", nil))
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	bot := chatbot.GetBot(chatbotName)
	if bot == nil {
		err := fmt.Errorf("fatal error: unsupported chatbot name %s\n", chatbotName)
		checkError(nil, err, true)
	}

	inputReader := bufio.NewReader(os.Stdin)
	begin, err := bot.Begin()
	checkError(bot, err, true)
	fmt.Println(begin)
	input, err := inputReader.ReadString('\n')
	checkError(bot, err, true)
	fmt.Println(bot.Hello(input[:len(input) - 1]))

	for {
		input, err := inputReader.ReadString('\n')
		if checkError(bot, err, false) {
			continue
		}
		output, end, err := bot.Talk(input)
		if checkError(bot, err, false) {
			continue
		}
		if output != "" {
			fmt.Println(output)
		}
		if end {
			err = bot.End()
			checkError(bot, err, false)
			os.Exit(0)
		}
	}

}

func checkError(chatbot chatbot.Chatbot, err error, exit bool) bool {
	if err == nil {
		return false
	}
	if chatbot != nil {
		fmt.Println(chatbot.ReportError(err))
	} else {
		fmt.Println(err)
	}
	if exit {
		debug.PrintStack()
		os.Exit(-1)
	}
	return true
}
