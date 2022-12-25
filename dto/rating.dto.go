package dto

import (
	"time"
	rating_enum "wemade_project/enums/rating"
)

/////////////////////////
//	 Create Request
/////////////////////////

//평점 별점 리뷰 생성 요청
type CreateRatingRequest struct {
	UserId string `json:"userId" binding:"required"`
	OrderListId string `json:"orderListId" binding:"required"`
	MenuId string `json:"menuId" binding:"required"`
	Rating rating_enum.RatingScore `json:"rating" binding:"required"` 
	ReviewMsg string `json:"reviewMsg" binding:"required"`
	Recommendation bool `json:"recommendation"` //binding:"required"
}

/////////////////////////
//		Resposne
/////////////////////////

//리뷰평가 읽기 데이터
type FullReadRatingResponse struct {
	Id string `json:"id"`
	UserId string `json:"userId"` //주문자 (_id)
	OderListId string `json:"orderListId"` //주문리스트 (_id)
	MenuId string `json:"menuId"` //메뉴 (_id)
	Rating rating_enum.RatingScore `json:"rating"`
	ReviewMsg string `json:"reviewMsg"`
	Recommendation bool `json:"recommendation"`
	CreateDate time.Time `json:"createDate"` //데이터 생성 시각
	UpdateDate time.Time `json:"updateDate"` //데이터 수정 시각
}

