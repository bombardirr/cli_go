package banner

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintBanner() {
	tapColor := color.New(color.FgHiGreen)

	// Разделяем баннер на части
	// TAP (цветной) + CLI (обычный)

	// Строка 1
	tapColor.Print("  ████████╗   █████╗    ██████╗     ")
	fmt.Println("██████╗  ██╗       ██╗")

	// Строка 2
	tapColor.Print("  ╚══██╔══╝  ██╔══██╗  ██╔══██║    ")
	fmt.Println("██╔════╝  ██║       ██║")

	// Строка 3
	tapColor.Print("     ██║     ███████║  ██████╔╝    ")
	fmt.Println("██║       ██║       ██║")

	// Строка 4
	tapColor.Print("     ██║     ██╔══██║  ██╔═══╝     ")
	fmt.Println("██║       ██║       ██║")

	// Строка 5
	tapColor.Print("     ██║     ██║  ██║  ██║         ")
	fmt.Println("╚██████╗  ███████╗  ██║")

	// Строка 6
	tapColor.Print("     ╚═╝     ╚═╝  ╚═╝  ╚═╝          ")
	fmt.Println("╚═════╝  ╚══════╝  ╚═╝")
}
