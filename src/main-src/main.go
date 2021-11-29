package yayx

import "os"

func main() {
	argv := os.Args
	argvLen := len(argv)

	Use(argvLen)
}
