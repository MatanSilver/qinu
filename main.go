package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "qinu"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "field, f",
			Usage: "Only take data from field",
			Value: 1,
		},
		cli.StringFlag{
			Name:  "delimiter, d",
			Usage: "Specify a delimiter",
			Value: " ",
		},
		cli.StringFlag{
			Name:  "file",
			Usage: "Specify a file to read from",
		},
	}
	app.Usage = "Suppresses the first instance of a line, and echos subsequent instances"
	app.Action = func(c *cli.Context) error {
		f := c.Int("field")
		d := c.String("delimiter")
		filename := c.String("file")
		var file *os.File
		if filename != "" {
      var err error
			file, err = os.Open(filename)
      defer file.Close()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			file = os.Stdin
		}
		var set map[string]bool
		set = make(map[string]bool)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Split(line, d)
			if set[fields[f-1]] == false {
				set[fields[f-1]] = true
			} else {
				fmt.Println(line)
			}
		}
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
		return nil
	}
	app.Run(os.Args)
}
