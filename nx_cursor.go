package SystemNX

import (
	"unsafe"
)

func GetCursorPosition(Position *CursorPosition) error {
	_, ret, err := user32.NewProc("GetCursorPos").Call(uintptr(unsafe.Pointer(Position)))
	if ret != 0 {
		return err
	}
	return nil

}

func SetCursorPosition(Position *CursorPosition) error {
	_, ret, err := user32.NewProc("SetCursorPos").Call(uintptr(Position.X), uintptr(Position.Y))
	if ret != 0 {
		return err
	}
	return nil
}
