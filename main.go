package main

import (
	"code.google.com/p/go.crypto/bcrypt"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	var cliOpts struct {
		Cost int `short:"c" long:"cost" description:"Custom cost factor"`
	}

	cliOpts.Cost = 12

	parser := flags.NewParser(&cliOpts, flags.Default)
	parser.Usage = "PASSWORD"

	args, err := parser.Parse()
	if err != nil {
		panic(err)
	}

	if cliOpts.Cost < bcrypt.MinCost {
		fmt.Printf("Minimum cost is %d.\n", bcrypt.MinCost)
		os.Exit(1)
	} else if cliOpts.Cost > bcrypt.MaxCost {
		fmt.Printf("Maximum cost is %d.\n", bcrypt.MaxCost)
		os.Exit(1)
	}

	if len(args) == 0 {
		parser.WriteHelp(os.Stderr)
		os.Exit(1)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(args[0]), cliOpts.Cost)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(hash))
}
