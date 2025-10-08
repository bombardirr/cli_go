package main

import (
	"bufio"
	"cli_go/pkg/banner"
	"cli_go/pkg/fixes"
	"fmt"
	"os"
)

func main() {
	fixes.FixLanguageSwitch()
	banner.PrintBanner()

	fmt.Println("Нажмите Enter для выхода")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
