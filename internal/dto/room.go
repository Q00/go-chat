package dto

type Room struct {
	ID        string `json:"id"`
	GroupId   int32  `json:"groupId"`
	LectureId int32  `json:"lectureId"`
	UserId    string `json:"userId"`
	Users     []User `json:"users"`
}
