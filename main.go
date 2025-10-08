package main

import (
	"bufio"
	"cli_go/pkg/fixes"
	"cli_go/pkg/menu"
	"fmt"
	"os"
)

func main() {
	fixes.FixLanguageSwitch()
	menu.ShowMenu()

	fmt.Println("Нажмите Enter для выхода")
	_, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return
	}
}
