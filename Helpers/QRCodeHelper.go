package helpers

import (
	data "BankGW-Gin/Data"
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/howeyc/crc16"
)

// func MerchantGenerate(pfi, pim, aid, billerID, ref01, ref02, merchantCode, currency, amount, countryCode, merchantName, terminalID string) string {
func QRMerchantGenerate(value data.QRGenerateData) string {
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
	merchantIdentifier = "30" + strconv.Itoa(len(merchantSum)) + merchantSum
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

	data += strconv.FormatInt(int64(crcResult), 16)

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

func QRConvertToBase64(qrText string) string {
	base64String := base64.StdEncoding.EncodeToString([]byte(qrText))

	return base64String
}
