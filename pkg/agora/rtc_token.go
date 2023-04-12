package agora

import (
	"Open_IM/pkg/common/log"
	"fmt"
	rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtctokenbuilder2"
)

const (
	tokenExpireTime     = 40
	privelegeExpireTime = 40
)

// 使用 RtcTokenBuilder 来生成 RTC Token
func GenerateRtcToken(int_uid, OperationID string, channelName string, role rtctokenbuilder.Role) (string, error) {

	appID := "7ab956ddab30495e85dab000ce22f77d"
	appCertificate := "49642a91127249619a730e9ac4db75ff"
	// AccessToken2 过期的时间，单位为秒
	// 当 AccessToken2 过期但权限未过期时，用户仍在频道里并且可以发流，不会触发 SDK 回调。
	// 但一旦用户和频道断开连接，用户将无法使用该 Token 加入同一频道。请确保 AccessToken2 的过期时间晚于权限过期时间。
	tokenExpireTimeInSeconds := uint32(tokenExpireTime)
	// 权限过期的时间，单位为秒。
	// 权限过期30秒前会触发 token-privilege-will-expire 回调。
	// 权限过期时会触发 token-privilege-did-expire 回调。
	// 为作演示，在此将过期时间设为 40 秒。你可以看到客户端自动更新 Token 的过程
	privilegeExpireTimeInSeconds := uint32(privelegeExpireTime)
	result, err := rtctokenbuilder.BuildTokenWithUserAccount(appID, appCertificate, channelName, int_uid, role, tokenExpireTimeInSeconds, privilegeExpireTimeInSeconds)
	if err != nil {
		log.Error(OperationID, fmt.Sprintf("build token with user account error: %v", err))
		return "", err
	}
	return result, err
}
