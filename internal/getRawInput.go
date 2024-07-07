package internal

import (
	"github.com/pkg/term"
)

func GetRawInput() byte {
	t, _ := term.Open("/dev/tty")

	err := term.RawMode(t)
	if err != nil {
		panic(err)
	}

	readBytes := make([]byte, 3)
	read, err := t.Read(readBytes)
	if err != nil {
		panic(err)
	}

	t.Restore()
	t.Close()

	if read == 3 {
		return readBytes[2]
	} else {
		return readBytes[0]
	}
}
