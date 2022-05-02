package identifiers

import (
	"fmt"
	// "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAppID(t *testing.T) {
	fmt.Println(time.Now())
	generateAndPrintRID[App]()
	generateAndPrintRID[Trace]()
}

func generateAndPrintRID[T ResourceType]() {
	fmt.Println(New[T]().String())
}
