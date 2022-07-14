package slackbot

import "github.com/kitabisa/golang-poc/internal/app/service"

type ISlackBot interface {
}

type slackBot struct {
	opt service.Option
}

func NewSlackBot(opt service.Option) ISlackBot {
	return &slackBot{
		opt: opt,
	}
}

func Execute() {

}
