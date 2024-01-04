package main

import "fmt"

type Avatar struct {
	URL  string
	Size int64
}

type Client struct {
	ID   int64
	Name string
	Age  int
	IMG  Avatar
}

func (c Client) HasAvatar() bool {
	if c.IMG.URL != "" {
		return true
	}
	return false
}

func (c *Client) UpdateAvatar(s string) {
	c.IMG.URL = s
}

func main() {
	cli := Client{
		ID:   1,
		Name: "Artem",
		Age:  28,
		IMG: Avatar{
			URL:  "https://ya.ru",
			Size: 128,
		},
	}
	fmt.Println(cli.HasAvatar())
	cli.UpdateAvatar("https://google.com")
	fmt.Println(cli.IMG.URL)

}
