package service

import (
	"io"
	"log"
	"context"

	"github.com/google/wire"
)

type MailSender interface {
	Send()
}


type MailConfig struct {

}
 

type mailSender struct {

}

func (e *mailSender) Send() {
	log.Println("send email")
}


func NewMailSender(m *MailConfig) *mailSender {
	return &mailSender{}
}


var MailSet = wire.NewSet(NewMailSender, wire.Bind(new(MailSender), new(*mailSender)))



type Options struct {
	Messages []string
	Writer io.Writer
}
type Greeter struct {
	msg []string
	w io.Writer
}


func NewGreeter(ctx context.Context, opts *Options) (*Greeter, error) {
	return &Greeter{
		msg: opts.Messages,
		w:opts.Writer,
	}, nil
}

var GreeterSet = wire.NewSet(NewGreeter,wire.Struct(new(Options), "*"))