package step

import (
	"bytes"
	"context"
	"testing"
)

func TestNoOpStep_Invoke(t *testing.T) {
	step := NoOpStep{}
	in := []byte("noop")
	out := bytes.Buffer{}
	err := step.Invoke(context.Background(), bytes.NewReader(in), &out)
	if err != nil {
		t.Fatal(err)
	}
	if out.String() != "noop" {
		t.Fatalf("want: noop, got: %s\n", out.String())
	}
}
