package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sunerpy/gitlens-patch/internal/config"
)

// InputHandler 输入处理器
type InputHandler struct{}

// NewInputHandler 创建输入处理器实例
func NewInputHandler() *InputHandler {
	return &InputHandler{}
}

// PromptForSelection 提示用户选择选项
func PromptForSelection(maxChoice int, prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(config.ActionInput + " " + prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err == nil && choice > 0 && choice <= maxChoice {
			return choice
		}
		fmt.Printf("[ERROR] 无效的选择，请输入 1 到 %d 之间的数字\n", maxChoice)
	}
}

// PromptCustomPath 提示用户输入自定义路径
func PromptCustomPath() string {
	fmt.Print("\n" + config.ActionInput + " 请输入自定义扩展目录路径: ")
	reader := bufio.NewReader(os.Stdin)
	customPath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("[ERROR] 读取输入失败")
		return ""
	}
	return strings.TrimSpace(customPath)
}

// WaitForKeyPress 等待用户按键
func WaitForKeyPress() {
	fmt.Print("\n" + config.ActionWait + " 按回车键退出...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
