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
	orderListRouterV1 := server.Group("/api/v1/rating")
	{
		orderListRouterV1.GET("menu/get",)
		// orderListRouterV1.POST("/menu/add",r.menuController.AddMenu())
		// orderListRouterV1.PUT("/menu/update", r.menuController.UpdateMenu())
		// orderListRouterV1.DELETE("/menu/delete/:menu_id", r.menuController.DeleteMenu4Logical())
	}
}