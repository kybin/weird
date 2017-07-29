package stringnode

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	tmpf, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal("could not create temp file:", err)
	}
	defer os.Remove(tmpf.Name())

	content := []byte("hi, there.\n this is a test for Read")
	if _, err := tmpf.Write(content); err != nil {
		t.Fatal("could not write content to temp file:", err)
	}
	if err := tmpf.Close(); err != nil {
		t.Fatal("could not close temp file:", err)
	}

	n := NewRead(ReadParm{
		fpath: tmpf.Name(),
	})

	want := string(content)
	got, err := n.Result()
	if err != nil {
		t.Fatalf("Result(): has unexpected error: %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("Result(): length of got should 1 but it's %v", len(got))
	}
	if got[0] != want {
		t.Fatalf("Result(): got %v, want %v", got, want)
	}
}
