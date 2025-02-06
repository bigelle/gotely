package methods

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bigelle/gotely/objects"
)

type SendInvoice struct {
	ChatId                    string
	Title                     string
	Description               string
	Payload                   string
	Currency                  string
	Prices                    []objects.LabeledPrice
	MessageThreadId           *int
	ProviderToken             *string
	MaxTipAmount              *int
	SuggestedTipAmounts       *[]int
	StartParameter            *string
	ProviderData              *string
	PhotoUrl                  *string
	PhotoSize                 *int
	PhotoWidth                *int
	PhotoHeight               *int
	NeedName                  *bool
	NeedPhoneNumber           *bool
	NeedEmail                 *bool
	NeedShippingAddress       *bool
	SendPhoneNumberToProvider *bool
	SendEmailToProvider       *bool
	IsFlexible                *bool
	DisableNotification       *bool
	ProtectContent            *bool
	AllowPaidBroadcast        *bool
	MessageEffectId           *string
	ReplyParameters           *objects.ReplyParameters
	ReplyMarkup               *objects.InlineKeyboardMarkup
	client                    *http.Client
	baseUrl                   string
}

func (s *SendInvoice) WithClient(c *http.Client) *SendInvoice {
	s.client = c
	return s
}

func (s SendInvoice) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendInvoice) WithApiBaseUrl(u string) *SendInvoice {
	s.baseUrl = u
	return s
}

func (s SendInvoice) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendInvoice) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendInvoice) Execute(token string) (*objects.Message, error) {
	return SendTelegramPostRequest[objects.Message](token, "sendInvoice", s)
}

type CreateInvoiceLink struct {
	Title                     string
	Description               string
	Payload                   string
	Currency                  string
	Prices                    []objects.LabeledPrice
	BusinessConnectionId      *string
	ProviderToken             *string
	SubscriptionPeriod        *int
	MaxTipAmount              *int
	SuggestedTipAmounts       *[]int
	ProviderData              *string
	PhotoUrl                  *string
	PhotoSize                 *int
	PhotoWidth                *int
	PhotoHeight               *int
	NeedName                  *bool
	NeedPhoneNumber           *bool
	NeedEmail                 *bool
	NeedShippingAddress       *bool
	SendPhoneNumberToProvider *bool
	SendEmailToProvider       *bool
	IsFlexible                *bool
	client                    *http.Client
	baseUrl                   string
}

func (s *CreateInvoiceLink) WithClient(c *http.Client) *CreateInvoiceLink {
	s.client = c
	return s
}

func (s CreateInvoiceLink) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *CreateInvoiceLink) WithApiBaseUrl(u string) *CreateInvoiceLink {
	s.baseUrl = u
	return s
}

func (s CreateInvoiceLink) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (c CreateInvoiceLink) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c CreateInvoiceLink) Execute(token string) (*string, error) {
	return SendTelegramPostRequest[string](token, "createInvoiceLink", c)
}

type AnswerShippingQuery struct {
	ShippingQueryId string
	Ok              bool
	ShippingOptions *[]objects.ShippingOption
	ErrorMessage    *string
	client          *http.Client
	baseUrl         string
}

func (s *AnswerShippingQuery) WithClient(c *http.Client) *AnswerShippingQuery {
	s.client = c
	return s
}

func (s AnswerShippingQuery) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *AnswerShippingQuery) WithApiBaseUrl(u string) *AnswerShippingQuery {
	s.baseUrl = u
	return s
}

func (s AnswerShippingQuery) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (a AnswerShippingQuery) ToRequestBody() ([]byte, error) {
	return json.Marshal(a)
}

func (a AnswerShippingQuery) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "answerShippingQuery", a)
}

type AnswerPreCheckoutQuery struct {
	PreCheckoutQueryId string
	Ok                 bool
	ErrorMessage       *bool
	client             *http.Client
	baseUrl            string
}

func (s *AnswerPreCheckoutQuery) WithClient(c *http.Client) *AnswerPreCheckoutQuery {
	s.client = c
	return s
}

func (s AnswerPreCheckoutQuery) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *AnswerPreCheckoutQuery) WithApiBaseUrl(u string) *AnswerPreCheckoutQuery {
	s.baseUrl = u
	return s
}

func (s AnswerPreCheckoutQuery) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (a AnswerPreCheckoutQuery) ToRequestBody() ([]byte, error) {
	return json.Marshal(a)
}

func (a AnswerPreCheckoutQuery) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "answerPreCheckoutQuery", a)
}

type GetStarTransactions struct {
	Offset  *int
	Limit   *int
	client  *http.Client
	baseUrl string
}

func (s *GetStarTransactions) WithClient(c *http.Client) *GetStarTransactions {
	s.client = c
	return s
}

func (s GetStarTransactions) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetStarTransactions) WithApiBaseUrl(u string) *GetStarTransactions {
	s.baseUrl = u
	return s
}

func (s GetStarTransactions) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (g GetStarTransactions) Validate() error {
	if g.Limit != nil {
		if *g.Limit < 1 || *g.Limit > 100 {
			return objects.ErrInvalidParam("limit parameter must be between 1 and 100")
		}
	}
	return nil
}

func (g GetStarTransactions) ToRequestBody() ([]byte, error) {
	return json.Marshal(g)
}

func (g GetStarTransactions) Execute(token string) (*objects.StarTransactions, error) {
	return SendTelegramGetRequest[objects.StarTransactions](token, "getStarTransactions", g)
}

type RefundStarPayment struct {
	// FIXME: all of the channel and supergroup ids SHOULD BE NEGATIVE,
	// should fix validation EVERYWHERE
	UserId                  int
	TelegramPaymentChargeId string
	client                  *http.Client
	baseUrl                 string
}

func (s *RefundStarPayment) WithClient(c *http.Client) *RefundStarPayment {
	s.client = c
	return s
}

func (s RefundStarPayment) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *RefundStarPayment) WithApiBaseUrl(u string) *RefundStarPayment {
	s.baseUrl = u
	return s
}

func (s RefundStarPayment) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (r RefundStarPayment) Validate() error {
	if r.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if strings.TrimSpace(r.TelegramPaymentChargeId) == "" {
		return objects.ErrInvalidParam("telegram_payment_charge_id parameter can't be empty")
	}
	return nil
}

func (r RefundStarPayment) ToRequestBody() ([]byte, error) {
	return json.Marshal(r)
}

func (r RefundStarPayment) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "refundStarPayment", r)
}

type EditUserStarSubscription struct {
	UserId                  int
	TelegramPaymentChargeId string
	IsCanceled              bool
	client                  *http.Client
	baseUrl                 string
}

func (s *EditUserStarSubscription) WithClient(c *http.Client) *EditUserStarSubscription {
	s.client = c
	return s
}

func (s EditUserStarSubscription) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *EditUserStarSubscription) WithApiBaseUrl(u string) *EditUserStarSubscription {
	s.baseUrl = u
	return s
}

func (s EditUserStarSubscription) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e EditUserStarSubscription) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditUserStarSubscription) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "editUserStarSubscription", e)
}
