package tap

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	Byte     = 1
	Kilobyte = 1024 * Byte
	Megabyte = 1024 * Kilobyte
)

type Tap struct {
	file        *os.File
	maxByteSize int64
}

func MakeTap(filename string) (Tap, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		return Tap{}, fmt.Errorf("could not tap file: %s", err)
	}
	return MakeTapFromRaw(file, 2*Megabyte), nil
}

func MakeTapFromRaw(file *os.File, maxByteSize int64) Tap {
	return Tap{file, maxByteSize}
}

func (t *Tap) Close() error {
	return t.Close()
}

func (t *Tap) NewReader() io.Reader {
	return t.file
}

// func (t *Tap) ReadAll() ([]byte, error) {
// 	err := checkSize(t.file, t.maxByteSize)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return io.ReadAll(t.file)
// }

func (t *Tap) Write(bytes []byte) (int, error) {
	return t.file.Write(bytes)
}

func checkSize(
	file *os.File,
	maxByteSize int64,
) error {
	if file == nil {
		return errors.New("file pointer is nil")
	}
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.Size() > maxByteSize {
		return fmt.Errorf(
			"Refusing to open file of size: %d", info.Size())
	}
	return nil
}
