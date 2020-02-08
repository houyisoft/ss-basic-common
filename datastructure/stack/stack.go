package stack

import (
	"errors"
	"fmt"
)

type Stack []interface{}

func (s *Stack)Push(x ...interface{}){
	*s = append(*s, x...)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("slice为空!")
	}
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result, nil
}

func (s *Stack) Len() int {
	return len(*s)
}

//修改指定索引的元素
func (s *Stack)Set(idx int,value interface{})(err error){
	if idx >= 0 && s.Len() > 0 && s.Len() > idx{
		(*s)[idx] = value
		return nil
	}
	return errors.New("Set失败!")
}

//返回指定索引的元素
func (s *Stack)Get(idx int)(value interface{}){
	if idx >= 0 && s.Len() > 0 && s.Len() > idx {
		return (*s)[idx]
	}
	return nil //read empty stack
}

//是否为空
func (s *Stack)Empty()(bool){
	if s == nil || s.Len() == 0 {
		return true
	}
	return false
}

//打印
func (s *Stack)Print(){
	for i := s.Len() - 1; i >= 0; i--{
		fmt.Println(i,"=>",(*s)[i])
	}
}
