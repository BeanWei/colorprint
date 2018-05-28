package colorprint

import (
	"fmt"
	"syscall"
)

//Print 设置终端字体颜色
func Print(s string, i int) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	fmt.Print(s)
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}

const Deepblue int = 1 | 8
const Green int = 2 | 8
const Skyblue int = 3 | 8
const Red int = 4 | 8
const Purple int = 5 | 8
const Yellow int = 6 | 8
const White int = 7 | 8
