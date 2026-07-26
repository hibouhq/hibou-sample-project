// Command app is a trivial entrypoint so the Go module produces a buildable
// binary (useful for SBOM generation). It only exercises the clean calc library.
package main

import (
	"fmt"

	calc "github.com/hibouhq/hibou-sample-project/go"
)

func main() {
	fmt.Println("2 + 3 =", calc.Add(2, 3))
	fmt.Println("classify(-1) =", calc.Classify(-1))
}
