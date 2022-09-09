package binaryext

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
)

type ByteArray struct {
	buf    bytes.Buffer
	Endian string
}

func NewByteArray(param ...interface{}) *ByteArray {
	obj := &ByteArray{Endian: "big"}
	if len(param) > 0 {
		for _, o := range param {
			switch o.(type) {
			case string:
				obj.Endian = o.(string)
			case []byte:
				buf := o.([]byte)
				obj.Wrap(buf)
			}
		}
		return obj
	}
	return &ByteArray{Endian: "big"}
}

func (obj *ByteArray) Wrap(data []byte) {
	obj.buf.Write(data)
}

func (obj *ByteArray) WriteShort(value int) {
	var buff = make([]byte, 2)
	if obj.Endian == "big" {
		binary.BigEndian.PutUint16(buff, uint16(value))
	} else if obj.Endian == "little" {
		binary.LittleEndian.PutUint16(buff, uint16(value))
	}
	obj.buf.Write(buff)
}

func (obj *ByteArray) ReadShort() []byte {
	var tempBuff = obj.buf.Bytes()
	var shortValue = tempBuff[:2]
	var restValue = tempBuff[2:]
	var byteByteArray bytes.Buffer
	byteByteArray.Write(restValue)
	obj.buf = byteByteArray
	return shortValue
}

func (obj *ByteArray) WriteInt(value int) {
	var buff = make([]byte, 4)
	if obj.Endian == "big" {
		binary.BigEndian.PutUint32(buff, uint32(value))
	} else if obj.Endian == "little" {
		binary.LittleEndian.PutUint32(buff, uint32(value))
	}
	obj.buf.Write(buff)
}

func (obj *ByteArray) ReadInt() []byte {
	var tempBuff = obj.buf.Bytes()
	var intValue = tempBuff[:4]
	var restValue = tempBuff[4:]
	var byteByteArray bytes.Buffer
	byteByteArray.Write(restValue)
	obj.buf = byteByteArray
	return intValue

}
func (obj *ByteArray) WriteLong(value int) {
	var buff = make([]byte, 8)
	if obj.Endian == "big" {
		binary.BigEndian.PutUint64(buff, uint64(value))
	} else if obj.Endian == "little" {
		binary.LittleEndian.PutUint64(buff, uint64(value))
	}
	obj.buf.Write(buff)
}

func (obj *ByteArray) ReadLong() []byte {
	var tempBuff = obj.buf.Bytes()
	var longValue = tempBuff[:8]
	var restValue = tempBuff[8:]
	var byteByteArray bytes.Buffer
	byteByteArray.Write(restValue)
	obj.buf = byteByteArray
	return longValue
}

func (obj *ByteArray) WriteFloat(value float32) {
	var bits = math.Float32bits(value)
	var buff = make([]byte, 4)
	if obj.Endian == "big" {
		binary.BigEndian.PutUint32(buff, bits)
	} else if obj.Endian == "little" {
		binary.LittleEndian.PutUint32(buff, bits)
	}
	obj.buf.Write(buff)
}

func (obj *ByteArray) ReadFloat() []byte {
	var tempBuff = obj.buf.Bytes()
	var floatValue = tempBuff[:4]
	var restValue = tempBuff[4:]
	var byteByteArray bytes.Buffer
	byteByteArray.Write(restValue)
	obj.buf = byteByteArray
	return floatValue
}

func (obj *ByteArray) WriteDouble(value float64) {
	var bits = math.Float64bits(value)
	var buff = make([]byte, 8)
	if obj.Endian == "big" {
		binary.BigEndian.PutUint64(buff, bits)
	} else if obj.Endian == "little" {
		binary.LittleEndian.PutUint64(buff, bits)
	}
	obj.buf.Write(buff)
}

func (obj *ByteArray) ReadDouble() []byte {
	var tempBuff = obj.buf.Bytes()
	var doubleValue = tempBuff[:8]
	var restValue = tempBuff[8:]
	var byteByteArray bytes.Buffer
	byteByteArray.Write(restValue)
	obj.buf = byteByteArray
	return doubleValue
}

func (obj *ByteArray) Write(value []byte) {
	obj.buf.Write(value)
}

func (obj *ByteArray) WriteByte(value byte) {
	var tempByte []byte
	tempByte = append(tempByte, value)
	obj.buf.Write(tempByte)
}

func (obj *ByteArray) Read(size int) []byte {
	var tempBuff = obj.buf.Bytes()
	var value = tempBuff[:size]
	var restValue = tempBuff[size:]
	var byteByteArray bytes.Buffer
	byteByteArray.Write(restValue)
	obj.buf = byteByteArray
	return value
}

func (obj *ByteArray) ReadByte() []byte {
	var tempBuff = obj.buf.Bytes()
	var value = tempBuff[:1]
	var restValue = tempBuff[1:]
	var byteByteArray bytes.Buffer
	byteByteArray.Write(restValue)
	obj.buf = byteByteArray
	return value
}

func (obj *ByteArray) Data() []byte {
	return obj.buf.Bytes()
}

func (obj *ByteArray) Size() int {
	return len(obj.buf.Bytes())
}

func (obj *ByteArray) Flip() {
	var bytesArr = obj.buf.Bytes()
	for i, j := 0, len(bytesArr)-1; i < j; i, j = i+1, j-1 {
		bytesArr[i], bytesArr[j] = bytesArr[j], bytesArr[i]
	}
	var byteByteArray bytes.Buffer
	byteByteArray.Write(bytesArr)
	obj.buf = byteByteArray
}

func (obj *ByteArray) Clear() {
	var byteByteArray bytes.Buffer
	obj.buf = byteByteArray
}

func (obj *ByteArray) Slice(start int, end int) error {
	var bytesArr = obj.buf.Bytes()
	if len(bytesArr) < (start + end) {
		return errors.New("ByteArray does not contain that much of limit")
	}
	bytesArr = bytesArr[start:end]
	var byteByteArray bytes.Buffer
	byteByteArray.Write(bytesArr)
	obj.buf = byteByteArray
	return nil
}

func (obj *ByteArray) GetString(len int) string {
	return string(bytes.TrimRight(obj.Read(len), "\x00"))
}

func (obj *ByteArray) GetShort() uint16 {
	if obj.Endian == "big" {
		return binary.BigEndian.Uint16(obj.ReadShort())
	} else if obj.Endian == "little" {
		return binary.LittleEndian.Uint16(obj.ReadShort())
	}
	return 0
}

func (obj *ByteArray) GetInt() uint32 {
	if obj.Endian == "big" {
		return binary.BigEndian.Uint32(obj.ReadInt())
	} else if obj.Endian == "little" {
		return binary.LittleEndian.Uint32(obj.ReadInt())
	}
	return 0
}

func (obj *ByteArray) GetLong() uint64 {
	if obj.Endian == "big" {
		return binary.BigEndian.Uint64(obj.ReadLong())
	} else if obj.Endian == "little" {
		return binary.LittleEndian.Uint64(obj.ReadLong())
	}
	return 0
}
