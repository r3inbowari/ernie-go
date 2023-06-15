package ernie

import (
	"bytes"
	"encoding/json"
)

const (
	QueryTypeDefault = iota
)

const (
	BaiduAPI = "https://chat-ws.baidu.com/aichat/api/conversation"

	InputTypeKeyboard = "keyboard"

	TypeWaiting    = "waiting-resp"
	TypeGenerating = "generating-resp"
	TypeCompleted  = "generate-complete"
)

type Content struct {
	Query    string   `json:"query"`
	Type     int      `json:"qtype"`
	BotQuery struct{} `json:"botQuery"`
}

type Message struct {
	Method    string  `json:"inputMethod"`
	IsRebuild bool    `json:"isRebuild"`
	Content   Content `json:"content"`
}

type Request struct {
	Message    Message `json:"message"`
	Session    string  `json:"sessionId"`
	AISearchID string  `json:"aisearchId"`
	PvID       string  `json:"pvId"`
}

type StreamSegment struct {
	Status    int    `json:"status"`
	QueryID   string `json:"qid"`
	PackageID string `json:"pkgId"`
	SessionID string `json:"sessionId"`
	IsDefault int    `json:"isDefault"`
	IsShow    int    `json:"isShow"`
	Data      struct {
		Message struct {
			ID         string `json:"msgId"`
			IsRebuild  bool   `json:"isRebuild"`
			UpdateTime string `json:"updateTime"`
			MetaData   struct {
				State    string `json:"state"`
				EndTurn  bool   `json:"endTurn"`
				UserInfo struct {
					Status int `json:"status"`
				} `json:"userInfo"`
			} `json:"metaData"`
			Content struct {
				Generator struct {
					Text       string `json:"text"`
					Type       string `json:"type"`
					ShowType   string `json:"showType"`
					AntiFlag   int    `json:"antiFlag"`
					IsFinished bool   `json:"isFinished"`
				} `json:"generator"`
			} `json:"content"`
		} `json:"message"`
	} `json:"data"`
}

type Ernie struct {
	Cookie  string
	Session string
}

func (bd *Ernie) messageWrapper(msg string) (*bytes.Reader, error) {
	baiduMsg := Request{
		Message: Message{
			Content: Content{
				Query: msg,
				Type:  QueryTypeDefault,
				// BotQuery: {},
			},
			Method: InputTypeKeyboard,
		},
		Session:    bd.Session,
		AISearchID: "",
		PvID:       "",
	}
	marshal, err := json.Marshal(baiduMsg)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(marshal), nil
}
