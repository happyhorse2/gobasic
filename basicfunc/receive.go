package basicfunc

//func (接收器变量 接收器类型) 方法名(参数列表) (返回参数){
//函数体
//}

// Bag : bag结构体
type Bag struct {
	items []int
}

//Insert 面向过程
func Insert(b *Bag, itemid int) {
	b.items = append(b.items, itemid)
}

//Insertmethod 类面向对象
func (b *Bag) Insertmethod(itemid int) {

}
