# subflag
subflag is a subcommand library for `flag` package.

## Installation
```bash
go get github.com/orisano/subflag
```

## How to use
```go
package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    
    "github.com/orisano/subflag"
)

func main() {
    if err := subflag.SubCommand(os.Args[1:], []subflag.Command{&EchoCommand{}}); err != nil {
        log.Fatal(err)
    }
}

type EchoCommand struct {
    text string
}

func (c *EchoCommand) FlagSet() *flag.FlagSet {
    flagSet := flag.NewFlagSet("echo", flag.ExitOnError)
    flagSet.StringVar(&c.text, "t", c.text, "echo text (required)")
    return flagSet
}

func (c *EchoCommand) Run(args []string) error {
    if len(c.text) == 0 {
        return flag.ErrHelp
    }
    fmt.Println(c.text)
    return nil
}
```

## Author
Nao Yonashiro (@orisano)

## License
MIT
