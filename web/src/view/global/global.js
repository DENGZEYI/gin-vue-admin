//定义变量或常量
const name ='张三'
// 订单的状态
const orderOptions = 
  { "PROCESSING": 1, "PASS": 2,"FAIL":3,"PURCHASING":4 }
//定义方法test
function test () {
  const text="Hello World"
  alert(text)
}
//导出设置使得Vue可引入，关键
export {
	test,
  orderOptions,
	name
}
