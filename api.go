package ernie

import (
	"errors"
	"fmt"
	"github.com/donovanhide/eventsource"
	"github.com/google/uuid"
	"net/http"
)

var (
	ErrEmptyPrompt = errors.New("empty prompt")
)

func New(token string) *Ernie {
	return &Ernie{Cookie: fmt.Sprintf("BDUSS=%s;", token), Session: uuid.New().String()}
}

func (bd *Ernie) SetCookie(cookie string) *Ernie {
	bd.Cookie = cookie
	return bd
}

func (bd *Ernie) Query(text string) (stream *eventsource.Stream, err error) {
	if len(text) == 0 || text == "\r" || text == "\r\n" {
		return nil, ErrEmptyPrompt
	}
	wrapper, err := bd.messageWrapper(text)
	if err != nil {
		return
	}
	req, err := http.NewRequest(http.MethodPost, BaiduAPI, wrapper)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", bd.Cookie)
	if err != nil {
		return
	}
	return eventsource.SubscribeWithRequest("", req)
}
