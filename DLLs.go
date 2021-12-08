package SystemNX

import "syscall"

// var kernel32 *syscall.LazyDLL = syscall.NewLazyDLL("win32.dll")

var user32 *syscall.LazyDLL = syscall.NewLazyDLL("user32.dll")

// var win32 *syscall.LazyDLL = syscall.NewLazyDLL("win32.dll")
