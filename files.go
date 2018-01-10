package files

import (
    "os"
    "io"
)

func Copy(src, dst string) error {
	srcHandle, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcHandle.Close()

	dstHandle, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstHandle.Close()

	if _, err = io.Copy(dstHandle, srcHandle); err != nil {
		return err
	}

	err = dstHandle.Sync()
	return err
}
