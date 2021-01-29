package tempconv

import (
	"flag"
	"testing"
)

func TestCelsiusFlag(t *testing.T) {
	var temp = CelsiusFlag("temp", 20.0, "temperature specification")

	flag.Parse()

	t.Log(temp)
}
