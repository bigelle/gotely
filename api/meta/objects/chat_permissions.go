package objects

type ChatPermissions struct {
	CanSendMessages       bool
	CanSendAudios         bool
	CanSendDocuments      bool
	CanSendPhotos         bool
	CanSendVideos         bool
	CanSendVideoNotes     bool
	CanSendVoiceNotes     bool
	CanSendPolls          bool
	CanSendOtherMessages  bool
	CanAddWebpagePreviews bool
	CanChangeInfo         bool
	CanInviteUsers        bool
	CanPinMessages        bool
	CanManageTopics       bool
}
