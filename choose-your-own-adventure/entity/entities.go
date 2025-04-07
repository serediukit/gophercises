package entity

import "fmt"

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

func (c *Arc) String() string {
	return fmt.Sprintf("{\n\tTitle:%s\n\tStory:%v\n\tOptions%v\n}\n", c.Title, c.Story, c.Options)
}
