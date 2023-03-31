package db

import "time"

type Register struct {
	Account        string `gorm:"column:account;primary_key;type:char(255)" json:"account"`
	Password       string `gorm:"column:password;type:varchar(255)" json:"password"`
	Ex             string `gorm:"column:ex;size:1024" json:"ex"`
	UserID         string `gorm:"column:user_id;type:varchar(255)" json:"userID"`
	AreaCode       string `gorm:"column:area_code;type:varchar(255)"`
	InvitationCode string `gorm:"column:invitation_code;type:varchar(255)"`
	RegisterIP     string `gorm:"column:register_ip;type:varchar(255)"`
}

type Invitation struct {
	InvitationCode string    `gorm:"column:invitation_code;primary_key;type:varchar(32)"`
	CreateTime     time.Time `gorm:"column:create_time"`
	UserID         string    `gorm:"column:user_id;index:userID"`
	LastTime       time.Time `gorm:"column:last_time"`
	Status         int32     `gorm:"column:status"`
}

// message FriendInfo{
// string OwnerUserID = 1;
// string Remark = 2;
// int64 CreateTime = 3;
// UserInfo FriendUser = 4;
// int32 AddSource = 5;
// string OperatorUserID = 6;
// string Ex = 7;
// }
// open_im_sdk.FriendInfo(FriendUser) != imdb.Friend(FriendUserID)
type Friend struct {
	OwnerUserID    string    `gorm:"column:owner_user_id;primary_key;size:64"`
	FriendUserID   string    `gorm:"column:friend_user_id;primary_key;size:64"`
	Remark         string    `gorm:"column:remark;size:255"`
	CreateTime     time.Time `gorm:"column:create_time"`
	AddSource      int32     `gorm:"column:add_source"`
	OperatorUserID string    `gorm:"column:operator_user_id;size:64"`
	Ex             string    `gorm:"column:ex;size:1024"`
}

// message FriendRequest{
// string  FromUserID = 1;
// string ToUserID = 2;
// int32 HandleResult = 3;
// string ReqMsg = 4;
// int64 CreateTime = 5;
// string HandlerUserID = 6;
// string HandleMsg = 7;
// int64 HandleTime = 8;
// string Ex = 9;
// }
// open_im_sdk.FriendRequest(nickname, farce url ...) != imdb.FriendRequest
type FriendRequest struct {
	FromUserID    string    `gorm:"column:from_user_id;primary_key;size:64"`
	ToUserID      string    `gorm:"column:to_user_id;primary_key;size:64"`
	HandleResult  int32     `gorm:"column:handle_result"`
	ReqMsg        string    `gorm:"column:req_msg;size:255"`
	CreateTime    time.Time `gorm:"column:create_time"`
	HandlerUserID string    `gorm:"column:handler_user_id;size:64"`
	HandleMsg     string    `gorm:"column:handle_msg;size:255"`
	HandleTime    time.Time `gorm:"column:handle_time"`
	Ex            string    `gorm:"column:ex;size:1024"`
}

func (FriendRequest) TableName() string {
	return "friend_requests"
}

//	message GroupInfo{
//	 string GroupID = 1;
//	 string GroupName = 2;
//	 string Notification = 3;
//	 string Introduction = 4;
//	 string FaceUrl = 5;
//	 string OwnerUserID = 6;
//	 uint32 MemberCount = 8;
//	 int64 CreateTime = 7;
//	 string Ex = 9;
//	 int32 Status = 10;
//	 string CreatorUserID = 11;
//	 int32 GroupType = 12;
//	}
//
// open_im_sdk.GroupInfo (OwnerUserID ,  MemberCount )> imdb.Group
type Group struct {
	//`json:"operationID" binding:"required"`
	//`protobuf:"bytes,1,opt,name=GroupID" json:"GroupID,omitempty"` `json:"operationID" binding:"required"`
	GroupID                string    `gorm:"column:group_id;primary_key;size:64" json:"groupID" binding:"required"`
	GroupName              string    `gorm:"column:name;size:255" json:"groupName"`
	Notification           string    `gorm:"column:notification;size:255" json:"notification"`
	Introduction           string    `gorm:"column:introduction;size:255" json:"introduction"`
	FaceURL                string    `gorm:"column:face_url;size:255" json:"faceURL"`
	CreateTime             time.Time `gorm:"column:create_time;index:create_time"`
	Ex                     string    `gorm:"column:ex" json:"ex;size:1024" json:"ex"`
	Status                 int32     `gorm:"column:status"`
	CreatorUserID          string    `gorm:"column:creator_user_id;size:64"`
	GroupType              int32     `gorm:"column:group_type"`
	NeedVerification       int32     `gorm:"column:need_verification"`
	LookMemberInfo         int32     `gorm:"column:look_member_info" json:"lookMemberInfo"`
	ApplyMemberFriend      int32     `gorm:"column:apply_member_friend" json:"applyMemberFriend"`
	NotificationUpdateTime time.Time `gorm:"column:notification_update_time"`
	NotificationUserID     string    `gorm:"column:notification_user_id;size:64"`
}

