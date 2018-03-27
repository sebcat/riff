# riff
--
    import "github.com/sebcat/riff"


## Usage

#### type Chunk

```go
type Chunk struct {
	Identifier [4]byte
	Offset     int64
	Size       uint32
}
```


#### type Reader

```go
type Reader struct {
}
```


#### func  NewReader

```go
func NewReader(r io.ReadSeeker) *Reader
```

#### func (*Reader) NextChunk

```go
func (r *Reader) NextChunk(c *Chunk) error
```
