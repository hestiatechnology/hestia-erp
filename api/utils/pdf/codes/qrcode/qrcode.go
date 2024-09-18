package qrcode

import (
	"bytes"
	"fmt"
	"io"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

/* func main() {
	qr()

	//issue17()
} */

// CustomReadWriteCloser wraps a bytes.Buffer and implements io.ReadWriteCloser
type customReadWriteCloser struct {
	bytes.Buffer
}

// Close is a no-op for CustomReadWriteCloser
func (c *customReadWriteCloser) Close() error {
	return nil
}

func Qr() []byte {
	qrc, err := qrcode.NewWith("A:500000000*B:123456789*C:PT*D:GT*E:N*F:20190720*G:GTG234CB/50987*H:GTVX4Y8B-50987*I1:0*N:0.00*O:0.00*Q:5uIg*R:9999",
		/* qrcode.WithVersion(25), */ qrcode.WithEncodingMode(qrcode.EncModeByte), qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionMedium))
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return nil
	}

	// Initialize buf with CustomReadWriteCloser
	var buf io.ReadWriteCloser = &customReadWriteCloser{}
	w := standard.NewWithWriter(buf)

	if err := qrc.Save(w); err != nil {
		panic(err)
	}

	// Now you can use the buffer
	return buf.(*customReadWriteCloser).Bytes()
}
