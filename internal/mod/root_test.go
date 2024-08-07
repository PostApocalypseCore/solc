package mod_test

import (
	"testing"

	"github.com/PostApocalypseCore/solc/internal/mod"
)

func TestModRoot(t *testing.T) {
	// if !strings.HasSuffix(mod.Root, "solc") {
	if mod.Root == "" {
		t.Fatalf("Unexpected module root: %q", mod.Root)
	}
}
