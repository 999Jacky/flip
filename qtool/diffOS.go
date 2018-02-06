package qtool

// https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
import (
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func() // create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) // Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") // Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") // Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] // runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          // if we defined a clear func for that platform:
		value() // we execute it
	} else { // unsupported platform
		panic("Your platform is unsupported! ")
	}
}

// open opens the specified URL in the default browser of the user.
// https://stackoverflow.com/questions/39320371/how-start-web-server-to-open-page-in-browser-in-golang
func Open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
