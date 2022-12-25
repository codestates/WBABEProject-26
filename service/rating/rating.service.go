package rating_service

import (
	"fmt"
	"time"
	"wemade_project/dto"
	rating_model "wemade_project/model/rating"
)

/////////////////////////
//	  Struct
/////////////////////////

type RatingService struct {
	ratingCollection rating_model.RatingCollection
}

/////////////////////////
//	  Init 
/////////////////////////


func InitWithSelf(model rating_model.RatingCollection) RatingService {
	return RatingService{ratingCollection: model }
}

/////////////////////////
//	  Find Data
/////////////////////////

//메뉴 id에 해당하는 리뷰 평점 데이터를 가져오는 함수
func (s *RatingService) FindList4MenuId(menuId string) ([]*dto.FullReadRatingResponse, error) {
	result, err := s.ratingCollection.FindListByMenuId(menuId)
	if err != nil {
		return nil, err
	}

	fmt.Println("result = ", result)

	var itemList []*dto.FullReadRatingResponse

	//빈 배열 반환
	if (len(result) == 0) {
		itemList = make([]*dto.FullReadRatingResponse, 0)
		return itemList, nil
	}

	for _, item := range result {
		itemList = append(itemList, changeEntity2NormalReadDto(item))
	}
	
	return itemList, nil
}


/////////////////////////
//	  Add Data
/////////////////////////

//리뷰 평점 데이터 등록
func (rs *RatingService) AddRatingItem(addDto dto.CreateRatingRequest) (*dto.FullReadRatingResponse, error) {
	//Model 조립
	userEntity := changeCreateRatingDto2Entity(addDto)
	userEntity.CreateDate = time.Now()
	userEntity.UpdateDate = userEntity.CreateDate

	//데이터 등록 처리
	result, err := rs.ratingCollection.AddEntity(userEntity);
	if err != nil {
		return nil, err
	}

	//등록된 아이템을 반환하기 위해 조회
	saveItem, err1 := rs.ratingCollection.FindByObjectId(result.InsertedID); 
	if err1 != nil {
		return nil, err1
	}

	return changeEntity2NormalReadDto(saveItem), nil
}




/////////////////////////
//	  Utils
/////////////////////////

//Entity 2 Read Dto
func changeEntity2NormalReadDto(entity *rating_model.RatingEntity) *dto.FullReadRatingResponse {
	return &dto.FullReadRatingResponse{Id : entity.ID.String(), UserId: entity.UserId, OderListId: entity.OderListId, MenuId: entity.MenuId, 
	Rating: entity.Rating, ReviewMsg: entity.ReviewMsg, Recommendation: entity.Recommendation, CreateDate: entity.CreateDate, UpdateDate: entity.UpdateDate,}
}

//UserRequest 데이터 중 일부를 Entity랑 매핑하는 함수
func changeCreateRatingDto2Entity(sendDto dto.CreateRatingRequest) rating_model.RatingEntity {
	var entity rating_model.RatingEntity
	
	entity.UserId = sendDto.UserId
	entity.OderListId = sendDto.OrderListId
	entity.MenuId = sendDto.MenuId
	entity.Rating = sendDto.Rating
	entity.ReviewMsg = sendDto.ReviewMsg
	entity.Recommendation = sendDto.Recommendation
	return entity
}


