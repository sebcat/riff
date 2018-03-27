package riff

import (
	"os"
	"testing"
)

func TestReader(t *testing.T) {
	// TODO: table driven test with expected values
	f, err := os.Open("testdata/demo.wav")
	if err != nil {
		t.Fail()
	}

	defer f.Close()
	r := NewReader(f)
	var chunk Chunk
	for err = r.NextChunk(&chunk); err == nil; err = r.NextChunk(&chunk) {
		t.Logf("chunk: %v %v %v\n", string(chunk.Identifier[:]), chunk.Offset,
			chunk.Size)
	}
	t.Log(err)
}
