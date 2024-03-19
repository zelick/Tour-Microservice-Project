package model

type TourPointRequest struct {
	ID          int    `json:"id" gorm:"column:Id;primaryKey;autoIncrement"`
	AuthorId    int    `json:"authorId"`
	TourPointId int    `json:"tourPointId"`
	Status      string `json:"status"`
}
