package ioutil

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestAtomicWriteFile(t *testing.T) {
	var (
		content = []byte("atomic write file test")
		mode    = os.FileMode(0777)
	)

	tmpDirectory, err := ioutil.TempDir("", "atomic-write-file-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDirectory)

	fname := tmpDirectory + "/" + "test"
	if err := AtomicWriteFile(fname, content, mode); err != nil {
		t.Fatal(err)
	}
	actual, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(content, actual) {
		t.Fatalf("expected %q data, got %q", content, actual)
	}
	fi, err := os.Stat(fname)
	if err != nil {
		t.Fatal(err)
	}
	if fi.Mode() != mode {
		t.Fatalf("expected %o mode, got %o", mode, fi.Mode())
	}
}
