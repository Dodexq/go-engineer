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

func main() {
	i := new(int)
	_ = i

	client := Client{}
	fmt.Printf("%#v\n", client)
	updateAvatar(&client)
	fmt.Printf("%#v\n", client)
	updateName(&client)
	fmt.Printf("%#v\n", client.Name)
}

func updateAvatar(client *Client) {
	client.IMG.URL = "updated_url"
}

func updateName(client *Client) {
	client.Name = "Artem"
}
