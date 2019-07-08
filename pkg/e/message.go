package e

var message = map[int]string{
	SUCCESS:        "请求成功",
	ERROR:          "请求失败",
	INVALID_PARAMS: "参数错误",

	ERROR_EXIST_TAG: "标签已存在",
	ERROR_NOT_EXIST_TAG: "标签不存在",

	ERROR_AUTH: "用户验证失败",
	ERROR_AUTH_TOKEN: "用户TOKEN生成失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL: "TOKEN解析错误",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "TOKEN超时",
}

func GetMsg(num int) string {
	if msg, ok := message[num]; ok {
		return msg
	}
	return ""
}
