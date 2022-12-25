package rating_route

import (
	rating_controller "wemade_project/controller/rating"

	"github.com/gin-gonic/gin"
)


type RatingRoute struct {
	ratingController rating_controller.RatingController
}

func InitWithSelf(ratingController rating_controller.RatingController ) RatingRoute {
	return RatingRoute{ratingController: ratingController}
}

func (r *RatingRoute) InitWithRoute(server *gin.Engine) {
	ratingRouterV1 := server.Group("/api/v1/rating")
	{
		ratingRouterV1.POST("/add",r.ratingController.AddRating()) //메뉴 리뷰 등록
		// orderListRouterV1.PUT("/menu/update", r.menuController.UpdateMenu())
		// orderListRouterV1.DELETE("/menu/delete/:menu_id", r.menuController.DeleteMenu4Logical())
	}
}