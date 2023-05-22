package helpers

import (
	"encoding/base64"
	"log"

	"github.com/skip2/go-qrcode"
)

func QRConvertToBase64String(qrText string) string {
	var png []byte
	png, err := qrcode.Encode(qrText, qrcode.Highest, 256)
	if err != nil {
		log.Fatal(err)
	}

	imgStr := base64.StdEncoding.EncodeToString(png)
	imgBase64Str := "data:image/png;base64," + imgStr

	return imgBase64Str
}
