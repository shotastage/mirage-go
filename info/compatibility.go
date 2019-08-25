package info

import "runtime"

func compatibility() {

	switch runtime.GOOS {
	case "windows":
		println("❌  Platform Windows does not be supported.")
	case "darwin":
		println("✅  Platform macOS is compatible.")
	case "linux":
		println("✅  Platform Linux is compatible.")
	case "freebsd":
		println("✅  Platform FreeBSD is compatible.")
	default:
		println("❓  Failed to detect platfrom.")
	}
}
