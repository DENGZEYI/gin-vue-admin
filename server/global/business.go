package global

// BusOrder的状态
const Processing uint = 1
const Pass uint = 2
const Fail uint = 3
const Purchasing uint = 4

type ApplyState struct {
	Name string `json:"name" form:"name" gorm:"column:name;comment:;"`
	ID   uint
}

// Chair 主任的authorityID
const Chair uint = 1
const Manager uint = 2
