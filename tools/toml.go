package main

import (
    "fmt"
    "github.com/BurntSushi/toml"
)

//Config 設定ファイル
type Config struct {
    User UserConfig
}

//UserConfig 設定ファイルのユーザ部分
type UserConfig struct {
    Name string
    Age  int
}

var config Config

func main() {

    _, err := toml.DecodeFile("./config.toml", &config)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(config.User.Name)
}
