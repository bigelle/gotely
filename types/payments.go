package types

import (
	"encoding/json"
	"fmt"

	"github.com/bigelle/tele.go/interfaces"
	"github.com/bigelle/tele.go/internal/assertions"
)

type PaidMediaInfo struct {
	StarCount string      `json:"star_count"`
	PaidMedia []PaidMedia `json:"paid_media"`
}

type PaidMedia struct {
	PaidMediaInterface
}

// TODO: connect it with types
type PaidMediaInterface interface {
	paidMediaContract()
	interfaces.Validator
}

func (p *PaidMedia) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type       string `json:"type"`
		Attributes json.RawMessage
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw.Type {
	case "preview":
		p.PaidMediaInterface = new(PaidMediaPreview)
	case "photo":
		p.PaidMediaInterface = new(PaidMediaPhoto)
	case "video":
		p.PaidMediaInterface = new(PaidMediaVideo)
	default:
		return fmt.Errorf("Unrecognized type: %T", p.PaidMediaInterface)
	}
	return json.Unmarshal(raw.Attributes, &p.PaidMediaInterface)
}

type PaidMediaPhoto struct {
	Type  string      `json:"type"`
	Photo []PhotoSize `json:"photo"`
}

func (p PaidMediaPhoto) paidMediaContract() {}

func (p PaidMediaPhoto) Validate() error {
	if err := assertions.ParamNotEmpty(p.Type, "Type"); err != nil {
		return err
	}
	if assertions.IsSliceEmpty(p.Photo) {
		return fmt.Errorf("Photo parameter can't be empty")
	}
	return nil
}

type PaidMediaPreview struct {
	Type     string `json:"type"`
	Width    *int   `json:"width,omitempty"`
	Height   *int   `json:"height,omitempty"`
	Duration *int   `json:"duration,omitempty"`
}

func (p PaidMediaPreview) paidMediaContract() {}

func (p PaidMediaPreview) Validate() error {
	if err := assertions.ParamNotEmpty(p.Type, "Type"); err != nil {
		return err
	}
	return nil
}

type PaidMediaVideo struct {
	Type  string `json:"type"`
	Video *Video `json:"video"`
}

func (p PaidMediaVideo) paidMediaContract() {}

func (p PaidMediaVideo) Validate() error {
	if err := assertions.ParamNotEmpty(p.Type, "Type"); err != nil {
		return err
	}
	if p.Video == nil {
		return fmt.Errorf("Photo parameter can't be empty")
	}
	return nil
}

type StarTransaction struct {
	Id       string              `json:"id"`
	Amount   int                 `json:"amount"`
	Date     int                 `json:"date"`
	Source   *TransactionPartner `json:"source,omitempty"`
	Receiver *TransactionPartner `json:"receiver,omitempty"`
}

type StarTransactions struct {
	Transactions []StarTransaction `json:"transactions"`
}

type TransactionPartner struct {
	TransactionPartnerInterface
}

type TransactionPartnerInterface interface {
	transactionPartnerContract()
}

func (t *TransactionPartner) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type       string `json:"type"`
		Attributes json.RawMessage
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw.Type {
	case "fragment":
		t.TransactionPartnerInterface = new(TransactionPartnerFragment)
	case "user":
		t.TransactionPartnerInterface = new(TransactionPartnerUser)
	case "telegram_ads":
		t.TransactionPartnerInterface = new(TransactionPartnerTelegramAds)
	case "other":
		t.TransactionPartnerInterface = new(TransactionPartnerOther)
	default:
		return fmt.Errorf("Unrecognized type: %T", t.TransactionPartnerInterface)
	}
	return json.Unmarshal(raw.Attributes, &t.TransactionPartnerInterface)
}

type TransactionPartnerFragment struct {
	Type            string                 `json:"type"`
	WithdrawalState RevenueWithdrawalState `json:"withdrawal_state,omitempty"`
}

func (t TransactionPartnerFragment) transactionPartnerContract() {}

type TransactionPartnerOther struct {
	Type string `json:"type"`
}

func (t TransactionPartnerOther) transactionPartnerContract() {}

type TransactionPartnerTelegramAds struct {
	Type string `json:"type"`
}

func (t TransactionPartnerTelegramAds) transactionPartnerContract() {}

type TransactionPartnerUser struct {
	Type             string       `json:"type"`
	User             User         `json:"user"`
	InvoicePayload   *string      `json:"invoice_payload,omitempty"`
	PaidMedia        *[]PaidMedia `json:"paid_media,omitempty"`
	PaidMediaPayload *string      `json:"paid_media_payload,omitempty"`
}

func (t TransactionPartnerUser) transactionPartnerContract() {}

type RevenueWithdrawalState struct {
	RevenueWithdrawalStateInterface
}

type RevenueWithdrawalStateInterface interface {
	revenueWithdrawalStateContract()
}

func (r *RevenueWithdrawalState) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type       string `json:"type"`
		Attributes json.RawMessage
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw.Type {
	case "pending":
		r.RevenueWithdrawalStateInterface = new(RevenueWithdrawalStatePending)
	case "succeeded":
		r.RevenueWithdrawalStateInterface = new(RevenueWithdrawalStateSucceeded)
	case "failed":
		r.RevenueWithdrawalStateInterface = new(RevenueWithdrawalStateFailed)
	default:
		return fmt.Errorf("Unrecognized type: %T", r.RevenueWithdrawalStateInterface)
	}
	return json.Unmarshal(raw.Attributes, &r.RevenueWithdrawalStateInterface)
}

type RevenueWithdrawalStateFailed struct {
	Type string `json:"type"`
}

func (r RevenueWithdrawalStateFailed) revenueWithdrawalStateContract() {}

type RevenueWithdrawalStatePending struct {
	Type string `json:"type"`
}

func (r RevenueWithdrawalStatePending) revenueWithdrawalStateContract() {}

type RevenueWithdrawalStateSucceeded struct {
	Type string `json:"type"`
	Date int    `json:"date"`
	Url  string `json:"url"`
}

func (r RevenueWithdrawalStateSucceeded) revenueWithdrawalStateContract() {}

type Invoice struct {
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	StartParameter string    `json:"start_parameter"`
	Currency       string    `json:"currency"`
	TotalAmount    string    `json:"total_amount"`
	Photo          PhotoSize `json:"photo"`
}

type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

func (l LabeledPrice) Validate() error {
	if assertions.IsStringEmpty(l.Label) {
		return fmt.Errorf("label parameter can't be empty")
	}
	if l.Amount < 0 {
		return fmt.Errorf("Amount can't be less than zero")
	}
	return nil
}

type OrderInfo struct {
	Name            *string          `json:"name,omitempty"`
	PhoneNumber     *string          `json:"phone_number,omitempty"`
	Email           *string          `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type PaidMediaPurchased struct {
	User             User   `json:"user"`
	PaidMediaPayload string `json:"paid_media_payload"`
}

type PreCheckoutQuery struct {
	Id               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionId *string    `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

type RefundedPayment struct {
	Currency                string  `json:"currency"`
	TotalAmount             int     `json:"total_amount"`
	InvoicePayload          string  `json:"invoice_payload"`
	TelegramPaymentChargeId string  `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId *string `json:"provider_payment_charge_id,omitempty"`
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line_1"`
	StreetLine2 string `json:"street_line_2"`
	PostCode    string `json:"post_code"`
}

type ShippingQuery struct {
	Id              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             string     `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	TelegramPaymentChargeId string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId string     `json:"provider_payment_charge_id"`
	ShippingOptionId        *string    `json:"shipping_option_id,omitempty"`
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`
}
