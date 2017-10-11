[![Build Status](https://travis-ci.org/corticph/slicewriteseek.svg?branch=master)](https://travis-ci.org/corticph/slicewriteseek)
[![GoDoc](https://godoc.org/github.com/corticph/slicewriteseek?status.svg)](https://godoc.org/github.com/corticph/slicewriteseek)
# SliceWriteSeek
SliceWriteSeeker implements WriteSeeker on a slice

### Sample Usage

```golang
s := slicewriteseek.New()
// write to it
if _, err := s.Write([]byte{1, 2, 3}); err != nil {
	panic(err)
}
// seek it
if off, err := s.Seek(0, io.SeekStart); err != nil || off != 0 {
	panic("Unexpected seek")
}
// read it
b := make([]byte, 1)
if _, err := s.Read(b); err != nil {
	panic(err)
}
```
