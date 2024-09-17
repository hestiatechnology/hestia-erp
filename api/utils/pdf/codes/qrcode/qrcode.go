package main

import (
	"fmt"
	"log"

	"github.com/yeqown/go-qrcode/v2"
)

func main() {
	repo()

	//issue17()
}

func repo() {
	qrc, err := qrcode.NewWith("A:500000000*B:123456789*C:PT*D:GT*E:N*F:20190720*G:GTG234CB/50987*H:GTVX4Y8B-50987*I1:0*N:0.00*O:0.00*Q:5uIg*R:9999",
		qrcode.WithVersion(25), qrcode.WithEncodingMode(qrcode.EncModeByte), qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionMedium))
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return
	}

	log.Println(qrc.Dimension())

	/*
		 	writer := new(io.WriteCloser)
			w := standard.NewWithWriter(writer)

			if err := qrc.Save(w); err != nil {
				panic(err)
			}
	*/
}
