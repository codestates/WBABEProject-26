package dto

/////////////////////////
//		Resposne
/////////////////////////


type NomalReadUserResponse struct {
	Name string `bson:"name"` //사용자 이름 
	Phone string `bson:"phone"` //사용자 폰번호
	Addr string `bson:"addr"` //주소
}