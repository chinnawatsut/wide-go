package main

import "fmt"

type database interface {
	query()
}

type httpAction interface {
	Do()
}

type mongo struct {
	c httpAction
}

func (m *mongo) query() {
	fmt.Println("QUERY...")
}

func newSession() mongo {
	return mongo{
		c: newClient(),
	}
}

func newClient() httpAction {
	return &client{
		port: 8080,
		host: "asia.network",
	}
}

type client struct {
	port int
	host string
}

func (c *client) Do() {
	fmt.Println("Connecting...")
	fmt.Println("Port:", c.port)
	fmt.Println("Host:", c.host)
}

func main() {

}
