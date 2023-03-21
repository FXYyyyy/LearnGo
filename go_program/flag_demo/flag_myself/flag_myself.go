package flag_myself

import (
	"errors"
	"strconv"
)

var errParse = errors.New("parse error")

// MyFlagValue 通用的类型接口
type MyFlagValue interface {
	String() string
	MySet(string) error
}

// MyIntValue 继承通用接口
type MyIntValue int

func (b *MyIntValue) MySet(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		err = errParse
	}
	*b = MyIntValue(v) // 类型转换
	return err
}

func (b *MyIntValue) String() string {
	return strconv.Itoa(int(*b))
}

// MyFlag 某个命令参数
type MyFlag struct {
	name     string
	usage    string
	value    MyFlagValue
	defValue MyFlagValue
}

// MyFlagSet 命令集
type MyFlagSet struct {
	name   string
	args   []string           // 参数
	formal map[string]*MyFlag // 记录已经解析过命令
}
