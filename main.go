package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()

	app.Name = "tcp"
	app.Usage = "Check if the tcp port is available"
	app.Version = "0.1.0"
	app.Author = "Axetroy"
	app.Email = "axetroy.dev@gmail.com"

	app.Action = func(c *cli.Context) error {
		address := c.Args().Get(0)

		host, port, err := net.SplitHostPort(address)

		if err != nil {
			return err
		}

		timeout := time.Second * 5

		if conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout); err != nil {
			fmt.Println("Connecting error:", err)
		} else if conn != nil {
			defer conn.Close()
			fmt.Println("Opened", net.JoinHostPort(host, port))
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
