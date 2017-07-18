
package hello

// Hello returns a friendly hello message.

const DefaultName = "Anonymoose"

func Hello(name string) string {
	if name == "" {
		name = DefaultName
	}
	return "Hello, " + name + "!"
}
