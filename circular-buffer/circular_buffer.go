package circular

import "errors"

type Buffer struct {
	slots []byte
	maxSize int
}

func NewBuffer(size int) *Buffer {
	var b = Buffer {[]byte{}, size}
	return &b
}

func (b *Buffer) ReadByte() (byte, error) {
	if len(b.slots) == 0 {
		return '0', errors.New("read failed")
	}
	var result = b.slots[0]
	b.slots = b.slots[1:]
	return result, nil
}

func (b *Buffer) WriteByte(c byte) error {
	b.slots = append(b.slots, c)
	return nil
}

func (b *Buffer) Overwrite(c byte) {

}

func (b *Buffer) Reset() {

}
