package err

var message = map[int]string{
	SUCCESS:        "请求成功",
	ERROR:          "请求失败",
	INVALID_PARAMS: "参数错误",
}

func GetMsg(num int) string {
	if msg, ok := message[num]; ok {
		return msg
	}
	return ""
}
