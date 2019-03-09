package main

import (
	"gh-activist/cli"
	"log"
	"os"
)

func main() {
	c, err := cli.New()
	if err != nil {
        log.Fatal(err)
	}

	os.Exit(c.Run())
}
