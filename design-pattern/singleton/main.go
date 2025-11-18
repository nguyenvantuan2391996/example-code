package main

import (
    "fmt"
    "sync"
)

type Config struct {
    URL string
}

var (
    instance *Config
    once     sync.Once
)

func GetConfig() *Config {
    once.Do(func() {
        fmt.Println("Creating new Config instance...")
        instance = &Config{
            URL: "https://example.com/api",
        }
    })
    return instance
}

func main() {
    c1 := GetConfig()
    c2 := GetConfig()
    c3 := GetConfig()

    fmt.Println("c1 URL:", c1.URL)
    fmt.Println("c2 URL:", c2.URL)
    fmt.Println("c3 URL:", c3.URL)

    if c1 == c2 && c2 == c3 {
        fmt.Println("All instances point to the same Config → Singleton works!")
    }
}