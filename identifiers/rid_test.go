package identifiers

import (
	"fmt"
	// "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func ExampleAppID(t *testing.T) {
	generateAndPrintRID[App]()
	generateAndPrintRID[Trace]()
}

func generateAndPrintRID[T ResourceType]() {
	fmt.Println(New[T]().String())
}
