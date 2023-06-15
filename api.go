package ernie

import (
	"github.com/donovanhide/eventsource"
	"github.com/google/uuid"
	"net/http"
)

func New(cookie string) *Ernie {
	return &Ernie{Cookie: cookie, Session: uuid.New().String()}
}

func (bd *Ernie) SetCookie(cookie string) *Ernie {
	bd.Cookie = cookie
	return bd
}

func (bd *Ernie) Query(text string) (stream *eventsource.Stream, err error) {
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