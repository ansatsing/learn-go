package ch8_string

import (
	"strconv"
	"strings"
	"testing"
)

func TestSplitAndJoin(t *testing.T) {
	s := "a,b,c"
	parts := strings.Split(s, ",")
	for idx, vle := range parts {
		t.Log(idx, vle)
	}

	s = strings.Join(parts, "-")
	t.Log(s)
}

func TestStringIntCovert(t *testing.T) {
	s := "10"
	//t.Log(strconv.Atoi(s)+10) //multiple-value strconv.Atoi() in single-value context
	if vle, err := strconv.Atoi(s); err != nil {
		t.Log("err")
	} else {
		t.Log(vle + 10)
	}

	s = strconv.Itoa(10)
	t.Logf("type of s is:%T", s)
}
