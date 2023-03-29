package api

import (
	apiAuth "Open_IM/internal/api/auth"
	clientInit "Open_IM/internal/api/client_init"
	"Open_IM/internal/api/cloud_wallet/account"
	"Open_IM/internal/api/conversation"
	"Open_IM/internal/api/friend"
	"Open_IM/internal/api/group"
	"Open_IM/internal/api/manage"
	apiChat "Open_IM/internal/api/msg"
	"Open_IM/internal/api/office"
	"Open_IM/internal/api/organization"
	apiThird "Open_IM/internal/api/third"
	"Open_IM/internal/api/user"
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/constant"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"os"

	promePkg "Open_IM/pkg/common/prometheus"

	"github.com/gin-gonic/gin"
)

func NewGinRouter() *gin.Engine {
	log.NewPrivateLog(constant.LogFileName)
	gin.SetMode(gin.DebugMode)
	f, _ := os.Create("./logs/api.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(utils.CorsHandler())
	log.Info("load config: ", config.Config)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if config.Config.Prometheus.Enable {
		promePkg.NewApiRequestCounter()
		promePkg.NewApiRequestFailedCounter()
		promePkg.NewApiRequestSuccessCounter()
		r.Use(promePkg.PromeTheusMiddleware)
		r.GET("/metrics", promePkg.PrometheusHandler())
	}

	// user routing group, which handles user registration and login services
	userRouterGroup := r.Group("/user")
	{
		userRouterGroup.POST("/update_user_info", user.UpdateUserInfo) //1
		userRouterGroup.POST("/set_global_msg_recv_opt", user.SetGlobalRecvMessageOpt)
		userRouterGroup.POST("/get_users_info", user.GetUsersPublicInfo)            //1
		userRouterGroup.POST("/get_self_user_info", user.GetSelfUserInfo)           //1
		userRouterGroup.POST("/get_users_online_status", user.GetUsersOnlineStatus) //1
		userRouterGroup.POST("/get_users_info_from_cache", user.GetUsersInfoFromCache)
		userRouterGroup.POST("/get_user_friend_from_cache", user.GetFriendIDListFromCache)
		userRouterGroup.POST("/get_black_list_from_cache", user.GetBlackIDListFromCache)
		userRouterGroup.POST("/get_all_users_uid", manage.GetAllUsersUid) //1
		userRouterGroup.POST("/account_check", manage.AccountCheck)       //1
		//	userRouterGroup.POST("/get_users_online_status", manage.GetUsersOnlineStatus) //1
		userRouterGroup.POST("/get_users", user.GetUsers)
	}
	//friend routing group
	friendRouterGroup := r.Group("/friend")
	{
		//	friendRouterGroup.POST("/get_friends_info", friend.GetFriendsInfo)
		friendRouterGroup.POST("/add_friend", friend.AddFriend)                              //1
		friendRouterGroup.POST("/delete_friend", friend.DeleteFriend)                        //1
		friendRouterGroup.POST("/get_friend_apply_list", friend.GetFriendApplyList)          //1
		friendRouterGroup.POST("/get_self_friend_apply_list", friend.GetSelfFriendApplyList) //1
		friendRouterGroup.POST("/get_friend_list", friend.GetFriendList)                     //1
		friendRouterGroup.POST("/add_friend_response", friend.AddFriendResponse)             //1
		friendRouterGroup.POST("/set_friend_remark", friend.SetFriendRemark)                 //1

		friendRouterGroup.POST("/add_black", friend.AddBlack)          //1
		friendRouterGroup.POST("/get_black_list", friend.GetBlacklist) //1
		friendRouterGroup.POST("/remove_black", friend.RemoveBlack)    //1

		friendRouterGroup.POST("/import_friend", friend.ImportFriend) //1
		friendRouterGroup.POST("/is_friend", friend.IsFriend)         //1
	}
	//group related routing group
	groupRouterGroup := r.Group("/group")
	{
		groupRouterGroup.POST("/create_group", group.CreateGroup)                                   //1
		groupRouterGroup.POST("/set_group_info", group.SetGroupInfo)                                //1
		groupRouterGroup.POST("/join_group", group.JoinGroup)                                       //1
		groupRouterGroup.POST("/quit_group", group.QuitGroup)                                       //1
		groupRouterGroup.POST("/group_application_response", group.ApplicationGroupResponse)        //1
		groupRouterGroup.POST("/transfer_group", group.TransferGroupOwner)                          //1
		groupRouterGroup.POST("/get_recv_group_applicationList", group.GetRecvGroupApplicationList) //1
		groupRouterGroup.POST("/get_user_req_group_applicationList", group.GetUserReqGroupApplicationList)
		groupRouterGroup.POST("/get_groups_info", group.GetGroupsInfo) //1
		groupRouterGroup.POST("/kick_group", group.KickGroupMember)    //1
		//	groupRouterGroup.POST("/get_group_member_list", group.GetGroupMemberList)        //no use
		groupRouterGroup.POST("/get_group_all_member_list", group.GetGroupAllMemberList) //1
		groupRouterGroup.POST("/get_group_members_info", group.GetGroupMembersInfo)      //1
		groupRouterGroup.POST("/invite_user_to_group", group.InviteUserToGroup)          //1
		//only for supergroup
		groupRouterGroup.POST("/invite_user_to_groups", group.InviteUserToGroups)
		groupRouterGroup.POST("/get_joined_group_list", group.GetJoinedGroupList)
		groupRouterGroup.POST("/dismiss_group", group.DismissGroup) //
		groupRouterGroup.POST("/mute_group_member", group.MuteGroupMember)
		groupRouterGroup.POST("/cancel_mute_group_member", group.CancelMuteGroupMember) //MuteGroup
		groupRouterGroup.POST("/mute_group", group.MuteGroup)
		groupRouterGroup.POST("/cancel_mute_group", group.CancelMuteGroup)
		groupRouterGroup.POST("/set_group_member_nickname", group.SetGroupMemberNickname)
		groupRouterGroup.POST("/set_group_member_info", group.SetGroupMemberInfo)
		groupRouterGroup.POST("/get_group_abstract_info", group.GetGroupAbstractInfo)
		//groupRouterGroup.POST("/get_group_all_member_list_by_split", group.GetGroupAllMemberListBySplit)
	}
	superGroupRouterGroup := r.Group("/super_group")
	{
		superGroupRouterGroup.POST("/get_joined_group_list", group.GetJoinedSuperGroupList)
		superGroupRouterGroup.POST("/get_groups_info", group.GetSuperGroupsInfo)
	}
	//certificate
	authRouterGroup := r.Group("/auth")
	{
		authRouterGroup.POST("/user_register", apiAuth.UserRegister) //1
		authRouterGroup.POST("/user_token", apiAuth.UserToken)       //1
		authRouterGroup.POST("/parse_token", apiAuth.ParseToken)     //1
		authRouterGroup.POST("/force_logout", apiAuth.ForceLogout)   //1
	}
	//Third service
	thirdGroup := r.Group("/third")
	{
		thirdGroup.POST("/tencent_cloud_storage_credential", apiThird.TencentCloudStorageCredential)
		thirdGroup.POST("/ali_oss_credential", apiThird.AliOSSCredential)
		thirdGroup.POST("/minio_storage_credential", apiThird.MinioStorageCredential)
		thirdGroup.POST("/minio_upload", apiThird.MinioUploadFile)
		thirdGroup.POST("/upload_update_app", apiThird.UploadUpdateApp)
		thirdGroup.POST("/get_download_url", apiThird.GetDownloadURL)
		thirdGroup.POST("/get_rtc_invitation_info", apiThird.GetRTCInvitationInfo)
		thirdGroup.POST("/get_rtc_invitation_start_app", apiThird.GetRTCInvitationInfoStartApp)
		thirdGroup.POST("/fcm_update_token", apiThird.FcmUpdateToken)
		thirdGroup.POST("/aws_storage_credential", apiThird.AwsStorageCredential)
		thirdGroup.POST("/set_app_badge", apiThird.SetAppBadge)
	}
	//Message
	chatGroup := r.Group("/msg")
	{
		chatGroup.POST("/newest_seq", apiChat.GetSeq)
		chatGroup.POST("/send_msg", apiChat.SendMsg)
		chatGroup.POST("/pull_msg_by_seq", apiChat.PullMsgBySeqList)
		chatGroup.POST("/del_msg", apiChat.DelMsg)
		chatGroup.POST("/del_super_group_msg", apiChat.DelSuperGroupMsg)
		chatGroup.POST("/clear_msg", apiChat.ClearMsg)
		chatGroup.POST("/manage_send_msg", manage.ManagementSendMsg)
		chatGroup.POST("/batch_send_msg", manage.ManagementBatchSendMsg)
		chatGroup.POST("/check_msg_is_send_success", manage.CheckMsgIsSendSuccess)
		chatGroup.POST("/set_msg_min_seq", apiChat.SetMsgMinSeq)

		chatGroup.POST("/set_message_reaction_extensions", apiChat.SetMessageReactionExtensions)
		chatGroup.POST("/get_message_list_reaction_extensions", apiChat.GetMessageListReactionExtensions)
		chatGroup.POST("/add_message_reaction_extensions", apiChat.AddMessageReactionExtensions)
		chatGroup.POST("/delete_message_reaction_extensions", apiChat.DeleteMessageReactionExtensions)
	}
	//Conversation
	conversationGroup := r.Group("/conversation")
	{ //1
		conversationGroup.POST("/get_all_conversations", conversation.GetAllConversations)
		conversationGroup.POST("/get_conversation", conversation.GetConversation)
		conversationGroup.POST("/get_conversations", conversation.GetConversations)
		//deprecated
		conversationGroup.POST("/set_conversation", conversation.SetConversation)
		conversationGroup.POST("/batch_set_conversation", conversation.BatchSetConversations)
		//deprecated
		conversationGroup.POST("/set_recv_msg_opt", conversation.SetRecvMsgOpt)
		conversationGroup.POST("/modify_conversation_field", conversation.ModifyConversationField)
	}
	// office
	officeGroup := r.Group("/office")
	{
		officeGroup.POST("/get_user_tags", office.GetUserTags)
		officeGroup.POST("/get_user_tag_by_id", office.GetUserTagByID)
		officeGroup.POST("/create_tag", office.CreateTag)
		officeGroup.POST("/delete_tag", office.DeleteTag)
		officeGroup.POST("/set_tag", office.SetTag)
		officeGroup.POST("/send_msg_to_tag", office.SendMsg2Tag)
		officeGroup.POST("/get_send_tag_log", office.GetTagSendLogs)

		officeGroup.POST("/create_one_work_moment", office.CreateOneWorkMoment)
		officeGroup.POST("/delete_one_work_moment", office.DeleteOneWorkMoment)
		officeGroup.POST("/like_one_work_moment", office.LikeOneWorkMoment)
		officeGroup.POST("/comment_one_work_moment", office.CommentOneWorkMoment)
		officeGroup.POST("/get_work_moment_by_id", office.GetWorkMomentByID)
		officeGroup.POST("/get_user_work_moments", office.GetUserWorkMoments)
		officeGroup.POST("/get_user_friend_work_moments", office.GetUserFriendWorkMoments)
		officeGroup.POST("/set_user_work_moments_level", office.SetUserWorkMomentsLevel)
		officeGroup.POST("/delete_comment", office.DeleteComment)
	}

	// CloudWallet
	cloudWalletGroup := r.Group("/cloudWalletGroup")
	{
		//cloudWalletGroup.Use(middleware.JWTAuth())

		// 用户账户管理
		cloudWalletGroup.POST("/account", account.Account)                                //获取账户信息
		cloudWalletGroup.POST("/id_card/real_name/auth", account.IdCardRealNameAuth)      //身份证实名认证
		cloudWalletGroup.POST("/set_payment_secret", account.SetPaymentSecret)            // 设置支付密码
		cloudWalletGroup.POST("/user_account_balance", account.UserAccountBalance)        // 查询用户余额
		cloudWalletGroup.POST("/cloud_wallet/record_list", account.CloudWalletRecordList) // 云钱包明细：云钱包收支情况

		//用户银行卡管理
		cloudWalletGroup.POST("/get_user_bankcard_list", account.GetUserBankCardList)         //获取用户银行卡列表
		cloudWalletGroup.POST("/bind_user_bankcard", account.BindUserBankCard)                //绑定银行卡(预提交)
		cloudWalletGroup.POST("/bind_user_bankcard/confirm", account.BindUserBankcardConfirm) //确认绑定银行卡-code验证
		cloudWalletGroup.POST("/Unbinding_user_bankcard", account.BindUserBankcardConfirm)    //解绑银行卡

		/*// 账户充值提现
		cloudWalletGroup.POST("/charge_account", account.ChargeAccount)
		cloudWalletGroup.POST("/draw_account", cloud_wallet.DrawAccount)

		// 红包管理
		cloudWalletGroup.POST("/send_red_packet", cloud_wallet.SendRedPacket)
		cloudWalletGroup.POST("/click_red_packet", cloud_wallet.ClickRedPacket)

		//通过红包id查红包状态
		cloudWalletGroup.POST("/get_red_packet_info", cloud_wallet.GetRedPacketInfo)
		// 红包领取明细
		cloudWalletGroup.POST("/red_packet_click_detail", cloud_wallet.RedPacketClickDetail)
		// 根据日期-》 查询用户的红包记录 ： userid- red.list
		cloudWalletGroup.POST("/list_red_packet_record", cloud_wallet.ListRedPacketRecord)
		// 红包支付确认 ： 当需要选择银行卡支付的时候存在短信验证码
		cloudWalletGroup.POST("/confirm_send_red_packet_code", cloud_wallet.ConfirmSendRedPacketCode)

		// 回调 ： 充值是异步的，提现结果也是异步
		cloudWalletGroup.POST("/send_red_packet_notify", cloud_wallet.SendRedPacketNotify)
		cloudWalletGroup.POST("/draw_notify", cloud_wallet.DrawNotify)

		// ====================== 规划：软删除========================
		// 删除红包记录 : UserID , data range ,RedIds
		cloudWalletGroup.POST("/del_red_packet_record", cloud_wallet.DelRedPacketRecord)
		// 删除领钱明细: UserID , data range ,RedIds
		cloudWalletGroup.POST("/del_account_change_record", cloud_wallet.DelAccountChangeRecord)*/

		// ===================== 脚本 ======================
		// 红包24小时未领取，通知
		// 红包24小时未领取退回
	}

	organizationGroup := r.Group("/organization")
	{
		organizationGroup.POST("/create_department", organization.CreateDepartment)
		organizationGroup.POST("/update_department", organization.UpdateDepartment)
		organizationGroup.POST("/get_sub_department", organization.GetSubDepartment)
		organizationGroup.POST("/delete_department", organization.DeleteDepartment)
		organizationGroup.POST("/get_all_department", organization.GetAllDepartment)

		organizationGroup.POST("/create_organization_user", organization.CreateOrganizationUser)
		organizationGroup.POST("/update_organization_user", organization.UpdateOrganizationUser)
		organizationGroup.POST("/delete_organization_user", organization.DeleteOrganizationUser)

		organizationGroup.POST("/create_department_member", organization.CreateDepartmentMember)
		organizationGroup.POST("/get_user_in_department", organization.GetUserInDepartment)
		organizationGroup.POST("/update_user_in_department", organization.UpdateUserInDepartment)

		organizationGroup.POST("/get_department_member", organization.GetDepartmentMember)
		organizationGroup.POST("/delete_user_in_department", organization.DeleteUserInDepartment)
		organizationGroup.POST("/get_user_in_organization", organization.GetUserInOrganization)
	}

	initGroup := r.Group("/init")
	{
		initGroup.POST("/set_client_config", clientInit.SetClientInitConfig)
		initGroup.POST("/get_client_config", clientInit.GetClientInitConfig)
	}

	return r
}
