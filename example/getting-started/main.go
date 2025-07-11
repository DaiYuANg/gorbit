package main

import "github.com/daiyuang/gorbit/framework"

func main() {
	app := framework.New(UserConfig{})
	app.Run()
}
