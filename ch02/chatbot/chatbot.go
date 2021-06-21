package chatbot

import (
	"errors"
)

type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

type Chatbot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
}

var (
	InvalidChatbot = errors.New("invalid chatbot")
	InvalidChatbotName = errors.New("invalid chatbot name")
	ExistingChatbot = errors.New("existing chatbot")
)

var chatbotMap = map[string]Chatbot{}

func Register(chatbot Chatbot) error {
	if chatbot == nil {
		return InvalidChatbot
	}
	if chatbot.Name() == "" {
		return InvalidChatbotName
	}
	if _, ok := chatbotMap[chatbot.Name()]; ok {
		return ExistingChatbot
	}
	chatbotMap[chatbot.Name()] = chatbot

	return nil
}

func GetBot(name string) Chatbot {
	return chatbotMap[name]
}
