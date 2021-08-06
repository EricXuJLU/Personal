package define

import (
	"fmt"
)

const (
	Port = ":10086"
	Address = "localhost"+Port
)

func Check(err error, id int){
	if err != nil{
		fmt.Println(err, id)
	}
}