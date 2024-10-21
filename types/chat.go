package types

type Chat struct {
	Id        int64   `json:"id"`
	Type      string  `json:"type"`
	Title     *string `json:"title,omitempty"`
	UserName  *string `json:"user_name,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	IsForum   *bool   `json:"is_forum,omitempty"`
}

type ChatFullInfo struct {
	Chat
	AccentColorId                      int                   `json:"accent_color_id"`
	MaxReactionCount                   int                   `json:"max_reaction_count"`
	Photo                              *ChatPhoto            `json:"photo,omitempty"`
	ActiveUsernames                    *[]string             `json:"active_usernames,omitempty"`
	BirthDate                          *BirthDate            `json:"birth_date,omitempty"`
	BusinessIntro                      *BusinessIntro        `json:"business_intro,omitempty"`
	BusinessLocation                   *BusinessLocation     `json:"business_location,omitempty"`
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours,omitempty"`
	PersonalChat                       *Chat                 `json:"personal_chat,omitempty"`
	AvailableReactions                 *[]ReactionType       `json:"available_reactions,omitempty"`
	BackgroundCustomEmojiId            *string               `json:"background_custom_emoji_id,omitempty"`
	ProfileAccentColorId               *bool                 `json:"profile_accent_color_id,omitempty"`
	ProfileBackgroundCustomEmojiId     *string               `json:"profile_background_custom_emoji_id,omitempty"`
	EmojiStatusCustomEmojiId           *string               `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate          *bool                 `json:"emoji_status_expiration_date,omitempty"`
	Bio                                *string               `json:"bio,omitempty"`
	HasPrivateForwards                 *bool                 `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages *bool                 `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages                 *bool                 `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      *bool                 `json:"join_by_request,omitempty"`
	Description                        *string               `json:"description,omitempty"`
	InviteLink                         *string               `json:"invite_link,omitempty"`
	PinnedMessage                      *Message              `json:"pinned_message,omitempty"`
	Permissions                        *ChatPermissions      `json:"permissions,omitempty"`
	CanSendPaidMedia                   *bool                 `json:"can_send_paid_media,omitempty"`
	SlowModeDelay                      *int                  `json:"slow_mode_delay,omitempty"`
	UnrestrictBoostCount               *int                  `json:"unrestrict_boost_count,omitempty"`
	MessageAutoDeleteTime              *int                  `json:"message_auto_delete_time,omitempty"`
	HasAggressiveAntiSpamEnabled       *string               `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                   *bool                 `json:"has_hidden_members,omitempty"`
	HasProtectedCount                  *bool                 `json:"has_protected_count,omitempty"`
	HasVisibleHistory                  *bool                 `json:"has_visible_history,omitempty"`
	StickerSetName                     *string               `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                   *bool                 `json:"can_set_sticker_set,omitempty"`
	CustomEmojiStickerSetName          *string               `json:"custom_emoji_sticker_set_name,omitempty"`
	LinkedChatId                       *int64                `json:"linked_chat_id,omitempty"`
	Location                           *ChatLocation         `json:"location,omitempty"`
}

type ChatBackground struct {
	Type BackgroundType `json:"type"`
}

type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

type ChatShared struct {
	RequestId string       `json:"request_id"`
	ChatId    int64        `json:"chat_id"`
	Title     *string      `json:"title,omitempty"`
	Username  *string      `json:"username,omitempty"`
	Photo     *[]PhotoSize `json:"photo,omitempty"`
}

type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}

type ChatInviteLink struct {
	InviteLink              string  `json:"invite_link"`
	Creator                 User    `json:"creator"`
	CreatesJoinRequest      bool    `json:"creates_join_request"`
	IsPrimary               bool    `json:"is_primary"`
	IsRevoked               bool    `json:"is_revoked"`
	Name                    *string `json:"name,omitempty"`
	ExpireDate              *int    `json:"expire_date,omitempty"`
	MemberLimit             *bool   `json:"member_limit,omitempty"`
	PendingJoinRequestCount *int    `json:"pending_join_request_count,omitempty"`
	SubscriptionPeriod      *int    `json:"subscription_period,omitempty"`
	SubscriptionPrice       *int    `json:"subscription_price,omitempty"`
}

type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`
	User       User            `json:"user"`
	UserChatId int64           `json:"user_chat_id"`
	Date       int             `json:"date"`
	Bio        *string         `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

type ChatPermissions struct {
	CanSendMessages       *bool `json:"can_send_messages,omitempty"`
	CanSendAudios         *bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      *bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         *bool `json:"can_send_photos,omitempty"`
	CanSendVideos         *bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     *bool `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     *bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          *bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  *bool `json:"can_send_other_messages,omitempty"`
	CanAddWebpagePreviews *bool `json:"can_add_webpage_previews,omitempty"`
	CanChangeInfo         *bool `json:"can_change_info,omitempty"`
	CanInviteUsers        *bool `json:"can_invite_users,omitempty"`
	CanPinMessages        *bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       *bool `json:"can_manage_topics,omitempty"`
}

type ChatAdministratorRights struct {
	IsAnonymous         bool  `json:"is_anonymous"`
	CanManageChat       bool  `json:"can_manage_chat"`
	CanDeleteMessages   bool  `json:"can_delete_messages"`
	CanManageVideoChats bool  `json:"can_manage_video_chats"`
	CanRestrictMembers  bool  `json:"can_restrict_members"`
	CanPromoteMembers   bool  `json:"can_promote_members"`
	CanChangeInfo       bool  `json:"can_change_info"`
	CanInviteUsers      bool  `json:"can_invite_users"`
	CanPostStories      bool  `json:"can_post_stories"`
	CanEditStories      bool  `json:"can_edit_stories"`
	CanDeleteStories    bool  `json:"can_delete_stories"`
	CanPostMessages     *bool `json:"can_post_messages,omitempty"`
	CanEditMessages     *bool `json:"can_edit_messages,omitempty"`
	CanPinMessages      *bool `json:"can_pin_messages,omitempty"`
	CanManageTopics     *bool `json:"can_manage_topics,omitempty"`
}
