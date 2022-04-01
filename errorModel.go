package piesta

import "errors"

//DB
const (
	MongoConnectionCode = 100
	PGConnectionCode    = 101
	CountDocumentsCode  = 102
	InsertOneCode       = 103
	FindCode            = 104
	FindOneCode         = 105
	CurCode             = 106
	CurDeCodeCode       = 107
	UpdateOneCode       = 108
	UpdateManyCode      = 109
	DeleteManyCode      = 110
	DeleteOneCode       = 111
	DistinctCode        = 112
	GetEnvInfo          = 113
)

var ErrRecordNotFound = errors.New("record not found")

type ErrorStruct struct {
	Message error
	Code    int
}

// 정보 없음
func NoUserStatusStatusError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("There is no user_status data"),
		Code:    201,
	}
}

func NoSocketInfoError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("There is no socket_info data"),
		Code:    202,
	}
}

func NoRoomInfoError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("There is no room_info data"),
		Code:    203,
	}
}

func NoObjectInfoError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("There is no user_object_status data"),
		Code:    204,
	}
}

func NoChannelInfoError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("There is no channel data"),
		Code:    205,
	}
}

func NoObjectIdError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("There is no objId data"),
		Code:    208,
	}
}

func NoHostError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("There is no host data"),
		Code:    209,
	}
}

func NoSocketIDError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("There is no socketID"),
		Code:    210,
	}
}

func NoFileInfoError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("There is no file data"),
		Code:    211,
	}
}

func ObjectIDHexError() *ErrorStruct {
	return &ErrorStruct{
		Code: 212,
	}
}

func NoUserInfoError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("There is no user_info data"),
		Code:    213,
	}
}

func NoHashingPassword() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("There is no hashing password data"),
		Code:    210,
	}
}

// 중복
func DuplicateUidError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("The user's uid already exists."),
		Code:    301,
	}
}

func DuplicateRoomIdError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("The roomId already exists."),
		Code:    302,
	}
}

func DuplicateUidAndRoomIdError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("The roomId`s uid already exists."),
		Code:    303,
	}
}

func DuplicateChannelNameError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("The channelName already exists."),
		Code:    304,
	}
}

func DuplicateNicknameError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("pq: duplicate key value violates unique constraint \"user_info_nickname_key\""),
		Code:    306,
	}
}

func DuplicateSocketInfoError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("The socketInfo already exists."),
		Code:    307,
	}
}

func DuplicateUserIdError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("pq: duplicate key value violates unique constraint \"user_info_user_id_key\""),
		Code:    308,
	}
}

func DuplicatePhoneError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("pq: duplicate key value violates unique constraint \"user_info_phone_key\""),
		Code:    309,
	}
}

//unmatch
func DoseNotMatchUidError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("user uid dose not match"),
		Code:    401,
	}
}

func DoseNotMatchHostError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("This is not a request from the host."),
		Code:    402,
	}
}

func InvalidPasswordError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("Invalid user password"),
		Code:    403,
	}
}

//json
func JSONMarsharError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("Cannot be converted to JSON"),
		Code:    500,
	}
}

func ParseIntError() *ErrorStruct {
	return &ErrorStruct{
		Code: 501,
	}
}

func ReadJSONFileError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("Cannot read object.json file"),
		Code:    502,
	}
}

func JwtParseError() *ErrorStruct {
	return &ErrorStruct{
		Code: 503,
	}
}

//zero
func UidZeroError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("uid cannot be zero"),
		Code:    600,
	}
}

func RoomidZeroError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("roomid cannot be zero"),
		Code:    601,
	}
}

func UidAndRoomidZeroError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("uid and roomid cannot be zero"),
		Code:    602,
	}
}

func SocketIdNullError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("socketId cannot be empty"),
		Code:    603,
	}
}

func ObjectIdZeroError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("objId cannot be zero"),
		Code:    604,
	}
}

func ChannelIDZeroError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("channelID cannot be zero"),
		Code:    605,
	}
}

func ChannelNameNullError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("channelName cannot be empty"),
		Code:    606,
	}
}

func FileNameNullError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("fileName cannot be empty"),
		Code:    607,
	}
}

func OriginFileNameNullError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("originFileName cannot be empty"),
		Code:    608,
	}
}

func FileIdNullError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("FileId cannot be empty"),
		Code:    609,
	}
}

func NoticeTitleNullError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("Notice title cannot be empty"),
		Code:    610,
	}
}

func NoticeContentNullError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("Notice content cannot be empty"),
		Code:    611,
	}
}

func UserIdNullError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("userId cannot be empty"),
		Code:    612,
	}
}

// wrong value
func TooLongForTypeError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("pq: value too long for type"),
		Code:    700,
	}
}

// token
func SignedStringError() *ErrorStruct {
	return &ErrorStruct{
		Code: 800,
	}
}

func SigningMethodError() *ErrorStruct {
	return &ErrorStruct{
		Code: 801,
	}
}

func TokenClaimsError() *ErrorStruct {
	return &ErrorStruct{
		Message: errors.New("token claims error"),
		Code:    802,
	}
}

func NoAccessTokenError() *ErrorStruct {
	return &ErrorStruct{
		Code: 803,
	}
}

func NoRefreshTokenError() *ErrorStruct {
	return &ErrorStruct{
		Code: 804,
	}
}

// hashing password
func HashingPasswordError() *ErrorStruct {
	return &ErrorStruct{
		Code: 900,
	}
}
