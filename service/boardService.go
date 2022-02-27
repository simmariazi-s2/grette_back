package service

import (
	"container/list"
	"log"
	"net/http"
	"work/grette_back/app"
	"work/grette_back/database/entities"
	"work/grette_back/message"
	"work/grette_back/model"
	"work/grette_back/repositories"

	"github.com/gin-gonic/gin"
)

// 게시글 목록 조회
func SelectBoardList(c *gin.Context) {
	gin := app.Gin{C: c}

	boardList, err := repositories.GetBoardList()

	if err != nil {
		gin.Response(http.StatusBadRequest, message.ERROR, boardList)
		log.Print(err.Error())
		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, boardList)
	log.Print("게시글 목록 조회 성공")
}

// 게시글 쓰기
func WriteBoard(c *gin.Context) {
	gin := app.Gin{C: c}

	resBoard := new(model.Board)

	if err := gin.C.ShouldBindJSON(&resBoard); err != nil {
		gin.Response(http.StatusBadRequest, message.CREATE_FAIL, resBoard)
		log.Print(err.Error())

		return
	}

	boardEntity := new(entities.Board)
	boardEntity.BoardNo = resBoard.BoardNo
	boardEntity.Title = &resBoard.Title
	boardEntity.Content = &resBoard.Content
	boardEntity.CategoryNo = &resBoard.CategoryNo
	boardEntity.Dislike = resBoard.Dislike
	boardEntity.UserNo = &resBoard.UserNo
	boardEntity.IsUsed = resBoard.IsUsed

	result, err := repositories.CreateBoard(*boardEntity)

	if err != nil {
		gin.Response(http.StatusBadRequest, message.SAVE_ERROR, result)
		log.Print(err.Error())
		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, result)
	return
}

// 게시글 수정
func UpdateBoard(c *gin.Context) {
	gin := app.Gin{C: c}

	resBoard := new(model.Board)

	if err := gin.C.ShouldBindJSON(&resBoard); err != nil {
		gin.Response(http.StatusBadRequest, message.CREATE_FAIL, resBoard)
		log.Print(err.Error())
		return
	}

	updateBoard := new(entities.Board)
	updateBoard.BoardNo = resBoard.BoardNo
	updateBoard.Title = &resBoard.Title
	updateBoard.Content = &resBoard.Content
	updateBoard.CategoryNo = &resBoard.CategoryNo
	updateBoard.Dislike = resBoard.Dislike
	updateBoard.UserNo = &resBoard.UserNo
	updateBoard.IsUsed = resBoard.IsUsed

	result, err := repositories.UpdateBoard(*updateBoard)

	if err != nil {
		gin.Response(http.StatusBadRequest, message.SAVE_ERROR, result)
		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, result)
	return
}

// 게시글 삭제
func DeleteBoard(c *gin.Context) {
	gin := app.Gin{C: c}

	resBoard := new(model.Board)

	if err := gin.C.ShouldBindJSON(&resBoard); err != nil {
		gin.Response(http.StatusBadRequest, message.CREATE_FAIL, resBoard)
		log.Print(err.Error())
		return
	}

	entityBoard := new(entities.Board)
	entityBoard.BoardNo = resBoard.BoardNo
	entityBoard.IsUsed = resBoard.IsUsed

	result, err := repositories.UpdateBoard(*entityBoard)

	if err != nil {
		gin.Response(http.StatusBadRequest, message.SELECT_FAIL, result)
		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, result)
	return
}

// 댓글 목록 조회
func SelectReplyList(c *gin.Context) {
	gin := app.Gin{C: c}

	result, err := repositories.GetReplyList()

	if err != nil {
		gin.Response(http.StatusBadRequest, message.SELECT_FAIL, result)
		log.Print(err.Error())
		return
	}

	retList := list.New()

	// entity model -> service model
	for value := range result {
		retList.PushBack(value)
	}

	if err != nil {
		gin.Response(http.StatusBadRequest, message.SELECT_FAIL, retList)
		log.Print(err.Error())
		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, retList)
	return
}

// 댓글 쓰기
func WriteReply(c *gin.Context) {
	gin := app.Gin{C: c}

	resReply := new(model.Reply)

	if err := gin.C.ShouldBindJSON(&resReply); err != nil {
		gin.Response(http.StatusBadRequest, message.CREATE_FAIL, resReply)
		log.Print(err.Error())
		return
	}

	reply := new(entities.Reply)
	reply.BoardNo = &resReply.BoardNo
	reply.Content = &resReply.Content
	reply.Dislike = resReply.Dislike
	reply.Report = resReply.Report
	reply.UserNo = &resReply.UserNo
	reply.Like = resReply.Like
	reply.Dislike = resReply.Dislike
	reply.IsUsed = resReply.IsUsed

	result, err := repositories.CreateReply(*reply)

	if err != nil {
		gin.Response(http.StatusBadRequest, message.CREATE_FAIL, result)
		log.Print(err.Error())
		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, result)
	return

}

// 댓글 수정
func UpdateReply(c *gin.Context) {
	gin := app.Gin{C: c}

	resReply := new(model.Reply)

	if err := gin.C.ShouldBindJSON(&resReply); err != nil {
		gin.Response(http.StatusBadRequest, message.CREATE_FAIL, resReply)
		log.Print(err.Error())
		return
	}

	entityReply := new(entities.Reply)
	entityReply.BoardNo = &resReply.BoardNo
	entityReply.Content = &resReply.Content
	entityReply.Dislike = resReply.Dislike
	entityReply.Like = resReply.Like
	entityReply.IsUsed = resReply.IsUsed
	entityReply.Report = resReply.Report

	result, err := repositories.UpdateReply(*entityReply)

	if err != nil {
		gin.Response(http.StatusBadRequest, message.UPDATE_FAIL, result)
		log.Print(err.Error())
		return
	}

	gin.Response(http.StatusOK, message.SUCCESS, result)
	return
}

// 댓글 삭제
func DeleteReply(c *gin.Context) {

}

// 좋아요 (게시글 or 댓글)
func SetRecommand(c *gin.Context) {
	// gin := app.Gin{C: c}

	// recommand := new(model.Recommand);

	// if err:= gin.C.ShouldBindJSON()
}
