package encoding

import (
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
)

// Calculates the length prefix bytes given the length of the original
// data bytes.
func getLengthBytes(length int) ([]byte, error) {
	strLength := strconv.FormatInt(int64(length), 10)
	lengthBytes := make([]byte, len(strLength)+2)
	lengthBytes[0] = 1
	for index, val := range strLength {
		if byteVal, err := strconv.Atoi(string(val)); err == nil {
			lengthBytes[index+1] = byte(byteVal)
		} else {
			return []byte{}, err
		}
	}
	lengthBytes[len(lengthBytes)-1] = 255
	return lengthBytes, nil
}

// The binary length encoder is responsible for adding the encoded
// length bytes to the front of an array of bytes. The JS reference
// implementation is here:
// https://github.com/socketio/engine.io-parser/blob/master/lib/index.js#L386
//
// The steps are:
// 1. Get the length of the data to send as a string.
// 2. Create a byte array, b, equal to the length of that string +2.
// 3. Set the first byte to 1 to indicate that the data is binary.
// 4. For bytes 1 through len(b) - 1, set the byte equal to the
//    corresponding length string digit value.
// 5. Set the last byte to 255
type BinaryLengthEncoder struct {
	toEncode  io.Reader
	readIndex int
}

// Reads the prefixed bytes passed to the BinaryLengthEncoder.
func (ble *BinaryLengthEncoder) Read(toFill []byte) (int, error) {
	if ble.toEncode == nil {
		return 0, io.EOF
	}
	return ble.toEncode.Read(toFill)
}

// Reads the bytes from reader and calculates the correct length
// prefix. The length prefix is combined with the original bytes and
// can be read out via the Read method.
func (ble *BinaryLengthEncoder) ReadFrom(reader io.Reader) (bytesRead int64, err error) {
	packetBytes, err := ioutil.ReadAll(reader)
	if err == nil {
		if lengthBytes, err := getLengthBytes(len(packetBytes)); err == nil {
			ble.toEncode = bytes.NewBuffer(append(lengthBytes, packetBytes...))
		}
	}
	return int64(len(packetBytes)), err
}

// Returns a *BinaryLengthEncoder containing the length-prefixed bytes
// read from the passed in Reader.
func NewBinaryLengthEncoder(reader io.Reader) *BinaryLengthEncoder {
	toReturn := &BinaryLengthEncoder{}
	toReturn.ReadFrom(reader)
	return toReturn
}
