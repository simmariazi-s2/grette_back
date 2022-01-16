package repositories

import (
	"errors"
	"work/grette_back/database"
	"work/grette_back/database/entities"

	"gorm.io/gorm"
)

/*
	게시글 리스트 가져오기
	Board : 게시글 리스트
*/
func GetBoardList() ([]entities.Board, error) {

	var boardList []entities.Board

	result := database.Db.Model(&boardList).Scan(boardList)

	return boardList, result.Error

}

/*
	댓글 리스트 가져오기
	Reply : 댓글 리스트
*/
func GetReplyList() ([]entities.Reply, error) {
	var replyList []entities.Reply

	result := database.Db.Model(&replyList).Scan(&replyList)

	return replyList, result.Error
}

/*
	좋아요 한 리스트 가져오기
	likeType 0 : 싫어요 , 1 : 좋아요
	Recommand : 좋아요 리스트
*/
func GetLikeRecommandList(likeType int) ([]entities.Recommand, error) {
	var recommandList []entities.Recommand

	result := database.Db.Model(&recommandList).Where("likeType", likeType).Scan(&recommandList)

	return recommandList, result.Error
}

func CreateBoard(board entities.Board) (int, error) {

	if &board == nil {
		return 0, errors.New("게시글 작성 오류 :: Board 정보가 비어있습니다.")
	}

	result := database.Db.Create(&board)

	return int(result.RowsAffected), result.Error

}

func UpdateBoard(board entities.Board) (int, error) {

	if &board == nil {
		return 0, errors.New("게시글 수정 오류 :: Board 정보가 비어있습니다.")
	}

	result := database.Db.Updates(&board)

	return int(result.RowsAffected), result.Error
}

func CreateReply(reply entities.Reply) (int, error) {

	if &reply == nil {
		return 0, errors.New("댓글 작성 오류 :: Reply 정보가 비어있습니다.")
	}

	result := database.Db.Create(&reply)

	return int(result.RowsAffected), result.Error
}

func UpdateReply(reply entities.Reply) (int, error) {

	if &reply == nil {
		return 0, errors.New("댓글 수정 오류 :: Reply 정보가 비어있습니다.")
	}

	result := database.Db.Updates(&reply)

	return int(result.RowsAffected), result.Error
}

/*
	좋아요, 싫어요 등록
	likeType 0 : 싫어요 , 1 : 좋아요
	Recommand : 좋아요 리스트
*/
func SetRecommand(rec *entities.Recommand) (int, error) {
	var result *gorm.DB
	var reRec entities.Recommand
	//reRec = database.Db.Model(&rec).Where("recNo=? AND likeType=?", rec.RecNo, rec.LikeType).Find(&entities.Recommand{})
	result = database.Db.Model(&rec).Where("userNo=? AND boardNo=? AND replyNo=?", rec.RecNo, rec.UserNo, rec.BoardNo, rec.ReplyNo).Scan(&reRec)

	if rec.LikeType == reRec.LikeType {
		result = database.Db.Delete(&rec)
	} else {
		result = database.Db.Save(&rec)
	}

	return int(result.RowsAffected), result.Error
}
