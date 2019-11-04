package ch16_time

import (
	"crypto/sha256"
	"fmt"
	"testing"
	"time"
)

//在Go语言中，“Y-m-d H:i:s”、 “yyyy-MM-dd HH:mm:ss”为特定的数字 “2006-01-02 15:04:05”是Go语言的创建时间，且必须为这几个准确的数字
const DATETIME_LAYOUT = "2006-01-02 15:04:05"

func TestInt642FormatStringTime(t *testing.T) {
	timeStamp := int64(1572837980) //time.Now().Unix()  //当前时间以秒计算
	//fmt.Println(timeStamp)
	formatStrDte := time.Unix(timeStamp, 0).Format(DATETIME_LAYOUT)
	fmt.Println(formatStrDte)
}

func TestSha(t *testing.T) {
	str := "ansatsing"
	hsh := sha256.Sum256([]byte(str))
	fmt.Printf("%x\n,%s\n,%d\n", hsh, hsh, hsh)
}
