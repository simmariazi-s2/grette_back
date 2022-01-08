package message

var MessageFlags = map[int]string{
	SUCCESS:        "OK",
	ERROR:          "Error",
	INVALID_PARAMS: "Invalid Params",
}

func GetMessage(code int) string {
	msg, ok := MessageFlags[code]
	if ok {
		return msg
	}

	return MessageFlags[ERROR]
}
