package main

import (
	"fmt"
	"os"

	"github.com/sunerpy/gitlens-patch/internal/app"
)

func main() {

	// 创建应用实例
	appInstance := app.NewApp()

	// 运行应用
	if err := appInstance.Run(); err != nil {
		fmt.Printf("运行失败: %v\n", err)
		app.WaitForKeyPress()
		os.Exit(1)
	}
}
