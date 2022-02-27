package message

const (
	SUCCESS        = 200 // 성공
	LOGIN_FAIL     = 201 // 로그인 실패
	UPDATE_FAIL    = 202 // 수정 실패
	CREATE_FAIL    = 203 // 생성 실패
	SELECT_FAIL    = 204 // 조회 실패
	INVALID_PARAMS = 400 // 파라미터 오류
	ERROR          = 403 // 오류
	SAVE_ERROR     = 404 // DB 저장 오류
)
