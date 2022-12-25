package user_service

import (
	"time"
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
//	  Add Data
/////////////////////////

//사용자 등록
func (us *UserService) AddMenuItem(addDto dto.CreateUserRequest) (*dto.NomalReadUserResponse, error) {
	//Model 조립
	userEntity := changeCreateUserDto2Entity(addDto)
	userEntity.CreateDate = time.Now()
	userEntity.UpdateDate = userEntity.CreateDate

	//데이터 등록 처리
	result, err := us.userCollection.AddEntity(userEntity);
	if err != nil {
		return nil, err
	}

	//등록된 아이템을 반환하기 위해 조회
	saveItem, err1 := us.userCollection.FindByObjectId(result.InsertedID); 
	if err1 != nil {
		return nil, err1
	}

	return changeUserEntity2UserReadDto(*saveItem), nil
}


/////////////////////////
//	  Utils
/////////////////////////

//UserRequest 데이터 중 일부를 Entity랑 매핑하는 함수
func changeCreateUserDto2Entity(sendDto dto.CreateUserRequest) user_model.UserEntity {
	var userEntity user_model.UserEntity
	
	userEntity.Name = sendDto.Name
	userEntity.Phone = sendDto.Phone
	userEntity.Addr = sendDto.Addr
	return userEntity
}



func changeUserEntity2UserReadDto(entity user_model.UserEntity) *dto.NomalReadUserResponse {
	return &dto.NomalReadUserResponse{Name:  entity.Name, Phone:  entity.Phone, Addr:  entity.Addr}
}

