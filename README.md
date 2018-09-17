# subflag
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
        return subflag.ErrInvalidArgument
    }
    fmt.Println(c.text)
    return nil
}
```

## Author
Nao YONASHIRO (@orisano)

## License
MIT
