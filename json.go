package main

import (
    "flag"
    "fmt"
    "os"
    "scanner"
)

type Decoder struct {
    encoding string
}

/*func (d *Decoder) decode(src io.Reader) {

}
*/
func main() {
    flag.Parse()
    if flag.NArg() != 1 {
        fmt.Println("wrong number of arguments")
        os.Exit(1)
    }

    file, err := os.Open(flag.Arg(0), os.O_RDONLY, 0)
    if file == nil {
        fmt.Printf("can't open file; err=%s\n", err.String())
        os.Exit(1)
    }

    var s scanner.Scanner
    s.Init(file)
    tok := s.Scan()
    for tok != scanner.EOF {
        fmt.Println(scanner.TokenString(tok))
        tok = s.Scan()
    }
    fmt.Println()

    file.Close()
}
