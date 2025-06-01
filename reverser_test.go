package reverser

import (
	"strings"
	"testing"
)

func TestProcessTexts(t *testing.T) {
	input := []string{"Hello", "qwerty", "Golang", "platypus", "тест", "level", "generics"}
	expectedReversed := []string{"olleH", "ytrewq", "gnaloG", "supytalp", "тсет", "level", "scireneg"}

	output := ProcessTexts(input, 3)

	if len(output) != len(input)+1 {
		t.Fatalf("expected %d lines, got %d", len(input)+1, len(output))
	}

	for i := 0; i < len(input); i++ {
		if !strings.HasPrefix(output[i], "line ") {
			t.Errorf("line %d missing prefix: %s", i+1, output[i])
		}
		if !strings.Contains(output[i], expectedReversed[i]) {
			t.Errorf("expected reversed %q in line %d, got: %s", expectedReversed[i], i+1, output[i])
		}
	}

	if output[len(output)-1] != "done" {
		t.Errorf("last line should be 'done', got: %s", output[len(output)-1])
	}
}
