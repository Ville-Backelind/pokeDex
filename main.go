package main

import (
    "fmt"
    "bufio"
    //"strings"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    fmt.Println(scanner.Text())
}
