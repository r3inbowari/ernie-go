package main

import (
	"bufio"
	"fmt"
	"github.com/r3inbowari/ernie"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"time"
)

type Config struct {
	Token string `json:"token"`
}

var ai *ernie.Ernie

func main() {
	config := Config{}
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("not found config file", err)
		time.Sleep(time.Second * 5)
		os.Exit(2)
	}
	err = yaml.Unmarshal(data, &config)
	if len(config.Token) == 0 {
		fmt.Println("please enter your bduss in config.yaml and restart")
		time.Sleep(time.Second * 5)
		return
	}
	ai = ernie.New(config.Token)

	for {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Print("Me> ")
		prompts, err := inputReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		query(strings.TrimRight(prompts, "\r\n"))
	}
}

func query(prompts string) {
	stream, err := ai.Query(prompts)
	if err == ernie.ErrEmptyPrompt {
		return
	}
	if err != nil {
		panic(err)
	}
	for {
		select {
		case event := <-stream.Events:
			seg, err1 := ernie.ParseStreamSegment(event.Data())
			if err1 != nil {
				continue
			}
			if seg.Status != 0 {
				panic(fmt.Sprintf("error response status code: %d, please check your bduss token\n", seg.Status))
			}
			if !seg.Empty() {
				ft := &ernie.FlippyText{
					TickerTime:  time.Millisecond * 10,
					TickerCount: 1,
					RandomChars: "啊?",
					Output:      os.Stdout,
				}
				_ = ft.Write("Ernie> " + seg.Text())
			}
			if seg.IsCompleted() {
				return
			}
		}
	}
}
