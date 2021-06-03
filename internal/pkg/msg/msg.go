package msg

type CountMsg struct {
	MsgType int32
	UserId  int64
	AnimeId int64
}

const (
	MsgTypeLike          = 1
	MsgTypeLikeCancel    = 2
	MsgTypeUnlike        = 3
	MsgTypeUnlikeCancel  = 4
	MsgTypeAddView       = 5
	MsgTypeCollect       = 6
	MsgTypeCollectCancel = 7
)
