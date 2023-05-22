package data

type QRGenerateData struct {
	PFI          string
	PIM          string
	AID          string
	BillerID     string
	Reference1   string
	Reference2   string
	MerchantCode string
	Currency     string
	Amount       string
	CountryCode  string
	MerchantName string
	TerminalID   string
	CRC          string
}

func (q *QRGenerateData) SetValues(billerID, reference1, reference2, amount, merchantName string) {
	q.PFI = "01"
	q.PIM = "12"
	q.AID = "A000000677010112" // Merchant TQR
	q.BillerID = billerID
	q.Reference1 = reference1
	q.Reference2 = reference2
	q.MerchantCode = "" // Can be null
	q.Currency = "764"  // THB
	q.Amount = amount
	q.CountryCode = "TH" // Thai
	q.MerchantName = merchantName
	q.TerminalID = "" // Can be null
	q.CRC = ""
}
