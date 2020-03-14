package busModel

type BusModel struct {
	Status int `json:"status" gorm:"type:tinyint;default 1;not null"`
}