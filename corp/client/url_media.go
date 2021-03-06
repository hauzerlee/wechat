// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/magicshui/wechat for the canonical source repository
// @license     https://github.com/magicshui/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

// https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE
func _MediaUploadURL(accesstoken, mediaType string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=" + accesstoken +
		"&type=" + mediaType
}

// https://qyapi.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
func _MediaDownloadURL(accesstoken, mediaId string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/media/get?access_token=" + accesstoken +
		"&media_id=" + mediaId
}
