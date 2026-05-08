package fakeerror
// improve run
import (
	"syscall"
	"unsafe"
)

func Run() {
	var title, text *uint16
	title, _ = syscall.UTF16PtrFromString("Fatal")
	text, _ = syscall.UTF16PtrFromString("Compiler failed: 0x988958")
	syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(0, uintptr(unsafe.Pointer(text)), uintptr(unsafe.Pointer(title)), 0)
}
