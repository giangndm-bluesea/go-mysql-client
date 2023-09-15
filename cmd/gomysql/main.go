package main

import (
	"flag"
	"log"
	"os"

	gomysqlclient "github.com/johejo/go-mysql-client"
)

func main() {
	var c gomysqlclient.Config
	var fileName string
	var command string

	flag.StringVar(&c.Host, "h", "localhost", "host")
	flag.UintVar(&c.Port, "P", 3306, "port")
	flag.StringVar(&c.User, "u", "", "user")
	flag.StringVar(&c.Password, "p", "", "password")
	flag.StringVar(&command, "c", "", "command")
	flag.StringVar(&fileName, "f", "", "file")
	flag.Parse()
	e := []string{}

	cli, err := gomysqlclient.NewCli(&c)
	if err != nil {
		log.Fatal(err)
	}
	if fileName != "" {
		f, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		e = gomysqlclient.QueriesFromReader(f)
	} else if command != "" {
		e = []string{command}
	}
	if err := cli.Run(e...); err != nil {
		log.Fatal(err)
	}
}
