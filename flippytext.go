package ernie

// Copy from https://github.com/bengadbois/flippytext
// LICENSE: MIT
// AUTHOR: bengadbois
// Change: r3inbowari
import (
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

const defaultTickerTime = time.Millisecond * 10
const defaultTickerCount = 10
const defaultRandomChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" // disabled...

type FlippyText struct {
	// How long to pause between each characer flipping
	TickerTime time.Duration
	// How many times each character should flip through before resolving
	TickerCount int
	// The list of characters to use while flipping
	RandomChars string
	// Where to write the output
	Output io.Writer
}

// Flip through the characters of word, printing to stdout
func (t *FlippyText) Write(word string) error {
	if word == "" {
		return nil
	}
	if t.RandomChars == "" {
		return errors.New("random char is empty")
	}
	if t.Output == nil {
		return errors.New("nil output for writing")
	}
	_, err := fmt.Fprint(t.Output, "")
	if err != nil {
		return errors.New("unable to write to output:" + err.Error())
	}
	word = strings.TrimLeft(word, "\n")
	for _, part := range word {
		time.Sleep(t.TickerTime)
		fmt.Fprintf(t.Output, string(part))
	}
	fmt.Println()
	return nil
}
