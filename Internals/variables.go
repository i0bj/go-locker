package variables

import (
	"os"
	"os/user"
	"path/filepath"
	"syscall"
)

// Variables that holds directory and user info
var (
	home = os.UserHomeDir()
	user = user.Current()
	desktop = 

)
