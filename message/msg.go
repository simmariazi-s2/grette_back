package message

var MessageFlags = map[int]string{
	SUCCESS:        "OK",
	ERROR:          "Error",
	INVALID_PARAMS: "Invalid Params",
	LOGIN_FAIL:     "Login Failed",
	UPDATE_FAIL:    "Update Faild",
	SELECT_FAIL:    "Select Failed",
	SAVE_ERROR:     "Database Save Failed",
}

func GetMessage(code int) string {
	msg, ok := MessageFlags[code]
	if ok {
		return msg
	}

	return MessageFlags[ERROR]
}
