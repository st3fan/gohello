
package hello_test

import (
	"testing"
	"github.com/st3fan/gohello"
)


func Test_Hello(t *testing.T) {
	if msg, _ := hello.Hello("Gopher", "en"); msg != "Hello, Gopher!" {
		t.Error("Expected 'Hello, Gopher!'")
	}
}

func Test_HelloAnonymous(t *testing.T) {
	if msg, _ := hello.Hello("", "en"); msg != "Hello, anonymous!" {
		t.Error("Expected 'Hello, anonymous!'")
	}
}

func Test_HelloWithUnknownLocale(t *testing.T) {
	if _, err := hello.Hello("Gopher", "nl"); err == nil {
		t.Error("Expected error")
	}
}

func Test_HelloAnonymousWithUnknownLocale(t *testing.T) {
	if _, err := hello.Hello("", "nl"); err == nil {
		t.Error("Expected error")
	}
}
