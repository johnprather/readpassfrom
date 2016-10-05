package readpassfrom

import (
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// ReadPasswordFrom reads a password from a suitable file
//
// file - the file to read and write to (like /dev/ttys)
// prompt - string displayed to user as password prompt
//
func ReadPasswordFrom(file string, prompt string) (*string, error) {
	tty, err := os.OpenFile(file, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer tty.Close()

	tty.Write([]byte(prompt))
	fd := tty.Fd()
	pass, err := terminal.ReadPassword(int(fd))
	tty.Write([]byte("\n"))
	if err != nil {
		return nil, err
	}
	passwordstr := string(pass)
	return &passwordstr, nil
}
