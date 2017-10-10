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
	toRead := sws.Index + int64(len(p))
	switch {
	case sws.Index+1 == sws.Len():
		p = []byte{}
	case toRead <= sws.Len():
		p = sws.Buffer[sws.Index : int(sws.Index)+len(p)]
	case toRead > sws.Len():
		p = sws.Buffer[sws.Index:]
	}
	n = len(p)
	sws.Index += int64(len(p))
	return
}

func (sws *SliceWriteSeeker) Write(p []byte) (n int, err error) {
	writeLen := int64(len(p))
	switch {
	case sws.Len() == 0:
		sws.Buffer = p
		sws.Index = int64(len(p)) - 1
	case sws.Index+1 == sws.Len():
		sws.Buffer = append(sws.Buffer, p...)
		sws.Index += writeLen
	case sws.Index+1 < sws.Len():
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
		sws.Index = (sws.Len() - 1) + offset
	}
	return sws.Index, nil
}
