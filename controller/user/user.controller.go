package user_controller

import user_service "wemade_project/service/user"

/////////////////////////
//	  Struct
/////////////////////////

type UserController struct {
	userService user_service.UserService
}


/////////////////////////
//	  Init func
/////////////////////////

//생성자 역할 함수
func InitWithSelf(userService user_service.UserService) UserController {
	return UserController{userService: userService}
}
