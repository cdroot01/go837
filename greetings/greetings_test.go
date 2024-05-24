package greetings

import (
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "cong"

	msg, err := Hello(name)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(msg)
}

func TestEmpty(t *testing.T) {
	name := ""

	msg, err := Hello(name)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(msg)
}
