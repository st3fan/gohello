
package hello

// Hello returns a friendly hello message.

func Hello(name string) string {
	if name == "" {
		name = "anonymous"
	}
	return "Hello, " + name + "!"
}
