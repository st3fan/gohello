
package hello

// Hello is it me you are looking for?

func Hello(name string) string {
	if name == "" {
		name = "anonymous"
	}
	return "Hello, " + name + "!"
}
