
package hello_test

import (
	"testing"
	"github.com/st3fan/gohello"
)


func Test_Hello(t *testing.T) {
	if hello.Hello("Gopher") != "Hello, Gopher!" {
		t.Error("Expected 'Hello, Gopher!'")
	}
}

func Test_HelloAnonymous(t *testing.T) {
	if hello.Hello("") != "Hello, anonymous!" {
		t.Error("Expected 'Hello, anonymous!'")
	}
}

func Test_Unicode(t *testing.T) {
	if hello.Hello("Amélie") != "Hello, Amélie!" {
		t.Error("Expected 'Hello, Amélie!'")
	}
}

