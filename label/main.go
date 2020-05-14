package main

import "fmt"

func main() {
	j := 1

	/**
	   * Label注意事项：
	   * 		WithCancel、必须定义在任何变量之下
	  		2、针对for、switch、select的break，Label位置必须紧挨着break操作的for/switch/select
	*/
FirstLoop:
	switch j {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
		break FirstLoop // invalid break label FirstLoop
	}
}
