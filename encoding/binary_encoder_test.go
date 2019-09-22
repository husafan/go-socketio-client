package encoding_test

import (
	"io/ioutil"
	"strings"
	"testing"

	. "github.com/husafan/go-socketio-client/encoding"
	. "github.com/smartystreets/goconvey/convey"
)

// Tests that the BinaryLengthEncoder prepends the length bytes to the
// given []byte.
func TestBinaryLengthEncoder(t *testing.T) {
	smallData := strings.NewReader("abcde")
	encoder := NewBinaryLengthEncoder(smallData)
	encodedBytes, err := ioutil.ReadAll(encoder)
	Convey("Calling BinaryLengthEncoder on a short 'abcde'", t, func() {
		Convey("should not result in an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("should return 8 encodedBytes", func() {
			So(encodedBytes, ShouldHaveLength, 8)
		})
		Convey("should have 1 as the first byte", func() {
			So(encodedBytes[0], ShouldEqual, 1)
		})
		Convey("should have 5 as the second byte", func() {
			So(encodedBytes[1], ShouldEqual, 5)
		})
		Convey("should have 255 as the third byte", func() {
			So(encodedBytes[2], ShouldEqual, 255)
		})
		Convey("should end with the original bytes", func() {
			So(encodedBytes[3:], ShouldResemble, []byte("abcde"))
		})

	})
	largeData := strings.Repeat("a", 45)
	encoder = NewBinaryLengthEncoder(strings.NewReader(largeData))
	encodedBytes, err = ioutil.ReadAll(encoder)
	Convey("Calling BinaryLengthEncoder on a 45 char string", t, func() {
		Convey("should not result in an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("should return 49 encodedBytes", func() {
			So(encodedBytes, ShouldHaveLength, 49)
		})
		Convey("should have 1 as the first byte", func() {
			So(encodedBytes[0], ShouldEqual, 1)
		})
		Convey("should have 4 as the second byte", func() {
			So(encodedBytes[1], ShouldEqual, 4)
		})
		Convey("should have 5 as the third byte", func() {
			So(encodedBytes[2], ShouldEqual, 5)
		})
		Convey("should have 255 as the fourth byte", func() {
			So(encodedBytes[3], ShouldEqual, 255)
		})
		Convey("should end with the original bytes", func() {
			So(encodedBytes[4:], ShouldResemble, []byte(largeData))
		})
	})
}
