package main

import "fmt"
import "github.com/carlescere/scheduler"
import "runtime"

func main() {
    scheduler.Every(5).Seconds().Run(printSuccess)
    runtime.Goexit()
}

func printSuccess() {
    fmt.Printf("success!! \n")
}
