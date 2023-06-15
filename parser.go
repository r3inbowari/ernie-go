package ernie

import (
	"encoding/json"
	"errors"
)

func ParseStreamSegment(raw string) (*StreamSegment, error) {
	if len(raw) == 0 {
		return nil, errors.New("empty data")
	}
	bdr := StreamSegment{}
	if err := json.Unmarshal([]byte(raw), &bdr); err != nil {
		return nil, err
	}
	return &bdr, nil
}

func (bdr *StreamSegment) IsCompleted() bool {
	return bdr.State() == TypeCompleted
}

func (bdr *StreamSegment) State() string {
	return bdr.Data.Message.MetaData.State
}

func (bdr *StreamSegment) Empty() bool {
	return bdr.Text() == ""
}

func (bdr *StreamSegment) Text() string {
	return bdr.Data.Message.Content.Generator.Text
}
