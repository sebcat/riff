package riff

import (
	"encoding/binary"
	"io"
)

type Reader struct {
	r io.ReadSeeker
}

type Chunk struct {
	Identifier [4]byte
	Offset     int64
	Size       uint32
}

func NewReader(r io.ReadSeeker) *Reader {
	return &Reader{r: r}
}

func (r *Reader) NextChunk(c *Chunk) error {
	var (
		id     [4]byte
		size   uint32
		offset int64
	)

	_, err := io.ReadFull(r.r, id[:])
	if err != nil {
		return err
	}

	if err := binary.Read(r.r, binary.LittleEndian, &size); err != nil {
		return err
	}

	if offset, err = r.r.Seek(0, io.SeekCurrent); err != nil {
		return err
	}

	if _, err = r.r.Seek(int64(size+1&^1), io.SeekCurrent); err != nil {
		return err
	}

	c.Identifier = id
	c.Offset = offset
	c.Size = size
	return nil
}
