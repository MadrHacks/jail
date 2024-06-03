package config

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

const procMountFlags = uintptr(unix.MS_NOSUID | unix.MS_NOEXEC | unix.MS_RELATIME)

func MountProc() error {
	if _, err := os.Stat("/jail/proc"); errors.Is(err, os.ErrNotExist) {
		return nil
	}

	if err := unix.Mount("", "/jail/proc", "proc", procMountFlags, ""); err != nil {
		return fmt.Errorf("mount procfs: %w", err)
	}
	return nil
}
