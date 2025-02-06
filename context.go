package gotely

import (
	"net/http"

	"github.com/bigelle/gotely/methods"
	"github.com/bigelle/gotely/objects"
)

type Context struct {
	Update     objects.Update
	Token      string
	Client     *http.Client
	ApiBaseUrl string
}

func (c Context) SendMessage(m methods.SendMessage) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) ForwardMessage(m methods.ForwardMessage) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) ForwardMessages(m methods.ForwardMessages) (*[]objects.MessageId, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) CopyMessage(m methods.CopyMessage) (*objects.MessageId, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) CopyMessages(m methods.CopyMessages) (*[]objects.MessageId, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendPhoto(m methods.SendPhoto) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendAudio(m methods.SendAudio) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendDocument(m methods.SendDocument) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendVideo(m methods.SendVideo) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendAnimation(m methods.SendAnimation) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendVoice(m methods.SendVoice) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendVideoNote(m methods.SendVideoNote) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendPaidMedia(m methods.SendPaidMedia) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendMediaGroup(m methods.SendMediaGroup) (*[]objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendLocation(m methods.SendLocation) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendVenue(m methods.SendVenue) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendContact(m methods.SendContact) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendPoll(m methods.SendPoll) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendDice(m methods.SendMessage) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendChatAction(m methods.SendChatAction) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetMessageReaction(m methods.SetMessageReaction) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetUserProfilePhotos(m methods.GetUserProfilePhotos) (*objects.UserProfilePhotos, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetUserEmojiStatus(m methods.SetUserEmojiStatus) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetFile(m methods.GetFile) (*objects.File, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) BanChatMember(m methods.BanChatMember) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) UnbanChatMember(m methods.UnbanChatMember) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) RestrictChatMember(m methods.RestrictChatMember) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) PromoteChatMember(m methods.PromoteChatMember) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetChatAdministratorCustomTitle(m methods.SetChatAdministratorCustomTitle) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) BanChatSenderChat(m methods.BanChatSenderChat) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) UnbanChatSenderChat(m methods.UnbanChatSenderChat) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetChatPermissions(m methods.SetChatPermissions) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) ExportChatInviteLink(m methods.ExportChatInviteLink) (*string, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) CreateInviteLink(m methods.CreateInviteLink) (*objects.ChatInviteLink, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) EditChatInviteLink(m methods.EditChatInviteLink) (*objects.ChatInviteLink, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) CreateChatSubscriptionInviteLink(m methods.CreateChatSubscriptionInviteLink) (*objects.ChatInviteLink, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) EditChatSubscriptionInviteLink(m methods.EditChatSubscriptionInviteLink) (*objects.ChatInviteLink, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) RevokeInviteLink(m methods.RevokeInviteLink) (*objects.ChatInviteLink, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) ApproveChatJoinRequest(m methods.ApproveChatJoinRequest) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) DeclineChatJoinRequest(m methods.DeclineChatJoinRequest) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetChatPhoto(m methods.SetChatPhoto) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) DeleteChatPhoto(m methods.DeleteChatPhoto) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetChatTitle(m methods.SetChatTitle) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetChatDescription(m methods.SetChatDescription) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) PinChatMessage(m methods.PinChatMessage) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) UnpinChatMessage(m methods.UnpinChatMessage) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) UnpinAllChatMessages(m methods.UnpinAllChatMessages) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) LeaveChat(m methods.LeaveChat) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetChat(m methods.GetChat) (*objects.ChatFullInfo, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetChatAdministrators(m methods.GetChatAdministrators) (*[]objects.ChatMember, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetChatMemberCount(m methods.GetChatMemberCount) (*int, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetChatStickerSet(m methods.SetChatStickerSet) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) DeleteChatStickerSet(m methods.DeleteChatStickerSet) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetForumTopicIconStickers(m methods.GetForumTopicIconStickers) (*[]objects.Sticker, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) CreateForumTopic(m methods.CreateForumTopic) (*objects.ForumTopic, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) EditForumTopic(m methods.EditForumTopic) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) CloseForumTopic(m methods.CloseForumTopic) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) ReopenForumTopic(m methods.ReopenForumTopic) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) DeleteForumTopic(m methods.DeleteForumTopic) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) UnpinAllForumTopicMessages(m methods.UnpinAllForumTopicMessages) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) EditGeneralForumTopic(m methods.EditGeneralForumTopic) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) CloseGeneralForumTopic(m methods.CloseGeneralForumTopic) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) ReopenGeneralForumTopic(m methods.ReopenGeneralForumTopic) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) HideGeneralForumTopic(m methods.HideGeneralForumTopic) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) UnhideGeneralForumTopic(m methods.UnhideGeneralForumTopic) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) UnpinAllGeneralForumTopicMessages(m methods.UnpinAllGeneralForumTopicMessages) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) AnswerCallbackQuery(m methods.AnswerCallbackQuery) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetUserChatBoosts(m methods.GetUserChatBoosts) (*objects.UserChatBoosts, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetBusinessConnection(m methods.GetBusinessConnection) (*objects.BusinessConnection, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetMyCommands(m methods.SetMyCommands) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) DeleteMyCommands(m methods.DeleteMyCommands) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetMyCommands(m methods.GetMyCommands) (*[]objects.BotCommand, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetMyName(m methods.SetMyName) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetMyName(m methods.GetMyName) (*objects.BotName, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetMyDescription(m methods.SetMyDescription) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetMyDescription(m methods.GetMyDescription) (*objects.BotDescription, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetMyShortDescription(m methods.SetMyShortDescription) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetMyShortDescription(m methods.GetMyShortDescription) (*objects.BotShortDescription, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetChatMenuButton(m methods.SetChatMenuButton) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetChatMenuButton(m methods.GetChatMenuButton) (*objects.MenuButton, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetMyDefaultAdministratorRights(m methods.SetMyDefaultAdministratorRights) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetMyDefaultAdministratorRights(m methods.GetMyDefaultAdministratorRights) (*objects.ChatAdministratorRights, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendGame(m methods.SendGame) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetGameHighScore(m methods.SetGameHighScore) (methods.MessageOrBool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetGameHighScores(m methods.GetGameHighScores) (*[]objects.GameHighScore, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) AnswerInlineQuery(m methods.AnswerInlineQuery) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) AnswerWebAppQuery(m methods.AnswerWebAppQuery) (*objects.SentWebAppMessage, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SavePreparedInlineMessage(m methods.SavePreparedInlineMessage) (*objects.PreparedInlineMessage, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetPassportDataErrors(m methods.SetPassportDataErrors) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendInvoice(m methods.SendInvoice) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) CreateInvoiceLink(m methods.CreateInvoiceLink) (*string, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) AnswerShippingQuery(m methods.AnswerShippingQuery) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) AnswerPreCheckoutQuery(m methods.AnswerPreCheckoutQuery) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetStarTransactions(m methods.GetStarTransactions) (*objects.StarTransactions, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) RefundStarPayment(m methods.RefundStarPayment) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) EditUserStarSubscription(m methods.EditUserStarSubscription) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendSticker(m methods.SendSticker) (*objects.Message, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetStickerSet(m methods.GetStickerSet) (*objects.StickerSet, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetCustomEmojiStickers(m methods.GetCustomEmojiStickers) (*[]objects.Sticker, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) UploadStickerFile(m methods.UploadStickerFile) (*objects.File, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) CreateNewStickerSet(m methods.CreateNewStickerSet) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) AddStickerToSet(m methods.AddStickerToSet) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetStickerPositionInSet(m methods.SetStickerPositionInSet) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) DeleteStickerFromSet(m methods.DeleteStickerFromSet) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) ReplaceStickerInSet(m methods.ReplaceStickerInSet) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetStickerEmojiList(m methods.SetStickerEmojiList) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetStickerKeywords(m methods.SetStickerKeywords) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetStickerMaskPosition(m methods.SetStickerMaskPosition) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetStickerSetTitle(m methods.SetStickerSetTitle) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetStickerSetThumbnail(m methods.SetStickerSetThumbnail) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SetCustomEmojiStickerSetThumbnail(m methods.SetCustomEmojiStickerSetThumbnail) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) DeleteStickerSet(m methods.DeleteStickerSet) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) GetAvailableGifts(m methods.GetAvailableGifts) (*objects.Gifts, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) SendGift(m methods.SendGift) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) EditMessageText(m methods.EditMessageText) (methods.MessageOrBool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) EditMessageCaption(m methods.EditMessageCaption) (methods.MessageOrBool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) EditMessageMedia(m methods.EditMessageMedia) (methods.MessageOrBool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) EditMessageLiveLocation(m methods.EditMessageLiveLocation) (methods.MessageOrBool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) StopMessageLiveLocation(m methods.StopMessageLiveLocation) (methods.MessageOrBool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) EditMessageReplyMarkup(m methods.EditMessageReplyMarkup) (methods.MessageOrBool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) StopPoll(m methods.StopPoll) (*objects.Poll, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) DeleteMessage(m methods.DeleteMessage) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}

func (c Context) DeleteMessages(m methods.DeleteMessages) (*bool, error) {
	return m.WithApiBaseUrl(c.ApiBaseUrl).WithClient(c.Client).Execute(c.Token)
}
