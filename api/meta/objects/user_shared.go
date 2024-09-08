package objects

type UserShared struct {
	UserId    int64
	FirstName string
	LastName  string
	Username  string
	Photo     []PhotoSize
}

type UsersShared struct {
	RequestId string
	Users     []UserShared
}
