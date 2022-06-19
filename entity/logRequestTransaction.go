package entity

const (
	TransactionStatusPending = "pending"

	// user success paid the transaction, after click the upload data in web, update to paid
	TransactionStatusPaid = "paid"

	// success transfer manual from admin, change to success
	// set Done to true
	TransactionStatusSuccess = "success"

	// failed for case user not paid, admin cannot transfer
	// set Done to true
	TransactionStatusFailed = "failed"

	// for more than 2 hours after request, change to expired if status in pending
	// set Done to true
	TransactionStatusExpired = "expired"
)

type LogRequestTransaction struct {
	Id             string `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	PersonalDataId string `json:"personal_data_id"`
	SenderNumber   string `json:"sender_number"`
	SenderWallet   string `json:"sender_wallet"`
	ReceiverName   string `json:"receiver_name"`
	ReceiverNumber string `json:"receiver_number"`
	ReceiverWallet string `json:"receiver_wallet"`
	AmountTransfer int    `json:"amount_transfer"`
	AdminFee       int    `json:"admin_fee"`
	AmountReceived int    `json:"amount_received"`
	Status         string `json:"status"` // [pending, paid, success, failed, expired]
	Done           bool   `json:"done"`   // default is false
}
