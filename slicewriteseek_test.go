package slicewriteseek

import (
	"errors"
	"io"
	"testing"
)

func TestWrite(t *testing.T) {
	s := New()
	if _, err := s.Write([]byte{1, 2, 4}); err != nil {
		t.Error(err)
	}
}

func TestWriteMany(t *testing.T) {
	s := New()
	if _, err := s.Write([]byte{1, 2, 3}); err != nil {
		t.Error(err)
	}
	if _, err := s.Write([]byte{4, 5, 6}); err != nil {
		t.Error(err)
	}
	if s.Len() != 6 {
		t.Errorf("Expecting Len == 6, got %d", s.Len())
	}
}

func TestLen(t *testing.T) {
	s := New()
	if _, err := s.Write([]byte{1, 2, 4}); err != nil {
		t.Error(err)
	}
	if s.Len() != 3 {
		t.Errorf("Expecting Len == 3, got %d", s.Len())
	}
}

func TestSeek(t *testing.T) {
	s := New()
	if off, err := s.Seek(0, io.SeekEnd); err != nil || off != 0 {
		t.Error("Expecting offset to be zero at the end of an empty slice")
	}
	if _, err := s.Write([]byte{1, 2, 4}); err != nil {
		t.Error(err)
	}
	if off, err := s.Seek(0, io.SeekCurrent); err != nil || off != 3 {
		t.Error("Unexpected seek")
	}
	if off, err := s.Seek(0, io.SeekStart); err != nil || off != 0 {
		t.Error("Unexpected seek")
	}
	if s.Buffer[s.Index] != 1 {
		t.Errorf("Expecting first item to be 1, got %v", s.Buffer[s.Index])
	}
	if off, err := s.Seek(0, io.SeekEnd); err != nil || off != 3 {
		t.Error("Unexpected seek")
	}
	if s.Buffer[s.Index-1] != 4 {
		t.Errorf("Expecting last item to be 4, got %v", s.Buffer[s.Index-1])
	}
	s.Index = 1
	if off, err := s.Seek(0, io.SeekCurrent); err != nil || off != 1 {
		t.Error("Unexpected seek")
	}
	if s.Buffer[s.Index] != 2 {
		t.Errorf("Expecting item to be 2, got %v", s.Buffer[s.Index])
	}
}

func TestRead(t *testing.T) {
	s := New()
	if _, err := s.Write([]byte{1, 2, 4}); err != nil {
		t.Error(err)
	}
	if _, err := s.Seek(0, io.SeekEnd); err != nil {
		t.Error(err)
	}
	p := make([]byte, 1)
	n, err := s.Read(p)
	if !errors.Is(err, io.EOF) {
		t.Errorf("Expecting error io.EOF, got %v", err)
	}
	if n != 0 {
		t.Errorf("Expecting to get back an empty slice, got %d", n)
	}
	if _, err = s.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
	}
	p = make([]byte, 2)
	n, err = s.Read(p)
	if err != nil {
		t.Error(err)
	}
	if n != 2 {
		t.Errorf("Expecting to get back an len 2 slice, got %d", n)
	}
	if _, err = s.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
	}
	p = make([]byte, 4)
	n, err = s.Read(p)
	if err != nil {
		t.Error(err)
	}
	if n != 3 {
		t.Errorf("Expecting to get back an len 3 slice, got: %d", n)
	}
}

func TestWriteAt(t *testing.T) {
	s := New()
	if _, err := s.Write([]byte{1, 2, 4}); err != nil {
		t.Error(err)
	}
	if _, err := s.Seek(1, io.SeekStart); err != nil {
		t.Error(err)
	}
	s.Write([]byte{5})
	if s.Len() != 3 {
		t.Errorf("Expecting Len == 3, got %d", s.Len())
	}
	if s.Buffer[1] != 5 {
		t.Errorf("Expecting 5, got %v", s.Buffer)
	}
	if _, err := s.Seek(0, io.SeekEnd); err != nil {
		t.Error(err)
	}
	s.Write([]byte{1})
	if s.Len() != 4 {
		t.Errorf("Expecting Len == 4, got %d", s.Len())
	}
	if _, err := s.Seek(1, io.SeekStart); err != nil {
		t.Error(err)
	}
	s.Write([]byte{2, 3, 4, 5})
	if s.Len() != 5 {
		t.Errorf("Expecting Len == 5, got %d", s.Len())
	}
	if s.Buffer[4] != 5 {
		t.Errorf("Expecting 5, got %v", s.Buffer[4])
	}
}
