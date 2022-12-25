package dto

/////////////////////////
//	 Create Request
/////////////////////////

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