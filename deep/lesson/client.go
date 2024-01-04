package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Avatar struct {
	URL  string
	size int64
}

type Client struct {
	id   int64
	name string
	age  int64
	img  Avatar
}

func (c Client) HasAvatar() bool {
	if c.img.URL != "" {
		return true
	}
	return false
}

func (c *Client) UpdateAvatar(u string) {
	c.img.URL = u
}

func NewClient(name string, age int64, img Avatar) Client {
	return Client{
		id:   generateRandomID(),
		name: name,
		age:  age,
		img:  img,
	}
}

func generateRandomID() int64 {
	rand.Seed(time.Now().UnixNano())
	randomID := rand.Int63() % 1000
	return randomID
}

func main() {
	client := Client{
		id:   5,
		name: "Artem",
		age:  25,
		img: Avatar{
			URL:  "https://ya.ru",
			size: 128,
		},
	}
	fmt.Println(client)

	client.UpdateAvatar("https://google.com")
	fmt.Println(client)

	c := NewClient("Vladimir", 34, Avatar{URL: "https://vk.com"})
	fmt.Println(c)
}
