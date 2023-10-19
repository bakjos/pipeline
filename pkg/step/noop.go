package step

import (
	"context"
	"io"
)

// NoOpStep is a step that does nothing but copy from the reader to the writer.
// It's good e.g. for testing purposes where you simply want to check the functionality of the pipeline itself.
type NoOpStep struct{}

func (n NoOpStep) Invoke(ctx context.Context, reader io.Reader, writer io.Writer) error {
	_, err := io.Copy(writer, reader)
	return err
}
