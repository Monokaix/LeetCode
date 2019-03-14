package main

import (
	"fmt"
)

func validUtf8(data []int) bool {
	nBytes := 0
	for i:= 0;i<len(data);i++{
		if nBytes == 0{
			if (data[i] & 0x80) != 0 {  //与操作之后不为0，说明首位为1
				for (data[i] & 0x80) != 0 {
					data[i] <<= 1 //左移一位
					nBytes++ //记录字符共占几个字节
				}

				if nBytes < 2 || nBytes > 6 { //因为UTF8编码单字符最多不超过6个字节
					return false
				}

				nBytes-- //减掉首字节的一个计数
			}
		}else{ //处理多字节字符
			if data[i] & 0xc0 != 0x80{ //判断多字节后面的字节是否是10开头
				return false
			}
			nBytes--
		}
	}
	return nBytes == 0
}
func main(){
	fmt.Println(validUtf8([]int{250,145,145,145,145}))
}
