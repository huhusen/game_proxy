package util

import (
	"bytes"
	"encoding/binary"
)

type BinaryReader struct {
	reader *bytes.Reader
}

func NewBinaryReader(buf []byte) *BinaryReader {
	reader := bytes.NewReader(buf)
	return &BinaryReader{reader: reader}
}
func (reader *BinaryReader) ReadBytes(len int) (buf []byte) {
	reader.check(len)
	buf = NewBuf(len)
	reader.reader.Read(buf)
	buf = bytes.TrimRight(buf, "\x00")
	return
}
func (reader *BinaryReader) ReadInt32() (i int32) {
	reader.check(4)
	binary.Read(reader.reader, binary.BigEndian, &i)
	return
}
func (reader *BinaryReader) ReadUInt32() (i uint32) {
	reader.check(4)
	binary.Read(reader.reader, binary.BigEndian, &i)
	return
}
func (reader *BinaryReader) ReadByte() (b byte) {
	reader.check(1)
	binary.Read(reader.reader, binary.BigEndian, &b)
	return
}

func (reader *BinaryReader) Position() (i int64) {
	i = reader.reader.Size() - int64(reader.reader.Len())
	return
}
func (reader *BinaryReader) Size() (i int64) {
	i = reader.reader.Size()
	return
}
func (reader *BinaryReader) Len() (i int) {
	i = reader.reader.Len()
	return
}
func (reader *BinaryReader) Seek(offset int64, whence int) {
	reader.reader.Seek(offset, whence)
	return
}
func (reader *BinaryReader) check(len int) {
	if reader.Size()-reader.Position() < int64(len) {
		panic("outside the bounds of the BinaryReader.")
	}
	return
}
