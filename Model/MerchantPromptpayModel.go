package model

type MerchantPromptpayModel struct {
	PFI          string
	PIM          string
	AID          string
	RecieveID    string
	MerchantCode string
	Currency     string
	Amount       string
	CountryCode  string
	MerchantCity string
	PostalCode   string
	CRC          string
}

func (m *MerchantPromptpayModel) SetValues(recieveID, amount string, onetime bool) {
	if onetime {
		m.PIM = "12"
	} else {
		m.PIM = "11"
	}

	m.PFI = "01"
	// m.PIM = pim
	m.AID = "A000000677010111" // Promptpay for Merchant
	m.RecieveID = recieveID    // mobile:13, nid/Taxid: 13, EwalletId: 15, BankAcc: up to 43 (BOTcode 3 digits + account no.)
	m.MerchantCode = ""        // Can be null
	m.Currency = "764"         // THB
	m.Amount = amount
	m.CountryCode = "TH" // Thai
	m.PostalCode = ""
	m.CRC = ""
}
