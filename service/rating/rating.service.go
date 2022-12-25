package rating_service

import rating_model "wemade_project/model/rating"

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