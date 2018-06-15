package sayings

import "fmt"

// Greet returns a greeting message
func Greet(s string) string {
	return fmt.Sprint("Greet ", s)
}
