package objects

import (
	"encoding/json"

	"github.com/bigelle/tele.go/api/meta/interfaces"
)

// TODO: connect it with types
type PaidMedia interface {
	interfaces.Validable
	interfaces.BotApiObject
}

type PaidMediaInfo struct {
	StarCount string
	PaidMedia []PaidMedia
}

type PaidMediaPhoto struct {
	Type  string
	Photo []PhotoSize
}

func (p PaidMediaPhoto) MarshalJSON() ([]byte, error) {
	type Alias PaidMediaPhoto
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "photo",
		Alias: (*Alias)(&p),
	})
}

type PaidMediaPreview struct {
	Type     string
	Width    int
	Height   int
	Duration int
}

func (p PaidMediaPreview) MarshalJSON() ([]byte, error) {
	type Alias PaidMediaPreview
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "preview",
		Alias: (*Alias)(&p),
	})
}

type PaidMediaVideo struct {
	Type  string
	Video Video
}

func (p PaidMediaVideo) MarshalJSON() ([]byte, error) {
	type Alias PaidMediaVideo
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "video",
		Alias: (*Alias)(&p),
	})
}

type StarTransaction struct {
	Id       string
	Amount   int
	Date     int
	Source   TransactionPartner
	Receiver TransactionPartner
}

type StarTransactions struct {
	Transactions []StarTransaction
}

type TransactionPartner interface {
	interfaces.Validable
	interfaces.BotApiObject
}

type TransactionPartnerFragment struct {
	Type            string
	WithdrawalState RevenueWithdrawalState
}

func (t TransactionPartnerFragment) MarshalJSON() ([]byte, error) {
	type Alias TransactionPartnerFragment
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "fragment",
		Alias: (*Alias)(&t),
	})
}

type TransactionPartnerOther struct {
	Type string
}

func (t TransactionPartnerOther) MarshalJSON() ([]byte, error) {
	type Alias TransactionPartnerOther
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "other",
		Alias: (*Alias)(&t),
	})
}

type TransactionPartnerTelegramAds struct {
	Type string
}

func (t TransactionPartnerTelegramAds) MarshalJSON() ([]byte, error) {
	type Alias TransactionPartnerTelegramAds
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "telegram_ads",
		Alias: (*Alias)(&t),
	})
}

type TransactionPartnerUser struct {
	Type           string
	User           User
	InvoicePayload string
	PaidMedia      []PaidMedia
}

func (t TransactionPartnerUser) MarshalJSON() ([]byte, error) {
	type Alias TransactionPartnerUser
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "user",
		Alias: (*Alias)(&t),
	})
}

type RevenueWithdrawalState interface {
	interfaces.Validable
	interfaces.BotApiObject
}

// TODO: type

type RevenueWithdrawalStateFailed struct {
	Type string
}

func (t RevenueWithdrawalStateFailed) MarshalJSON() ([]byte, error) {
	type Alias RevenueWithdrawalStateFailed
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "failed",
		Alias: (*Alias)(&t),
	})
}

type RevenueWithdrawalStatePending struct {
	Type string
}

func (t RevenueWithdrawalStatePending) MarshalJSON() ([]byte, error) {
	type Alias RevenueWithdrawalStatePending
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "pending",
		Alias: (*Alias)(&t),
	})
}

type RevenueWithdrawalStateSucceeded struct {
	Type string
	Date int
	Url  string
}

func (t RevenueWithdrawalStateSucceeded) MarshalJSON() ([]byte, error) {
	type Alias RevenueWithdrawalStateSucceeded
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "succeeded",
		Alias: (*Alias)(&t),
	})
}

type RefundedPayment struct {
	Currency                string
	TotalAmount             int
	InvoicePayload          string
	TelegramPaymentChargeId string
	ProviderPaymentChargeId string
}

type OrderInfo struct {
	Name            string
	PhoneNumber     string
	Email           string
	ShippingAddress ShippingAddress
}

type ShippingAddress struct {
	CountryCode string
	State       string
	City        string
	StreetLine1 string
	StreetLine2 string
	PostCode    string
}

type SuccessfulPayment struct {
	Currency                string
	TotalAmount             string
	InvoicePayload          string
	ShippingOptionId        string
	OrderInfo               OrderInfo
	TelegramPaymentChargeId string
	ProviderPaymentChargeId string
}
