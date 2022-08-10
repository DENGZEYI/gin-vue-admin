package global

const Processing uint = 1
const Pass uint = 2
const Ingress uint = 3 //以入库
const Egress uint = 4  //已出库
const Fail uint = 5

type ApplyState struct {
	Name string `json:"name" form:"name" gorm:"column:name;comment:;"`
	ID   uint
}
