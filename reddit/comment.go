package reddit

import (
	"fmt"
)

type Comment struct {
	Body string `json:"body"`
	Ups int `json:"ups"`
	Downs int `json:"downs"`
	Author string `json:"author"`
	Subreddit string `json:"subreddit"`
}

func (c *Comment) String() string {
	return fmt.Sprintf("Ups: %d, Downs: %d\n Author: %s\n=>\n%s\n<=\n=========\n", c.Ups, c.Downs, c.Author, c.Body)
}

func (c *Comment) EstimateTime() int {
	return 0
}
