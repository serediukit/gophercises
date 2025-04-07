package decoder

import (
	"encoding/json"
	"fmt"
)

type Option struct {
	Text string
	Arc  string
}

type Arc struct {
	Title   string
	Story   []string
	Options []Option
}

type Stories map[string]*Arc

func Decode(data []byte) (Stories, error) {
	stories := make(Stories)
	err := json.Unmarshal(data, &stories)
	if err != nil {
		return nil, err
	}
	return stories, nil
}

func (c *Arc) String() string {
	return fmt.Sprintf("{\n\tTitle:%s\n\tStory:%v\n\tOptions%v\n}\n", c.Title, c.Story, c.Options)
}
