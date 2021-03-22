package slicewriteseek

import "io"

// SliceWriteSeeker implements WriteSeeker on a slice
type SliceWriteSeeker struct {
	Buffer []byte
	Index  int64
}

// New creates a new SliceWriteReader
func New() *SliceWriteSeeker {
	return &SliceWriteSeeker{Buffer: []byte{}}
}

// Len returns the length of the underlying slice
func (sws *SliceWriteSeeker) Len() int64 {
	return int64(len(sws.Buffer))
}

func (sws *SliceWriteSeeker) Read(p []byte) (n int, err error) {
	if sws.Index == sws.Len() {
		return 0, io.EOF
	}
	readTo := sws.Index + int64(len(p))
	if readTo > sws.Len() {
		readTo = sws.Len()
	}
	n = int(readTo - sws.Index)
	copy(p, sws.Buffer[sws.Index:readTo])
	sws.Index = readTo
	return
}

func (sws *SliceWriteSeeker) Write(p []byte) (n int, err error) {
	writeLen := int64(len(p))
	switch {
	case sws.Len() == 0:
		sws.Buffer = p
		sws.Index = int64(len(p))
	case sws.Index == sws.Len():
		sws.Buffer = append(sws.Buffer, p...)
		sws.Index += writeLen
	case sws.Index < sws.Len():
		switch {
		case sws.Index+writeLen > sws.Len():
			sws.Buffer = append(sws.Buffer[:sws.Index], p...)
		case sws.Index+writeLen <= sws.Len():
			sws.Buffer = append(sws.Buffer[:sws.Index], append(p, sws.Buffer[sws.Index+writeLen:]...)...)
		}
		sws.Index += writeLen
	}
	return len(p), err
}

// Seek sets the offset for the next Read or Write to offset, interpreted according to whence
func (sws *SliceWriteSeeker) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		sws.Index = 0 + offset
	case io.SeekCurrent:
		sws.Index = sws.Index + offset
	case io.SeekEnd:
		sws.Index = sws.Len() + offset
	}
	return sws.Index, nil
}
