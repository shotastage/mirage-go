package os

import (
	"runtime"
)

func CheckPlatform() {
	print("😃  Your platform is ")

	switch runtime.GOOS {
	case "windows":
		print("Windows")
	case "darwin":
		print("macOS")
	case "linux":
		print("Linux")
	case "freebsd":
		print("FreeBSD")
	default:
		print("Other OS")
	}

	print(".\n")

	return
}
