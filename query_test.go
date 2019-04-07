package main

import "testing"

func TestQuery(t *testing.T) {
	a :=make([]string,2)

	a=append(a,"测试")

	Query(a,false,false,false,false)
}
