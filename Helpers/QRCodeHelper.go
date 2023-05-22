package helpers

import (
	"encoding/base64"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"

	model "github.com/ThaiQR/QRGenerate-Gin/Model"
	"github.com/skip2/go-qrcode"

	"github.com/howeyc/crc16"
)

func QRMerchantPromptpayGenerate(value model.MerchantPromptpayModel) string {
	var merchantIdentifier, merchantSum string

	value.PFI = preprocessValue("00", value.PFI)

	value.PIM = preprocessValue("01", value.PIM)

	/** Merchant Identifier */
	value.AID = preprocessValue("00", value.AID)
	value.RecieveID = preprocessRecieveID(value.RecieveID)

	merchantSum = value.AID + value.RecieveID
	merchantIdentifier = preprocessValue("29", merchantSum)
	/* */

	value.MerchantCode = preprocessValue("52", value.MerchantCode)

	value.Currency = preprocessValue("53", value.Currency)

	value.Amount = preprocessAmount(value.Amount)

	value.CountryCode = preprocessValue("58", value.CountryCode)

	value.MerchantCity = preprocessValue("60", value.MerchantCity)

	value.PostalCode = preprocessValue("61", value.PostalCode)

	crc := "6304"

	data := value.PFI + value.PIM + merchantIdentifier + value.MerchantCode + value.Currency + value.Amount + value.CountryCode + value.MerchantCity + value.PostalCode + crc
	dataBuffer := []byte(data)
	crcResult := crc16.ChecksumCCITTFalse(dataBuffer)

	data += strings.ToUpper(strconv.FormatInt(int64(crcResult), 16))

	return data
}

// func MerchantGenerate(pfi, pim, aid, billerID, ref01, ref02, merchantCode, currency, amount, countryCode, merchantName, terminalID string) string {
func QRMerchantBillpaymentGenerate(value model.MerchantBillpaymentModel) string {
	var merchantIdentifier, merchantSum string
	value.Reference1 = strings.ToUpper(value.Reference1)
	value.Reference2 = strings.ToUpper(value.Reference2)

	value.PFI = preprocessValue("00", value.PFI)

	value.PIM = preprocessValue("01", value.PIM)

	/** Merchant Identifier */
	value.AID = preprocessValue("00", value.AID)
	value.BillerID = preprocessValue("01", value.BillerID)
	value.Reference1 = preprocessValue("02", value.Reference1)
	value.Reference2 = preprocessValue("03", value.Reference2)

	merchantSum = value.AID + value.BillerID + value.Reference1 + value.Reference2
	// merchantIdentifier = "30" + strconv.Itoa(len(merchantSum)) + merchantSum
	merchantIdentifier = preprocessValue("30", merchantSum)
	/* */

	value.MerchantCode = preprocessValue("52", value.MerchantCode)

	value.Currency = preprocessValue("53", value.Currency)

	value.Amount = preprocessAmount(value.Amount)

	value.CountryCode = preprocessValue("58", value.CountryCode)

	value.MerchantName = preprocessValue("59", value.MerchantName)

	value.TerminalID = preprocessValue("07", value.TerminalID)

	crc := "6304"

	data := value.PFI + value.PIM + merchantIdentifier + value.MerchantCode + value.Currency + value.Amount + value.CountryCode + value.MerchantName + value.TerminalID + crc
	dataBuffer := []byte(data)
	crcResult := crc16.ChecksumCCITTFalse(dataBuffer)

	data += strings.ToUpper(strconv.FormatInt(int64(crcResult), 16))

	return data
}

func preprocessValue(prefix, value string) string {
	if len(value) != 0 {
		if len(value) < 10 {
			value = prefix + "0" + strconv.Itoa(len(value)) + value
		} else {
			value = prefix + strconv.Itoa(len(value)) + value
		}
	} else {
		value = ""
	}
	return value
}

func preprocessAmount(value string) string {
	checkAmount := strings.Split(value, ".")

	if len(checkAmount) > 1 {
		if checkAmount[1] == "" || len(checkAmount[1]) == 0 {
			checkAmount[1] = "00"
		} else if len(checkAmount[1]) == 1 {
			checkAmount[1] += "0"
		} else if len(checkAmount[1]) > 2 {
			checkAmount[1] = checkAmount[1][:2]
		}

		value = checkAmount[0] + "." + checkAmount[1]
	} else if len(checkAmount) == 1 {
		value = value + "." + "00"
	}

	value = preprocessValue("54", value)

	return value
}

func preprocessRecieveID(value string) string {
	if len(value) == 10 && value[0] == '0' { // Phone Number
		value = "0066" + trimFirstRune(value)
		value = preprocessValue("01", value)
	} else if len(value) == 13 { // National ID or Tax ID
		value = preprocessValue("02", value)
	} else if len(value) == 15 { // E-Wallet ID
		value = preprocessValue("03", value)
	} else { // Bank Account
		value = preprocessValue("04", value)
	}

	return value
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func QRConvertToBase64String(qrText string) string {
	// base64String := base64.StdEncoding.EncodeToString([]byte(qrText))

	// qrCode, err := qrcode.New(qrText, qrcode.Highest)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// qrCode.VersionNumber = 1
	// // qrCode.BackgroundColor = qrCode.BackgroundColor.RGBA()

	// img, err := qrCode.PNG(4)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var buf bytes.Buffer
	// err = png.Encode(&buf, img)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var png []byte
	png, err := qrcode.Encode(qrText, qrcode.Highest, 256)
	if err != nil {
		log.Fatal(err)
	}

	imgStr := base64.StdEncoding.EncodeToString(png)
	imgBase64Str := "data:image/png;base64," + imgStr

	return imgBase64Str
}
