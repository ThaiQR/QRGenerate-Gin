package model

type CustomerPromptpayModel struct {
	PFI          string
	PIM          string
	AID          string
	ReceiveID    string
	MerchantCode string
	Currency     string
	Amount       string
	CountryCode  string
	MerchantCity string
	PostalCode   string
	CRC          string
}

func (m *CustomerPromptpayModel) SetValues(receiveID, amount string, onetime bool) {
	if onetime {
		m.PIM = "12"
	} else {
		m.PIM = "11"
	}

	m.PFI = "01"
	// m.PIM = pim
	m.AID = "A000000677010114" // Promptpay for Merchant
	m.ReceiveID = receiveID    // mobile:13, nid/Taxid: 13, EwalletId: 15, BankAcc: up to 43 (BOTcode 3 digits + account no.)
	m.MerchantCode = ""        // Can be null
	m.Currency = "764"         // THB
	m.Amount = amount
	m.CountryCode = "TH" // Thai
	m.PostalCode = ""
	m.CRC = ""
}
