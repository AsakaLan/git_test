package test

import "testing"

type addTest struct{
	arg1,arg2,expected int 
}
var addTests = []addTest{
	addTest{2,3,5},
	addTest{5,2,7},
	addTest{4,8,12},
	addTest{3,10,13},
}
func TestAdd(t *testing.T){
	for _,test := range addTests{
		if outputs:=Add(test.arg1,test.arg2); outputs != test.expected{
			t.Error("TestAdd failed. arg1:",test.arg1,"arg2:",test.arg2,"expected:",test.expected,"but got:",outputs)
		}
	}
}