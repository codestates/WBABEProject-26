package user_service

import (
	"wemade_project/dto"
	user_model "wemade_project/model/user"
)

/////////////////////////
//	  Struct
/////////////////////////

type UserService struct {
	userCollection user_model.UserCollection
}

/////////////////////////
//	  Init 
/////////////////////////


func InitWithSelf(model user_model.UserCollection) UserService {
	return UserService{userCollection: model }
}


/////////////////////////
//	  Find Data
/////////////////////////


//메뉴 id로 엔티티 데이터를 조회하는 함수
func (s *UserService) Find4UserPhone(phone string) (*dto.NomalReadUserResponse, error) {
	findUser, findErr := s.userCollection.FindByPhone(phone)
	if findErr != nil {
		return nil, findErr
	}
	
	return changeUserEntity2UserReadDto(*findUser), nil
}


/////////////////////////
//	  Utils
/////////////////////////

func changeUserEntity2UserReadDto(entity user_model.UserEntity) *dto.NomalReadUserResponse {
	return &dto.NomalReadUserResponse{Name:  entity.Name, Phone:  entity.Phone, Addr:  entity.Addr}
}

