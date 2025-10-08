package fixes

import (
	"os/exec"
	"runtime"
)

// FixLanguageSwitch исправляет проблему с переключением языка в Windows
func FixLanguageSwitch() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("rundll32", "user32.dll,SetFocus", "0")
		cmd.Run()
	}
}
