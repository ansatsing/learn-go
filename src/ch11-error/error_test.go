package ch11_error

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func getFanbonacci(num int) ([]int, error) {
	if num < 2 {
		return nil, errors.New("num should greater than 2")
	}
	if num > 100 {
		return nil, errors.New("num should less than 100")
	}
	arr := []int{1, 1}
	for i := 0; i < num; i++ {
		arr = append(arr, arr[i]+arr[i+1])
	}
	return arr, nil
}

var LessThanTwoError = errors.New("num should greater than 2")
var LargerThenHundredError = errors.New("num should less than 100")

func getFanbonacci1(num int) ([]int, error) {
	if num < 2 {
		return nil, LessThanTwoError
	}
	if num > 100 {
		return nil, LargerThenHundredError
	}
	arr := []int{1, 1}
	for i := 0; i < num; i++ {
		arr = append(arr, arr[i]+arr[i+1])
	}
	return arr, nil
}

func TestError(t *testing.T) {
	if vle, err := getFanbonacci(4); err == nil {
		t.Log(vle)
	} else {
		t.Error(err)
	}
}

func GetFabnocciByString(s string) {
	var (
		num    int
		result []int
		err    error
	)
	if num, err = strconv.Atoi(s); err != nil {
		fmt.Println("Error:", err)
		return
	}
	if result, err = getFanbonacci(num); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
func TestError1(t *testing.T) {
	s := "10"
	GetFabnocciByString(s)

	s = "1"
	GetFabnocciByString(s)
}
