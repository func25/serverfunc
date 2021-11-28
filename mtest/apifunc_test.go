package apifunc_test

import (
	"fmt"
	"testing"

	"github.com/func25/serverfunc/apifunc"
)

func TestStringToInt64s(t *testing.T) {
	x := "1,2,3,4"
	arr, err := apifunc.StringToInt64s(x, ",")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(arr)

	x = "1,2,3,,,4,"
	arr, err = apifunc.StringToInt64s(x, ",")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(arr)

	x = "1,2,,"
	arr, err = apifunc.StringToInt64s(x, ",")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(arr)

	x = "1,a,asd,3,"
	arr, err = apifunc.StringToInt64s(x, ",")
	if err == nil {
		t.Error("error of course")
		return
	}
	fmt.Println(arr)
}
