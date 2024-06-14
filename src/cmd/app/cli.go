package main

import (
	"fmt"
	"os"

	"github.com/danielgtaylor/huma/v2/humacli"
)

/*
*
This Repreesnts the CLI option that we can give while running the program from CLI.
*/
type CliOptions struct {
	Debug bool   `doc:"Enable debug logging`
	Host  string `doc:"HOST ON WHICH TO LISTEN`
	Port  string `doc:"Port to Listen to"`
}

func main() {

	cli := humacli.New(func(hooks humacli.Hooks, options *CliOptions) {
		fmt.Printf("CLI was started with the following options Debug=%v Host=%v Port=%v", options.Debug, options.Host, options.Port)
		os.Exit(1)
	})

	cli.Run()
}
