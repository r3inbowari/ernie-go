package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/r3inbowari/ernie"
	"gopkg.in/yaml.v3"
	"os"
	"syscall"
	"time"
)

type Config struct {
	Cookie string `json:"cookie"`
}

var ai *ernie.Ernie

func main() {
	config := Config{}
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		if errors.Is(err, syscall.ERROR_FILE_NOT_FOUND) {
			fmt.Println("not found config file", err)
		} else {
			panic(err)
		}
		return
	}
	err = yaml.Unmarshal(data, &config)
	if len(config.Cookie) == 0 {
		fmt.Println("please enter your bduss in config.yaml and restart")
		time.Sleep(time.Second * 5)
		return
	}
	ai = ernie.New(config.Cookie)

	for {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Print("Me> ")
		prompts, err := inputReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		query(prompts)
	}
}

func query(prompts string) {
	stream, err := ai.Query(prompts)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case event := <-stream.Events:
			seg, err1 := ernie.ParseStreamSegment(event.Data())
			if seg.Status != 0 {
				panic(fmt.Sprintf("error response status code: %d, please check your bduss token\n", seg.Status))
			}
			if err1 != nil {
				continue
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