// message GroupMemberFullInfo {
// string GroupID = 1 ;
// string UserID = 2 ;
// int32 roleLevel = 3;
// int64 JoinTime = 4;
// string NickName = 5;
// string FaceUrl = 6;
// int32 JoinSource = 8;
// string OperatorUserID = 9;
// string Ex = 10;
// int32 AppMangerLevel = 7; //if >0
// }  open_im_sdk.GroupMemberFullInfo(AppMangerLevel) > imdb.GroupMember
type GroupMember struct {
	GroupID        string    `gorm:"column:group_id;primary_key;size:64"`
	UserID         string    `gorm:"column:user_id;primary_key;size:64"`
	Nickname       string    `gorm:"column:nickname;size:255"`
	FaceURL        string    `gorm:"column:user_group_face_url;size:255"`
	RoleLevel      int32     `gorm:"column:role_level"`
	JoinTime       time.Time `gorm:"column:join_time"`
	JoinSource     int32     `gorm:"column:join_source"`
	InviterUserID  string    `gorm:"column:inviter_user_id;size:64"`
	OperatorUserID string    `gorm:"column:operator_user_id;size:64"`
	MuteEndTime    time.Time `gorm:"column:mute_end_time"`
	Ex             string    `gorm:"column:ex;size:1024"`
}

// message GroupRequest{
// string UserID = 1;
// string GroupID = 2;
// string HandleResult = 3;
// string ReqMsg = 4;
// string  HandleMsg = 5;
// int64 ReqTime = 6;
// string HandleUserID = 7;
// int64 HandleTime = 8;
// string Ex = 9;
// }open_im_sdk.GroupRequest == imdb.GroupRequest
type GroupRequest struct {
	UserID        string    `gorm:"column:user_id;primary_key;size:64"`
	GroupID       string    `gorm:"column:group_id;primary_key;size:64"`
	HandleResult  int32     `gorm:"column:handle_result"`
	ReqMsg        string    `gorm:"column:req_msg;size:1024"`
	HandledMsg    string    `gorm:"column:handle_msg;size:1024"`
	ReqTime       time.Time `gorm:"column:req_time"`
	HandleUserID  string    `gorm:"column:handle_user_id;size:64"`
	HandledTime   time.Time `gorm:"column:handle_time"`
	JoinSource    int32     `gorm:"column:join_source"`
	InviterUserID string    `gorm:"column:inviter_user_id;size:64"`
	Ex            string    `gorm:"column:ex;size:1024"`
}

// string UserID = 1;
// string Nickname = 2;
// string FaceUrl = 3;
// int32 Gender = 4;
// string PhoneNumber = 5;
// string Birth = 6;
// string Email = 7;
// string Ex = 8;
// string CreateIp = 9;
// int64 CreateTime = 10;
// int32 AppMangerLevel = 11;
// open_im_sdk.User == imdb.User
type User struct {
	UserID           string    `gorm:"column:user_id;primary_key;size:64"`
	Nickname         string    `gorm:"column:name;size:255"`
	FaceURL          string    `gorm:"column:face_url;size:255"`
	Gender           int32     `gorm:"column:gender"`
	PhoneNumber      string    `gorm:"column:phone_number;size:32"`
	Birth            time.Time `gorm:"column:birth"`
	Email            string    `gorm:"column:email;size:64"`
	Ex               string    `gorm:"column:ex;size:1024"`
	status           int32     `gorm:"column:status"`
	AppMangerLevel   int32     `gorm:"column:app_manger_level"`
	GlobalRecvMsgOpt int32     `gorm:"column:global_recv_msg_opt"`
	CreateTime       time.Time `gorm:"column:create_time;index:create_time"`
}

