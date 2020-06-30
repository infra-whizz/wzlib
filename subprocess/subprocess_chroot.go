package wzlib_subprocess

import (
	"os"
	"syscall"
)

/*
Chroot returns a function that escapes it back.

Example usage:

	exit, _ := Chroot("/path/to/the/new/root")

	// call stuff

	exit()
*/

func Chroot(path string) (func() error, error) {
	root, err := os.Open("/")
	if err != nil {
		return nil, err
	}

	if err := syscall.Chroot(path); err != nil {
		root.Close()
		return nil, err
	}

	return func() error {
		defer root.Close()
		if err := root.Chdir(); err != nil {
			return err
		}
		return syscall.Chroot(".")
	}, nil
}
