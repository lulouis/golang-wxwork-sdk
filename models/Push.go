package models

/******************************************************************************************************************
touser、toparty、totag不能同时为空
text类型,使用text属性
textcard,使用textcard属性
file类型，使用file属性

参数	是否必须	说明
touser	否	成员ID列表（消息接收者，多个接收者用‘|’分隔，最多支持1000个）。特殊情况：指定为@all，则向该企业应用的全部成员发送
toparty	否	部门ID列表，多个接收者用‘|’分隔，最多支持100个。当touser为@all时忽略本参数
totag	否	标签ID列表，多个接收者用‘|’分隔，最多支持100个。当touser为@all时忽略本参数
msgtype	是	消息类型，此时固定为：text
agentid	是	企业应用的id，整型。企业内部开发，可在应用的设置页面查看；第三方服务商，可通过接口 获取企业授权信息 获取该参数值
content	是	消息内容，最长不超过2048个字节，超过将截断
safe	否	表示是否是保密消息，0表示否，1表示是，默认0

media_id	是	文件id，可以调用上传临时素材接口获取
******************************************************************************************************************/

type File struct {
	Media_id string `json:"media_id"`
}

type Text struct {
	Content string `json:"content"`
}

type Textcard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type Message struct {
	Touser   string   `json:"touser"`
	Toparty  string   `json:"toparty"`
	Totag    string   `json:"totag"`
	Msgtype  string   `json:"msgtype"`
	Agentid  int      `json:"agentid"`
	Safe     int      `json:"safe"`
	Text     Text     `json:"text"`     //Msgtype为text时方可使用
	File     File     `json:"file"`     //Msgtype为file时方可使用
	Textcard Textcard `json:"textcard"` //Msgtype为textcard时方可使用

}

//文件上传时的返回
type UploadMediaRes struct {
	PublicResult
	Type     string `json:"type"`
	Media_id string `json:"media_id"`
}