type UserIpRecord struct {
	UserID        string    `gorm:"column:user_id;primary_key;size:64"`
	CreateIp      string    `gorm:"column:create_ip;size:15"`
	LastLoginTime time.Time `gorm:"column:last_login_time"`
	LastLoginIp   string    `gorm:"column:last_login_ip;size:15"`
	LoginTimes    int32     `gorm:"column:login_times"`
}

// ip limit login
type IpLimit struct {
	Ip            string    `gorm:"column:ip;primary_key;size:15"`
	LimitRegister int32     `gorm:"column:limit_register;size:1"`
	LimitLogin    int32     `gorm:"column:limit_login;size:1"`
	CreateTime    time.Time `gorm:"column:create_time"`
	LimitTime     time.Time `gorm:"column:limit_time"`
}

// ip login
type UserIpLimit struct {
	UserID     string    `gorm:"column:user_id;primary_key;size:64"`
	Ip         string    `gorm:"column:ip;primary_key;size:15"`
	CreateTime time.Time `gorm:"column:create_time"`
}

// message BlackInfo{
// string OwnerUserID = 1;
// int64 CreateTime = 2;
// PublicUserInfo BlackUserInfo = 4;
// int32 AddSource = 5;
// string OperatorUserID = 6;
// string Ex = 7;
// }
// open_im_sdk.BlackInfo(BlackUserInfo) != imdb.Black (BlockUserID)
type Black struct {
	OwnerUserID    string    `gorm:"column:owner_user_id;primary_key;size:64"`
	BlockUserID    string    `gorm:"column:block_user_id;primary_key;size:64"`
	CreateTime     time.Time `gorm:"column:create_time"`
	AddSource      int32     `gorm:"column:add_source"`
	OperatorUserID string    `gorm:"column:operator_user_id;size:64"`
	Ex             string    `gorm:"column:ex;size:1024"`
}

