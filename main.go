package main

import (
	"fmt"
	"os"

	"shortLink/web"
)

func main() {
	// 运行短连接服务器
	fmt.Fprintln(os.Stderr, web.New(8000).Run().Error())
}
