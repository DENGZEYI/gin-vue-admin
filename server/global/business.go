package global

const Processing uint = 1
const Pass uint = 2
const Fail uint = 3

type ApplyState struct {
	Name string `json:"name" form:"name" gorm:"column:name;comment:;"`
	ID   uint
}