type ChatLog struct {
	ServerMsgID      string    `gorm:"column:server_msg_id;primary_key;type:char(64)" json:"serverMsgID"`
	ClientMsgID      string    `gorm:"column:client_msg_id;type:char(64)" json:"clientMsgID"`
	SendID           string    `gorm:"column:send_id;type:char(64);index:send_id,priority:2" json:"sendID"`
	RecvID           string    `gorm:"column:recv_id;type:char(64);index:recv_id,priority:2" json:"recvID"`
	SenderPlatformID int32     `gorm:"column:sender_platform_id" json:"senderPlatformID"`
	SenderNickname   string    `gorm:"column:sender_nick_name;type:varchar(255)" json:"senderNickname"`
	SenderFaceURL    string    `gorm:"column:sender_face_url;type:varchar(255);" json:"senderFaceURL"`
	SessionType      int32     `gorm:"column:session_type;index:session_type,priority:2;index:session_type_alone" json:"sessionType"`
	MsgFrom          int32     `gorm:"column:msg_from" json:"msgFrom"`
	ContentType      int32     `gorm:"column:content_type;index:content_type,priority:2;index:content_type_alone" json:"contentType"`
	Content          string    `gorm:"column:content;type:varchar(3000)" json:"content"`
	Status           int32     `gorm:"column:status" json:"status"`
	SendTime         time.Time `gorm:"column:send_time;index:sendTime;index:content_type,priority:1;index:session_type,priority:1;index:recv_id,priority:1;index:send_id,priority:1" json:"sendTime"`
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`
	Ex               string    `gorm:"column:ex;type:varchar(1024)" json:"ex"`
}

func (ChatLog) TableName() string {
	return "chat_logs"
}

type BlackList struct {
	UserId           string    `gorm:"column:uid"`
	BeginDisableTime time.Time `gorm:"column:begin_disable_time"`
	EndDisableTime   time.Time `gorm:"column:end_disable_time"`
}
type Conversation struct {
	OwnerUserID           string `gorm:"column:owner_user_id;primary_key;type:char(128)" json:"OwnerUserID"`
	ConversationID        string `gorm:"column:conversation_id;primary_key;type:char(128)" json:"conversationID"`
	ConversationType      int32  `gorm:"column:conversation_type" json:"conversationType"`
	UserID                string `gorm:"column:user_id;type:char(64)" json:"userID"`
	GroupID               string `gorm:"column:group_id;type:char(128)" json:"groupID"`
	RecvMsgOpt            int32  `gorm:"column:recv_msg_opt" json:"recvMsgOpt"`
	UnreadCount           int32  `gorm:"column:unread_count" json:"unreadCount"`
	DraftTextTime         int64  `gorm:"column:draft_text_time" json:"draftTextTime"`
	IsPinned              bool   `gorm:"column:is_pinned" json:"isPinned"`
	IsPrivateChat         bool   `gorm:"column:is_private_chat" json:"isPrivateChat"`
	BurnDuration          int32  `gorm:"column:burn_duration;default:30" json:"burnDuration"`
	GroupAtType           int32  `gorm:"column:group_at_type" json:"groupAtType"`
	IsNotInGroup          bool   `gorm:"column:is_not_in_group" json:"isNotInGroup"`
	UpdateUnreadCountTime int64  `gorm:"column:update_unread_count_time" json:"updateUnreadCountTime"`
	AttachedInfo          string `gorm:"column:attached_info;type:varchar(1024)" json:"attachedInfo"`
	Ex                    string `gorm:"column:ex;type:varchar(1024)" json:"ex"`
}

func (Conversation) TableName() string {
	return "conversations"
}

type Department struct {
	DepartmentID   string    `gorm:"column:department_id;primary_key;size:64" json:"departmentID"`
	FaceURL        string    `gorm:"column:face_url;size:255" json:"faceURL"`
	Name           string    `gorm:"column:name;size:256" json:"name" binding:"required"`
	ParentID       string    `gorm:"column:parent_id;size:64" json:"parentID" binding:"required"` // "0" or Real parent id
	Order          int32     `gorm:"column:order" json:"order" `                                  // 1, 2, ...
	DepartmentType int32     `gorm:"column:department_type" json:"departmentType"`                //1, 2...
	RelatedGroupID string    `gorm:"column:related_group_id;size:64" json:"relatedGroupID"`
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`
	Ex             string    `gorm:"column:ex;type:varchar(1024)" json:"ex"`
}

func (Department) TableName() string {
	return "departments"
}

type OrganizationUser struct {
	UserID      string    `gorm:"column:user_id;primary_key;size:64"`
	Nickname    string    `gorm:"column:nickname;size:256"`
	EnglishName string    `gorm:"column:english_name;size:256"`
	FaceURL     string    `gorm:"column:face_url;size:256"`
	Gender      int32     `gorm:"column:gender"` //1 ,2
	Mobile      string    `gorm:"column:mobile;size:32"`
	Telephone   string    `gorm:"column:telephone;size:32"`
	Birth       time.Time `gorm:"column:birth"`
	Email       string    `gorm:"column:email;size:64"`
	CreateTime  time.Time `gorm:"column:create_time"`
	Ex          string    `gorm:"column:ex;size:1024"`
}

func (OrganizationUser) TableName() string {
	return "organization_users"
}

type DepartmentMember struct {
	UserID       string    `gorm:"column:user_id;primary_key;size:64"`
	DepartmentID string    `gorm:"column:department_id;primary_key;size:64"`
	Order        int32     `gorm:"column:order" json:"order"` //1,2
	Position     string    `gorm:"column:position;size:256" json:"position"`
	Leader       int32     `gorm:"column:leader" json:"leader"` //-1, 1
	Status       int32     `gorm:"column:status" json:"status"` //-1, 1
	CreateTime   time.Time `gorm:"column:create_time"`
	Ex           string    `gorm:"column:ex;type:varchar(1024)" json:"ex"`
}

func (DepartmentMember) TableName() string {
	return "department_members"
}

type AppVersion struct {
	Version     string `gorm:"column:version;size:64" json:"version"`
	Type        int    `gorm:"column:type;primary_key" json:"type"`
	UpdateTime  int    `gorm:"column:update_time" json:"update_time"`
	ForceUpdate bool   `gorm:"column:force_update" json:"force_update"`
	FileName    string `gorm:"column:file_name" json:"file_name"`
	YamlName    string `gorm:"column:yaml_name" json:"yaml_name"`
	UpdateLog   string `gorm:"column:update_log" json:"update_log"`
}

