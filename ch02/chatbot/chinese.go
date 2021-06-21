package chatbot

import (
	"fmt"
	"strings"
)

type simpleCN struct {
	name string
	talk Talk
}

func (s simpleCN) Name() string {
	return s.name
}

func (s simpleCN) Begin() (string, error) {
	return "请输入你的名字: ", nil
}

func (s simpleCN) Hello(userName string) string {
	userName = strings.TrimSpace(userName)
	if s.talk != nil {
		return s.talk.Hello(userName)
	}
	return fmt.Sprintf("你好， %s! 我可以为你做些什么？", userName)
}

func (s simpleCN) Talk(heard string) (saying string, end bool, err error) {
	heard = strings.TrimSpace(heard)
	if s.talk != nil {
		return s.talk.Talk(heard)
	}

	switch heard {
	case "":
		return
	case "没有", "再见":
		saying = "再见!"
		end = true
		return
	default:
		saying = "对不起, 我没听懂你说的"
		return
	}
}

func (s simpleCN) ReportError(err error) string {
	return fmt.Sprintf("发生了一个错误: %s\n", err)
}

func (s simpleCN) End() error {
	return nil
}

func NewSimpleCN(name string, talk Talk) Chatbot {
	return &simpleCN{
		name: name,
		talk: talk,
	}
}
