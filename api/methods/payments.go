package methods

import (
	"strings"

	"github.com/bigelle/gotely/api/objects"
)

type SendInvoice struct {
	ChatId                    string                        `json:"chat_id"`
	Title                     string                        `json:"title"`
	Description               string                        `json:"description"`
	Payload                   string                        `json:"payload"`
	Currency                  string                        `json:"currency"`
	Prices                    []objects.LabeledPrice        `json:"prices"`
	MessageThreadId           *int                          `json:"message_thread_id,omitempty"`
	ProviderToken             *string                       `json:"provider_token,omitempty"`
	MaxTipAmount              *int                          `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       *[]int                        `json:"suggested_tip_amounts,omitempty"`
	StartParameter            *string                       `json:"start_parameter,omitempty"`
	ProviderData              *string                       `json:"provider_data,omitempty"`
	PhotoUrl                  *string                       `json:"photo_url,omitempty"`
	PhotoSize                 *int                          `json:"photo_size,omitempty"`
	PhotoWidth                *int                          `json:"photo_width,omitempty"`
	PhotoHeight               *int                          `json:"photo_height,omitempty"`
	NeedName                  *bool                         `json:"need_name,omitempty"`
	NeedPhoneNumber           *bool                         `json:"need_phone_number,omitempty"`
	NeedEmail                 *bool                         `json:"need_email,omitempty"`
	NeedShippingAddress       *bool                         `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider *bool                         `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       *bool                         `json:"send_email_to_provider,omitempty"`
	IsFlexible                *bool                         `json:"is_flexible,omitempty"`
	DisableNotification       *bool                         `json:"disable_notification,omitempty"`
	ProtectContent            *bool                         `json:"protect_content,omitempty"`
	AllowPaidBroadcast        *bool                         `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId           *string                       `json:"message_effect_id,omitempty"`
	ReplyParameters           *objects.ReplyParameters      `json:"reply_parameters,omitempty"`
	ReplyMarkup               *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
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

type CreateInvoiceLink struct {
	Title                     string                 `json:"title"`
	Description               string                 `json:"description"`
	Payload                   string                 `json:"payload"`
	Currency                  string                 `json:"currency"`
	Prices                    []objects.LabeledPrice `json:"prices"`
	BusinessConnectionId      *string                `json:"business_connection_id,omitempty"`
	ProviderToken             *string                `json:"provider_token,omitempty"`
	SubscriptionPeriod        *int                   `json:"subscription_period,omitempty"`
	MaxTipAmount              *int                   `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       *[]int                 `json:"suggested_tip_amounts,omitempty"`
	ProviderData              *string                `json:"provider_data,omitempty"`
	PhotoUrl                  *string                `json:"photo_url,omitempty"`
	PhotoSize                 *int                   `json:"photo_size,omitempty"`
	PhotoWidth                *int                   `json:"photo_width,omitempty"`
	PhotoHeight               *int                   `json:"photo_height,omitempty"`
	NeedName                  *bool                  `json:"need_name,omitempty"`
	NeedPhoneNumber           *bool                  `json:"need_phone_number,omitempty"`
	NeedEmail                 *bool                  `json:"need_email,omitempty"`
	NeedShippingAddress       *bool                  `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider *bool                  `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       *bool                  `json:"send_email_to_provider,omitempty"`
	IsFlexible                *bool                  `json:"is_flexible,omitempty"`
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

type AnswerShippingQuery struct {
	ShippingQueryId string                    `json:"shipping_query_id"`
	Ok              bool                      `json:"ok"`
	ShippingOptions *[]objects.ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    *string                   `json:"error_message,omitempty"`
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

type AnswerPreCheckoutQuery struct {
	PreCheckoutQueryId string `json:"pre_checkout_query_id"`
	Ok                 bool   `json:"ok"`
	ErrorMessage       *bool  `json:"error_message,omitempty"`
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

type GetStarTransactions struct {
	Offset *int `json:"offset,omitempty"`
	Limit  *int `json:"limit,omitempty"`
}

func (g GetStarTransactions) Validate() error {
	if g.Limit != nil {
		if *g.Limit < 1 || *g.Limit > 100 {
			return objects.ErrInvalidParam("limit parameter must be between 1 and 100")
		}
	}
	return nil
}

type RefundStarPayment struct {
	UserId                  int    `json:"user_id"`
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
} // FIXME: all of the channel and supergroup ids SHOULD BE NEGATIVE,
// should fix validation EVERYWHERE

func (r RefundStarPayment) Validate() error {
	if r.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if strings.TrimSpace(r.TelegramPaymentChargeId) == "" {
		return objects.ErrInvalidParam("telegram_payment_charge_id parameter can't be empty")
	}
	return nil
}

type EditUserStarSubscription struct {
	UserId                  int    `json:"user_id"`
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	IsCanceled              bool   `json:"is_canceled"`
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
