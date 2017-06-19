
package hello

import "fmt"

// Hello returns a friendly hello message in the specified locale.

func Hello(name string, locale string) (string, error) {
	if locale != "en" {
		return "", fmt.Errorf("unsupported locale: " + locale)
	}

	if name == "" {
		name = "anonymous"
	}
	return "Hello, " + name + "!", nil
}
