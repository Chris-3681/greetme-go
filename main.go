package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"
    "time"
)

// greet takes a raw name string and a shout flag.
// It cleans the name, falls back to "stranger" if empty,
// and returns a greeting. If shout is true, the greeting is uppercased.
func greet(name string, shout bool) string {
    name = strings.TrimSpace(name)
    if name == "" {
        name = "stranger"
    }

    msg := "Hello, " + name + "! Welcome to Go 🎉"

    if shout {
        msg = strings.ToUpper(msg)
    }

    return msg
}

func main() {
    // Define --shout flag. Usage example:
    //   go run main.go -- --shout
    shout := flag.Bool("shout", false, "Print greeting in uppercase")
    flag.Parse()

    fmt.Print("Enter your name: ")

    reader := bufio.NewReader(os.Stdin)
    name, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    fmt.Println(greet(name, *shout))
    fmt.Println("Generated at:", time.Now().Format(time.RFC1123))
}