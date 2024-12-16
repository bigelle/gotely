package types

import (
	"encoding/json"
	"errors"
	"strings"
)

// This object represents a portion of the price for goods or services.
type LabeledPrice struct {
	//Portion label
	Label string `json:"label"`
	//Price of the product in the smallest units of the currency (integer, not float/double).
	//For example, for a price of US$ 1.45 pass amount = 145.
	//See the exp parameter in https://core.telegram.org/bots/payments/currencies.json,
	//it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	Amount int `json:"amount"`
}

func (l LabeledPrice) Validate() error {
	if strings.TrimSpace(l.Label) == "" {
		return ErrInvalidParam("label parameter can't be empty")
	}
	if l.Amount < 0 {
		return ErrInvalidParam("amount can't be less than zero")
	}
	return nil
}

// This object contains basic information about an invoice.
type Invoice struct {
	//Product name
	Title string `json:"title"`
	//Product description
	Description string `json:"description"`
	//Unique bot deep-linking parameter that can be used to generate this invoice
	StartParameter string `json:"start_parameter"`
	//Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars
	Currency string `json:"currency"`
	//Total price in the smallest units of the currency (integer, not float/double).
	//For example, for a price of US$ 1.45 pass amount = 145.
	//See the exp parameter in https://core.telegram.org/bots/payments/currencies.json,
	//it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount string `json:"total_amount"`
}

// This object represents a shipping address.
type ShippingAddress struct {
	//Two-letter ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`
	//State, if applicable
	State string `json:"state"`
	//City
	City string `json:"city"`
	//First line for the address
	StreetLine1 string `json:"street_line1"`
	//Second line for the address
	StreetLine2 string `json:"street_line2"`
	//Address post code
	PostCode string `json:"post_code"`
}

