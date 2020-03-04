package basictype

import (
	"fmt"
)

// var val interface{}，空接口，实现接口的方法，相当于接口的实现类，空接口没有方法，相当于所有数据类型的都实现了该接口，相当于object

// Error : error接口
type Error interface {
	error() string
}

// RPCError 结构体
type RPCError struct {
	Code    int64
	Message string
}

// 实现了该接口，接收器
func (e *RPCError) error() string {
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

//上面的代码编译完成，就说明RPCError实现了Error接口，go语言中，接口的继承是隐式继承的
//在 Java 中：实现接口需要显式的声明接口并实现所有方法；
