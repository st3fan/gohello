
package hello

func Hello(name string) string {
	if name == "" {
		name = "anonymous"
	}
	return "Hello, " + name + "!"
}
