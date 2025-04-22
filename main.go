package main

import (
	"os"

	"github.com/Fuzz-Head/go-texttospeech/cmd"
)

func main() {
  cli := &cmd.CLI{ErrStream: os.Stderr}
  os.Exit(cli.Run(os.Args))
}
