package menu

import (
	"cli_go/pkg/banner"
	"fmt"

	"github.com/eiannone/keyboard"
)

// Чистим экран
func clearScreen() {
	// Перемещаем курсор в начало
	fmt.Print("\033[H")
	// Очищаем от курсора до конца экрана
	fmt.Print("\033[J")
}

func ShowMenu() {

	// Активируем клаву
	err := keyboard.Open()
	if err != nil {
		return
	}
	defer func() {
		err := keyboard.Close()
		if err != nil {

		}
	}()

	selected := 1

	for {
		clearScreen()
		banner.PrintBanner()

		fmt.Println("")
		fmt.Println("Используйте стрелки ↑↓ для выбора, Enter для подтверждения:")

		if selected == 1 {
			fmt.Println("► 1. Настройка SSH соединения с Git")
		} else {
			fmt.Println("  1. Настройка SSH соединения с Git")
		}

		if selected == 2 {
			fmt.Println("► 2. Изменение старой SSH ссылки под Ваш config")
		} else {
			fmt.Println("  2. Изменение старой SSH ссылки под Ваш config")
		}

		if selected == 0 {
			fmt.Println("► 0. Выход (Esc)")
		} else {
			fmt.Println("  0. Выход (Esc)")
		}

		_, key, err := keyboard.GetSingleKey()
		if err != nil {
			continue
		}

		switch key {
		case keyboard.KeyArrowUp:
			selected--
			if selected < 0 {
				selected = 2
			}

		case keyboard.KeyArrowDown:
			selected++
			if selected > 2 {
				selected = 0
			}

		case keyboard.KeyEnter:
			switch selected {
			case 1:
				fmt.Println("Настройка SSH...")
				return
			case 2:
				fmt.Println("Изменение SSH ссылки...")
				return
			case 0:
				fmt.Println("Выход...")
				return
			}
		case keyboard.KeyEsc:
			fmt.Println("Выход...")
			return
		case keyboard.KeyBackspace:
			fmt.Println("Выход...")
			return
		default:
			continue
		}

	}
}
