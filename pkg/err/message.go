package err

var message = map[int]string{
	SUCCESS:        "请求成功",
	ERROR:          "请求失败",
	INVALID_PARAMS: "参数错误",
	ERROR_EXIST_TAG: "标签已存在",
	ERROR_NOT_EXIST_TAG: "标签不存在",
}

func GetMsg(num int) string {
	if msg, ok := message[num]; ok {
		return msg
	}
	return ""
}