func (AppVersion) TableName() string {
	return "app_version"
}

type RegisterAddFriend struct {
	UserID string `gorm:"column:user_id;primary_key;size:64"`
}

func (RegisterAddFriend) TableName() string {
	return "register_add_friend"
}

type ClientInitConfig struct {
	DiscoverPageURL string `gorm:"column:discover_page_url;size:64" json:"version"`
}

func (ClientInitConfig) TableName() string {
	return "client_init_config"
}

type FNcountAccount struct {
	Id              int32     `gorm:"column:id" json:"id"`
	UserId          string    `gorm:"column:user_id" json:"user_id"`                     //用户id
	MainAccountId   string    `gorm:"column:main_account_id" json:"main_account_id"`     //主账号id
	PacketAccountId string    `gorm:"column:packet_account_id" json:"packet_account_id"` //红包账户id
	Mobile          string    `gorm:"column:mobile" json:"mobile"`                       //手机号码
	RealName        string    `gorm:"column:realname" json:"realname"`                   //身份证
	IdCard          string    `gorm:"column:id_card" json:"id_card"`                     //身份证
	PaySwitch       int32     `gorm:"column:pay_switch" json:"pay_switch"`               //支付开关(0关闭、1默认开启)
	BodPaySwitch    int32     `gorm:"column:bod_pay_switch" json:"bod_pay_switch"`       //指纹支付/人脸支付开关(0默认关闭、1开启)
	PaymentPassword string    `gorm:"column:payment_password" json:"payment_password"`   //支付密码(md5加密)
	OpenStatus      int32     `gorm:"column:open_status" json:"open_status"`             //开通状态
	OpenStep        int32     `gorm:"column:open_step" json:"open_step"`                 //开通认证步骤(1身份证认证、2支付密码、3绑定银行卡或者刷脸)
	CreatedTime     time.Time `gorm:"column:created_time" json:"created_time"`           //
	UpdatedTime     time.Time `gorm:"column:updated_time" json:"updated_time"`           //
}

func (FNcountAccount) TableName() string {
	return "f_ncount_account"
}

/*  `ncount_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '账户类型(1主账户，2红包账户)',*/
// 用户银行卡绑定表
type FNcountBankCard struct {
	Id                int32     `gorm:"column:id" json:"id"`
	UserId            string    `gorm:"column:user_id" json:"user_id"`                         //用户id
	NcountUserId      string    `gorm:"column:ncount_user_id" json:"ncount_user_id"`           //新生支付用户id
	MerOrderId        string    `gorm:"column:mer_order_id" json:"mer_order_id"`               //平台订单号
	NcountOrderId     string    `gorm:"column:ncount_order_id" json:"ncount_order_id"`         //第三方签约订单号
	BindCardAgrNo     string    `gorm:"column:bind_card_agr_no" json:"bind_card_agr_no"`       //第三方绑卡协议号
	NcountType        int       `gorm:"column:ncount_type" json:"ncount_type"`                 //账户类型(1主账户，2红包账户)
	Mobile            string    `gorm:"column:mobile" json:"mobile"`                           //手机号码
	CardOwner         string    `gorm:"column:card_owner" json:"card_owner"`                   //持卡者名字
	BankCardNumber    string    `gorm:"column:bank_card_number" json:"bank_card_number"`       //银行卡号
	CardAvailableDate string    `gorm:"column:card_available_date" json:"card_available_date"` //信用卡有效期
	Cvv2              string    `gorm:"column:cvv2" json:"cvv2"`                               //cvv2
	BankCode          string    `gorm:"column:bank_code" json:"bank_code"`                     //银行简写
	IsBind            int       `gorm:"column:is_bind" json:"is_bind"`                         //是否绑定成功(0预提交、1绑定成功)
	IsDelete          int       `gorm:"column:is_delete" json:"is_delete"`                     //是否删除(0未删除，1已删除)
	CreatedTime       time.Time `gorm:"column:created_time" json:"created_time"`               //
	UpdatedTime       time.Time `gorm:"column:updated_time" json:"updated_time"`               //
}

func (FNcountBankCard) TableName() string {
	return "f_ncount_bank_card"
}
