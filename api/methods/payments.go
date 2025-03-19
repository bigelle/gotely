// TODO: make optional and required fields more obvious
package methods

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/bigelle/gotely/api/objects"
)

// Use this method to send invoices.
// On success, the sent [objects.Message] is returned.
type SendInvoice struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	// Product name, 1-32 characters
	Title string `json:"title"`
	// Product description, 1-255 characters
	Description string `json:"description"`
	// Bot-defined invoice payload, 1-128 bytes.
	// This will not be displayed to the user, use it for your internal processes.
	Payload string `json:"payload"`
	// Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	ProviderToken *string `json:"provider_token,omitempty"`
	// Three-letter ISO 4217 currency code, see more on currencies. Pass “XTR” for payments in Telegram Stars.
	Currency string `json:"currency"`
	// Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.).
	// Must contain exactly one item for payments in Telegram Stars.
	Prices []objects.LabeledPrice `json:"prices"`
	// The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double).
	// For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json,
	// it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	// Defaults to 0. Not supported for payments in Telegram Stars.
	MaxTipAmount *int `json:"max_tip_amount,omitempty"`
	// A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double).
	// At most 4 suggested tip amounts can be specified.
	// The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	SuggestedTipAmounts *[]int `json:"suggested_tip_amounts,omitempty"`
	// Unique deep-linking parameter. If left empty, forwarded copies of the sent message will have a Pay button,
	// allowing multiple users to pay directly from the forwarded message, using the same invoice.
	// If non-empty, forwarded copies of the sent message will have a URL button with a deep link to the bot (instead of a Pay button),
	// with the value used as the start parameter
	StartParameter *string `json:"start_parameter,omitempty"`
	// JSON-serialized data about the invoice, which will be shared with the payment provider.
	// A detailed description of required fields should be provided by the payment provider.
	ProviderData *string `json:"provider_data,omitempty"`
	// URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	// People like it better when they see what they are paying for.
	PhotoUrl *string `json:"photo_url,omitempty"`
	// Photo size in bytes
	PhotoSize *int `json:"photo_size,omitempty"`
	// Photo width
	PhotoWidth *int `json:"photo_width,omitempty"`
	// Photo height
	PhotoHeight *int `json:"photo_height,omitempty"`
	// Pass True if you require the user's full name to complete the order.
	// Ignored for payments in Telegram Stars.
	NeedName *bool `json:"need_name,omitempty"`
	// Pass True if you require the user's phone number to complete the order.
	// Ignored for payments in Telegram Stars.
	NeedPhoneNumber *bool `json:"need_phone_number,omitempty"`
	// Pass True if you require the user's email address to complete the order.
	// Ignored for payments in Telegram Stars.
	NeedEmail *bool `json:"need_email,omitempty"`
	// Pass True if you require the user's shipping address to complete the order.
	// Ignored for payments in Telegram Stars.
	NeedShippingAddress *bool `json:"need_shipping_address,omitempty"`
	// Pass True if the user's phone number should be sent to the provider.
	// Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider *bool `json:"send_phone_number_to_provider,omitempty"`
	// Pass True if the user's email address should be sent to the provider.
	// Ignored for payments in Telegram Stars.
	SendEmailToProvider *bool `json:"send_email_to_provider,omitempty"`
	// Pass True if the final price depends on the shipping method.
	// Ignored for payments in Telegram Stars.
	IsFlexible *bool `json:"is_flexible,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown.
	// If not empty, the first button must be a Pay button.
	ReplyMarkup *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (s SendInvoice) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if len(s.Title) < 1 || len(s.Title) > 32 {
		return objects.ErrInvalidParam("title parameter must be between 1 and 32 characters long")
	}
	if len(s.Description) < 1 || len(s.Description) > 255 {
		return objects.ErrInvalidParam("description parameter must be between 1 and 255 characters long")
	}
	if len([]byte(s.Payload)) < 1 || len([]byte(s.Payload)) > 128 {
		return objects.ErrInvalidParam("payload parameter must be between 1 and 128 bytes long")
	}
	if len(s.Prices) < 1 {
		return objects.ErrInvalidParam("prices parameter can't be empty")
	}
	for _, price := range s.Prices {
		if err := price.Validate(); err != nil {
			return err
		}
	}
	if s.SuggestedTipAmounts != nil {
		if len(*s.SuggestedTipAmounts) > 4 {
			return objects.ErrInvalidParam("invalid suggested_tip_amounts parameter: at most 4 suggested tip amounts can be specified")
		}
		if (*s.SuggestedTipAmounts)[0] < 0 {
			return objects.ErrInvalidParam("invalid suggested_tip_amounts parameter: prices must be positive")
		}
		if s.MaxTipAmount != nil && (*s.SuggestedTipAmounts)[0] > *s.MaxTipAmount {
			return objects.ErrInvalidParam("invalid suggested_tip_amounts parameter: prices must not exceed max_tip_amount.")
		}
		for i := 1; i < len(*s.SuggestedTipAmounts); i++ {
			if (*s.SuggestedTipAmounts)[i] < 0 {
				return objects.ErrInvalidParam("invalid suggested_tip_amounts parameter: prices must be positive")
			}
			if (*s.SuggestedTipAmounts)[i-1] > (*s.SuggestedTipAmounts)[i] {
				return objects.ErrInvalidParam("invalid suggested_tip_amounts parameter: prices must be passed in a strictly increased order")
			}
			if s.MaxTipAmount != nil && (*s.SuggestedTipAmounts)[i] > *s.MaxTipAmount {
				return objects.ErrInvalidParam("invalid suggested_tip_amounts parameter: prices must not exceed max_tip_amount.")
			}
		}
	}
	return nil
}

func (s SendInvoice) Endpoint() string {
	return "sendInvoice"
}

func (s SendInvoice) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s SendInvoice) ContentType() string {
	return "application/json"
}

// Use this method to create a link for an invoice.
// Returns the created invoice link as String on success.
type CreateInvoiceLink struct {
	// Unique identifier of the business connection on behalf of which the link will be created.
	// For payments in Telegram Stars only.
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Product name, 1-32 characters
	Title string `json:"title"`
	// Product description, 1-255 characters
	Description string `json:"description"`
	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Payload string `json:"payload"`
	// Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	ProviderToken *string `json:"provider_token,omitempty"`
	// Three-letter ISO 4217 currency code, see more on currencies. Pass “XTR” for payments in Telegram Stars.
	Currency string `json:"currency"`
	// Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.).
	// Must contain exactly one item for payments in Telegram Stars.
	Prices []objects.LabeledPrice `json:"prices"`
	// The number of seconds the subscription will be active for before the next payment.
	// The currency must be set to “XTR” (Telegram Stars) if the parameter is used.
	// Currently, it must always be 2592000 (30 days) if specified. \
	// Any number of subscriptions can be active for a given bot at the same time, including multiple concurrent subscriptions from the same user.
	// Subscription price must no exceed 2500 Telegram Stars.
	SubscriptionPeriod *int `json:"subscription_period,omitempty"`
	// The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double).
	// For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json,
	// it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	// Defaults to 0. Not supported for payments in Telegram Stars.
	MaxTipAmount *int `json:"max_tip_amount,omitempty"`
	// A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double).
	// At most 4 suggested tip amounts can be specified.
	// The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	SuggestedTipAmounts *[]int `json:"suggested_tip_amounts,omitempty"`
	// JSON-serialized data about the invoice, which will be shared with the payment provider.
	// A detailed description of required fields should be provided by the payment provider.
	ProviderData *string `json:"provider_data,omitempty"`
	// URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	PhotoUrl *string `json:"photo_url,omitempty"`
	// Photo size in bytes
	PhotoSize *int `json:"photo_size,omitempty"`
	// Photo width
	PhotoWidth *int `json:"photo_width,omitempty"`
	// Photo height
	PhotoHeight *int `json:"photo_height,omitempty"`
	// Pass True if you require the user's full name to complete the order.
	// Ignored for payments in Telegram Stars.
	NeedName *bool `json:"need_name,omitempty"`
	// Pass True if you require the user's phone number to complete the order.
	// Ignored for payments in Telegram Stars.
	NeedPhoneNumber *bool `json:"need_phone_number,omitempty"`
	// Pass True if you require the user's email address to complete the order.
	// Ignored for payments in Telegram Stars.
	NeedEmail *bool `json:"need_email,omitempty"`
	// Pass True if you require the user's shipping address to complete the order.
	// Ignored for payments in Telegram Stars.
	NeedShippingAddress *bool `json:"need_shipping_address,omitempty"`
	// Pass True if the user's phone number should be sent to the provider.
	// Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider *bool `json:"send_phone_number_to_provider,omitempty"`
	// Pass True if the user's email address should be sent to the provider.
	// Ignored for payments in Telegram Stars.
	SendEmailToProvider *bool `json:"send_email_to_provider,omitempty"`
	// Pass True if the final price depends on the shipping method.
	// Ignored for payments in Telegram Stars.
	IsFlexible *bool `json:"is_flexible,omitempty"`
}

func (c CreateInvoiceLink) Validate() error {
	if len(c.Title) < 1 || len(c.Title) > 32 {
		return objects.ErrInvalidParam("title parameter must be between 1 and 32 characters long")
	}
	if len(c.Description) < 1 || len(c.Description) > 255 {
		return objects.ErrInvalidParam("description parameter must be between 1 and 255 characters long")
	}
	if len([]byte(c.Payload)) < 1 || len([]byte(c.Payload)) > 128 {
		return objects.ErrInvalidParam("payload parameter must be between 1 and 128 bytes long")
	}
	if len(c.Prices) < 1 {
		return objects.ErrInvalidParam("prices parameter can't be empty")
	}
	for _, price := range c.Prices {
		if err := price.Validate(); err != nil {
			return err
		}
	}
	if c.SuggestedTipAmounts != nil {
		if len(*c.SuggestedTipAmounts) > 4 {
			return objects.ErrInvalidParam("invalid suggested_tip_amounts parameter: at most 4 suggested tip amounts can be specified")
		}
		if (*c.SuggestedTipAmounts)[0] < 0 {
			return objects.ErrInvalidParam("invalid suggested_tip_amounts parameter: prices must be positive")
		}
		if c.MaxTipAmount != nil && (*c.SuggestedTipAmounts)[0] > *c.MaxTipAmount {
			return objects.ErrInvalidParam("invalid suggested_tip_amounts parameter: prices must not exceed max_tip_amount.")
		}
		for i := 1; i < len(*c.SuggestedTipAmounts); i++ {
			if (*c.SuggestedTipAmounts)[i] < 0 {
				return objects.ErrInvalidParam("invalid suggested_tip_amounts parameter: prices must be positive")
			}
			if (*c.SuggestedTipAmounts)[i-1] > (*c.SuggestedTipAmounts)[i] {
				return objects.ErrInvalidParam("invalid suggested_tip_amounts parameter: prices must be passed in a strictly increased order")
			}
			if c.MaxTipAmount != nil && (*c.SuggestedTipAmounts)[i] > *c.MaxTipAmount {
				return objects.ErrInvalidParam("invalid suggested_tip_amounts parameter: prices must not exceed max_tip_amount.")
			}
		}
	}
	return nil
}

func (s CreateInvoiceLink) Endpoint() string {
	return "createInvoiceLink"
}

func (s CreateInvoiceLink) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s CreateInvoiceLink) ContentType() string {
	return "application/json"
}

// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified,
// the Bot API will send an [objects.Update] with a shipping_query field to the bot. Use this method to reply to shipping queries.
// On success, True is returned.
type AnswerShippingQuery struct {
	// Unique identifier for the query to be answered
	ShippingQueryId string `json:"shipping_query_id"`
	//Pass True if delivery to the specified address is possible and False if there are any problems
	//(for example, if delivery to the specified address is not possible)
	Ok bool `json:"ok"`
	// Required if ok is True. A JSON-serialized array of available shipping options.
	ShippingOptions *[]objects.ShippingOption `json:"shipping_options,omitempty"`
	//Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order
	//(e.g. “Sorry, delivery to your desired address is unavailable”). Telegram will display this message to the user.
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (a AnswerShippingQuery) Validate() error {
	if strings.TrimSpace(a.ShippingQueryId) == "" {
		return objects.ErrInvalidParam("shipping_query_id parameter can't be empty")
	}
	if a.Ok && a.ShippingOptions == nil {
		return objects.ErrInvalidParam("shipping_options parameter can't be empty if ok == true")
	}
	if !a.Ok && a.ErrorMessage == nil {
		return objects.ErrInvalidParam("error_message parameter can't be empty if ok == false")
	}
	if a.ShippingOptions != nil {
		for _, opt := range *a.ShippingOptions {
			if err := opt.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s AnswerShippingQuery) Endpoint() string {
	return "answerShippingQuery"
}

func (s AnswerShippingQuery) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s AnswerShippingQuery) ContentType() string {
	return "application/json"
}

// Once the user has confirmed their payment and shipping details,
// the Bot API sends the final confirmation in the form of an [objects.Update] with the field pre_checkout_query.
// Use this method to respond to such pre-checkout queries. On success, True is returned.
// Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
type AnswerPreCheckoutQuery struct {
	// Unique identifier for the query to be answered
	PreCheckoutQueryId string `json:"pre_checkout_query_id"`
	// Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order.
	// Use False if there are any problems.
	Ok bool `json:"ok"`
	//Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout
	//(e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details.
	//Please choose a different color or garment!").
	//Telegram will display this message to the user.
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (a AnswerPreCheckoutQuery) Validate() error {
	if strings.TrimSpace(a.PreCheckoutQueryId) == "" {
		return objects.ErrInvalidParam("pre_checkout_query_id parameter can't be empty")
	}
	if !a.Ok && a.ErrorMessage == nil {
		return objects.ErrInvalidParam("error_message parameter can't be empty if ok == false")
	}
	return nil
}

func (s AnswerPreCheckoutQuery) Endpoint() string {
	return "answerPreCheckoutQuery"
}

func (s AnswerPreCheckoutQuery) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s AnswerPreCheckoutQuery) ContentType() string {
	return "application/json"
}

// Returns the bot's Telegram Star transactions in chronological order.
// On success, returns a [objects.StarTransactions] object.
type GetStarTransactions struct {
	// Number of transactions to skip in the response
	Offset *int `json:"offset,omitempty"`
	// The maximum number of transactions to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit *int `json:"limit,omitempty"`
}

func (g GetStarTransactions) Validate() error {
	if g.Limit != nil {
		if *g.Limit < 1 || *g.Limit > 100 {
			return objects.ErrInvalidParam("limit parameter must be between 1 and 100")
		}
	}
	return nil
}

func (s GetStarTransactions) Endpoint() string {
	return "getStarTransactions"
}

func (s GetStarTransactions) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s GetStarTransactions) ContentType() string {
	return "application/json"
}

// Refunds a successful payment in Telegram Stars. Returns True on success.
type RefundStarPayment struct {
	// Identifier of the user whose payment will be refunded
	UserId int `json:"user_id"`
	// Telegram payment identifier
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
} // FIXME: all of the channel and supergroup ids SHOULD BE NEGATIVE,
// should fix validation EVERYWHERE

func (r RefundStarPayment) Validate() error {
	if r.UserId <= 0 {
		return objects.ErrInvalidParam("user_id parameter can't be empty or negative")
	}
	if strings.TrimSpace(r.TelegramPaymentChargeId) == "" {
		return objects.ErrInvalidParam("telegram_payment_charge_id parameter can't be empty")
	}
	return nil
}

func (s RefundStarPayment) Endpoint() string {
	return "refundStarPayment"
}

func (s RefundStarPayment) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s RefundStarPayment) ContentType() string {
	return "application/json"
}

// Allows the bot to cancel or re-enable extension of a subscription paid in Telegram Stars. Returns True on success.
type EditUserStarSubscription struct {
	// Identifier of the user whose subscription will be edited
	UserId int `json:"user_id"`
	// Telegram payment identifier for the subscription
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	// Pass True to cancel extension of the user subscription;
	// the subscription must be active up to the end of the current subscription period.
	// Pass False to allow the user to re-enable a subscription that was previously canceled by the bot.
	IsCanceled bool `json:"is_canceled"`
}

func (e EditUserStarSubscription) Validate() error {
	if e.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if strings.TrimSpace(e.TelegramPaymentChargeId) == "" {
		return objects.ErrInvalidParam("telegram_payment_charge_id parameter can't be empty")
	}
	return nil
}

func (s EditUserStarSubscription) Endpoint() string {
	return "editUserStarSubscription"
}

func (s EditUserStarSubscription) Reader() (io.Reader, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (s EditUserStarSubscription) ContentType() string {
	return "application/json"
}
