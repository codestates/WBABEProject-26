package dto

/////////////////////////
//	 Create Request
/////////////////////////

/*
required를 통해서 필수 값에 대해서 validation을 check 하도록 구현하신 점 좋습니다.
*/
//메뉴 생성에 사용하는 데이터
type CreateUserRequest struct {
	Name string `json:"name" binding:"required" ` //사용자 이름 
	Phone string `json:"phone" binding:"required"` //사용자 폰번호
	Addr string `json:"addr" binding:"required"` //주소
}


/////////////////////////
//		Resposne
/////////////////////////


type NomalReadUserResponse struct {
	Name string `bson:"name"` //사용자 이름 
	Phone string `bson:"phone"` //사용자 폰번호
	Addr string `bson:"addr"` //주소
}