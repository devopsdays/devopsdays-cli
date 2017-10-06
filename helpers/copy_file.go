package helpers

import (
	"io"
	"os"
)

// Copyfile takes a source and destination file path and copies the file. Destination needs to be complete path of file, not a directory
func CopyFile(srcPath, destPath string) (err error) {
	// open files r and w
	r, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer w.Close()

	// do the actual work
	_, err = io.Copy(w, r) // <------ here !
	if err != nil {
		return err
	}

	// fmt.Printf("Copied %v bytes\n", n)
	return nil
}
