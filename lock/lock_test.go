package lock

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	file, err := ioutil.TempFile("", "lock-test")
	if err != nil {
		t.Fatal(err)
	}
	file.Close()
	defer func() {
		if err := os.Remove(file.Name()); err != nil {
			t.Fatal(err)
		}
	}()

	// lock file
	lock := New(file.Name())
	if err := lock.Lock(); err != nil {
		t.Fatal(err)
	}

	// try to lock a locked file
	dupl := New(file.Name())
	if err := dupl.TryLock(); err != ErrLockLocked {
		t.Fatal(err)
	}
	if err := lock.Unlock(); err != nil {
		t.Fatal(err)
	}

	// lock duplicate
	if err := dupl.Lock(); err != nil {
		t.Fatal(err)
	}

	// block while `dupl` is locked
	locked := make(chan bool, 1)
	go func() {
		blocker := New(file.Name())
		if err := blocker.Lock(); err != nil {
			t.Fatal(err)
		}
		locked <- true
		if err := blocker.Unlock(); err != nil {
			t.Fatal(err)
		}
	}()

	select {
	case <-locked:
		t.Errorf("blocker didn't block")
	case <-time.After(1 * time.Second):
	}

	// unlock dupl to unblock
	if err = dupl.Unlock(); err != nil {
		t.Fatal(err)
	}

	// blocker should be unblocked
	select {
	case <-locked:
	case <-time.After(2 * time.Second):
		t.Errorf("blocker didn't unblock")
	}
}
