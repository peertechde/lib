package ioutil

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// todo:
// AtomicWriteFile
func AtomicWriteFile(fname string, data []byte, perm os.FileMode) error {
	fd, err := ioutil.TempFile(filepath.Dir(fname), ".tmp-"+filepath.Base(fname))
	if err != nil {
		return err
	}
	if err = os.Chmod(fd.Name(), perm); err != nil {
		fd.Close()
		return err
	}
	n, err := fd.Write(data)
	if err == nil && n < len(data) {
		fd.Close()
		return io.ErrShortWrite
	}
	if err != nil {
		fd.Close()
		return err
	}
	if err := fd.Sync(); err != nil {
		fd.Close()
		return err
	}
	if err := fd.Close(); err != nil {
		return err
	}
	return os.Rename(fd.Name(), fname)
}