// This object represents information about an order.
type OrderInfo struct {
	//Optional. User name
	Name *string `json:"name,omitempty"`
	//Optional. User's phone number
	PhoneNumber *string `json:"phone_number,omitempty"`
	//Optional. User email
	Email *string `json:"email,omitempty"`
	//Optional. User shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// This object represents one shipping option.
type ShippingOption struct {
	//Shipping option identifier
	Id string `json:"id"`
	//Option title
	Title string `json:"title"`
	//List of price portions
	Prices []LabeledPrice `json:"prices"`
}

func (s ShippingOption) Validate() error {
	if strings.TrimSpace(s.Id) == "" {
		return ErrInvalidParam("id parameter can't be empty")
	}
	if strings.TrimSpace(s.Title) == "" {
		return ErrInvalidParam("title parameter can't be empty")
	}
	for _, price := range s.Prices {
		if err := price.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This object contains basic information about a successful payment.
type SuccessfulPayment struct {
	//Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars
	Currency string `json:"currency"`
	//Total price in the smallest units of the currency (integer, not float/double).
	//For example, for a price of US$ 1.45 pass amount = 145.
	//See the exp parameter in https://core.telegram.org/bots/payments/currencies.json,
	//it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount string `json:"total_amount"`
	//Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	//Optional. Expiration date of the subscription, in Unix time; for recurring payments only
	SubscriptionExpirationDate *int `json:"subscription_expiration_date,omitempty"`
	//Optional. True, if the payment is a recurring payment for a subscription
	IsRecurring *bool `json:"is_recurring,omitempty"`
	//Optional. True, if the payment is the first payment for a subscription
	IsFirstRecurring *bool `json:"is_first_recurring,omitempty"`
	//Optional. Identifier of the shipping option chosen by the user
	ShippingOptionId *string `json:"shipping_option_id,omitempty"`
	//Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
	//Telegram payment identifier
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	//Provider payment identifier
	ProviderPaymentChargeId string `json:"provider_payment_charge_id"`
}

// This object contains basic information about a refunded payment.
type RefundedPayment struct {
	//Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars.
	//Currently, always “XTR”
	Currency string `json:"currency"`
	//Total refunded price in the smallest units of the currency (integer, not float/double).
	//For example, for a price of US$ 1.45, total_amount = 145.
	//See the exp parameter in https://core.telegram.org/bots/payments/currencies.json,
	//it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
	//Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	//Telegram payment identifier
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	//Optional. Provider payment identifier
	ProviderPaymentChargeId *string `json:"provider_payment_charge_id,omitempty"`
}

// This object contains information about an incoming shipping query.
type ShippingQuery struct {
	//Unique query identifier
	Id string `json:"id"`
	//User who sent the query
	From User `json:"from"`
	//Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	//User specified shipping address
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

// This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	//Unique query identifier
	Id string `json:"id"`
	//User who sent the query
	From *User `json:"from"`
	//Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars
	Currency string `json:"currency"`
	//Total price in the smallest units of the currency (integer, not float/double).
	//For example, for a price of US$ 1.45 pass amount = 145.
	//See the exp parameter in https://core.telegram.org/bots/payments/currencies.json,
	//it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
	//Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	//Optional. Identifier of the shipping option chosen by the user
	ShippingOptionId *string `json:"shipping_option_id,omitempty"`
	//Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}

// This object contains information about a paid media purchase.
type PaidMediaPurchased struct {
	//User who purchased the media
	User User `json:"user"`
	//Bot-specified paid media payload
	PaidMediaPayload string `json:"paid_media_payload"`
}

// This object describes the state of a revenue withdrawal operation.
// Currently, it can be one of
//
// - RevenueWithdrawalStatePending
//
// - RevenueWithdrawalStateSucceeded
//
// - RevenueWithdrawalStateFailed
type RevenueWithdrawalState struct {
	RevenueWithdrawalStateInterface
}

type RevenueWithdrawalStateInterface interface {
	GetRevenueWithdrawalStateType() string
}

func (r *RevenueWithdrawalState) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "pending":
		tmp := RevenueWithdrawalStatePending{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		r.RevenueWithdrawalStateInterface = tmp
	case "succeeded":
		tmp := RevenueWithdrawalStateSucceeded{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		r.RevenueWithdrawalStateInterface = tmp
	case "failed":
		tmp := RevenueWithdrawalStateFailed{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		r.RevenueWithdrawalStateInterface = tmp
	default:
		return errors.New("type must be pending, succeeded or failed")
	}
	return nil
}

// The withdrawal is in progress.
type RevenueWithdrawalStatePending struct {
	//Type of the state, always “pending”
	Type string `json:"type"`
}

func (r RevenueWithdrawalStatePending) GetRevenueWithdrawalStateType() string {
	return "pending"
}

// The withdrawal succeeded.
type RevenueWithdrawalStateSucceeded struct {
	//Type of the state, always “succeeded”
	Type string `json:"type"`
	//Date the withdrawal was completed in Unix time
	Date int `json:"date"`
	//An HTTPS URL that can be used to see transaction details
	Url string `json:"url"`
}

func (r RevenueWithdrawalStateSucceeded) GetRevenueWithdrawalStateType() string {
	return "succeeded"
}

// The withdrawal failed and the transaction was refunded.
type RevenueWithdrawalStateFailed struct {
	//Type of the state, always “failed”
	Type string `json:"type"`
}

func (r RevenueWithdrawalStateFailed) GetRevenueWithdrawalStateType() string {
	return "failed"
}

// Contains information about the affiliate that received a commission via this transaction.
type AffiliateInfo struct {
	//Optional. The bot or the user that received an affiliate commission if it was received by a bot or a user
	AffiliateUser *User
	//Optional. The chat that received an affiliate commission if it was received by a chat
	AffiliateChat *Chat
	//The number of Telegram Stars received by the affiliate for each 1000 Telegram Stars received by the bot from referred users
	CommissionPerMile int
	//Integer amount of Telegram Stars received by the affiliate from the transaction, rounded to 0; can be negative for refunds
	Amount int
	//Optional. The number of 1/1000000000 shares of Telegram Stars received by the affiliate;
	//from -999999999 to 999999999; can be negative for refunds
	NanostarAmount *int
}

// This object describes the source of a transaction, or its recipient for outgoing transactions.
// Currently, it can be one of
// - TransactionPartnerUser
//
// - TransactionPartnerAffiliateProgram
//
// - TransactionPartnerFragment
//
// - TransactionPartnerTelegramAds
//
// - TransactionPartnerTelegramApi
//
// - TransactionPartnerOther
type TransactionPartner struct {
	//FIXME: FINALLY FOUND A BETTER SOLUTION
	//should make an object that contains EVERY POSSIBLE FIELD
	//and has a method that returns only that type THAT YOU NEED
	//OR an error if you tried to use wrong method
	TransactionPartnerInterface
}

type TransactionPartnerInterface interface {
	GetTransactionPartnerType() string
}

func (t *TransactionPartner) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case "user":
		tmp := TransactionPartnerUser{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		t.TransactionPartnerInterface = tmp
	case "affiliate_program":
		tmp := TransactionPartnerAffiliateProgram{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		t.TransactionPartnerInterface = tmp
	case "fragment":
		tmp := TransactionPartnerFragment{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		t.TransactionPartnerInterface = tmp
	case "telegram_ads":
		tmp := TransactionPartnerTelegramAds{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		t.TransactionPartnerInterface = tmp
	case "telegram_api":
		tmp := TransactionPartnerTelegramApi{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		t.TransactionPartnerInterface = tmp
	case "other":
		tmp := TransactionPartnerOther{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		t.TransactionPartnerInterface = tmp
	default:
		return errors.New("type must be fragment, user, telegram_ads or other")
	}
	return nil
}

// Describes a transaction with a user.
type TransactionPartnerUser struct {
	//Type of the transaction partner, always “user”
	Type string `json:"type"`
	//Information about the user
	User User `json:"user"`
	//Optional. Information about the affiliate that received a commission via this transaction
	Affiliate *AffiliateInfo `json:"affiliate,omitempty"`
	//Optional. Bot-specified invoice payload
	InvoicePayload *string `json:"invoice_payload,omitempty,"`
	//Optional. The duration of the paid subscription
	SubscriptionPeriod *int `json:"subscription_period"`
	//Optional. Information about the paid media bought by the user
	PaidMedia *[]PaidMedia `json:"paid_media,omitempty,"`
	//Optional. Bot-specified paid media payload
	PaidMediaPayload *string `json:"paid_media_payload,omitempty,"`
	//Optional. The gift sent to the user by the bot
	Gift *Gift `json:"gift,omitempty"`
}

func (t TransactionPartnerUser) GetTransactionPartnerType() string {
	return "user"
}

// Describes the affiliate program that issued the affiliate commission received via this transaction.
type TransactionPartnerAffiliateProgram struct {
	//Type of the transaction partner, always “affiliate_program”
	Type string `json:"type"`
	//Optional. Information about the bot that sponsored the affiliate program
	SponsorUser *User `json:"sponsor_user,omitempty"`
	//The number of Telegram Stars received by the bot for each 1000 Telegram Stars received by the affiliate program sponsor from referred users
	CommissionPerMile int `json:"commission_per_mile"`
}

func (TransactionPartnerAffiliateProgram) GetTransactionPartnerType() string {
	return "affiliate_program"
}

// Describes a withdrawal transaction with Fragment.
type TransactionPartnerFragment struct {
	//Type of the transaction partner, always “fragment”
	Type string `json:"type"`
	//Optional. State of the transaction if the transaction is outgoing
	WithdrawalState *RevenueWithdrawalState `json:"withdrawal_state,omitempty"`
}

func (t TransactionPartnerFragment) GetTransactionPartnerType() string {
	return "fragment"
}

// Describes a withdrawal transaction to the Telegram Ads platform.
type TransactionPartnerTelegramAds struct {
	//Type of the transaction partner, always “telegram_ads”
	Type string `json:"type"`
}

func (t TransactionPartnerTelegramAds) GetTransactionPartnerType() string {
	return "telegram_ads"
}

// Describes a transaction with payment for https://core.telegram.org/bots/api#paid-broadcasts.
type TransactionPartnerTelegramApi struct {
	//Type of the transaction partner, always “telegram_api”
	Type string `json:"type"`
	//The number of successful requests that exceeded regular limits and were therefore billed
	RequestCount int `json:"request_count"`
}

func (t TransactionPartnerTelegramApi) GetTransactionPartnerType() string {
	return "telegram_api"
}

// Describes a transaction with an unknown source or recipient.
type TransactionPartnerOther struct {
	//Type of the transaction partner, always “other”
	Type string `json:"type"`
}

func (t TransactionPartnerOther) GetTransactionPartnerType() string {
	return "other"
}

// Describes a Telegram Star transaction.
type StarTransaction struct {
	//Unique identifier of the transaction.
	//Coincides with the identifier of the original transaction for refund transactions.
	//Coincides with SuccessfulPayment.telegram_payment_charge_id for successful incoming payments from users.
	Id string `json:"id"`
	//Integer amount of Telegram Stars transferred by the transaction
	Amount int `json:"amount"`
	//Optional. The number of 1/1000000000 shares of Telegram Stars transferred by the transaction; from 0 to 999999999
	NanostarAmount *int `json:"nanostar_amount,omitempty"`
	//Date the transaction was created in Unix time
	Date int `json:"date"`
	//Optional. Source of an incoming transaction
	//(e.g., a user purchasing goods or services, Fragment refunding a failed withdrawal). Only for incoming transactions
	Source *TransactionPartner `json:"source,omitempty,"`
	//Optional. Receiver of an outgoing transaction
	//(e.g., a user for a purchase refund, Fragment for a withdrawal). Only for outgoing transactions
	Receiver *TransactionPartner `json:"receiver,omitempty,"`
}

// Contains a list of Telegram Star transactions.
type StarTransactions struct {
	//The list of transactions
	Transactions []StarTransaction `json:"transactions"`
}
