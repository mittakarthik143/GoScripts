// this function we define to take one or more parameters

package main

import "fmt"

func take_multi_params(data ...string) {
	fmt.Println("Data is: ", data, " And length of the data is: ", len(data))
	//fmt.Println(data[0])
}

func take_multi_int(data ...int) {
	if len(data) == 0 || len(data) == 1 {
		fmt.Println("Integer values is: ", data, " and length is: ", len(data))
	} else {
		fmt.Println("Integer values are: ", data, " and length is: ", len(data))
	}
}

func main() {
	take_multi_params("10", "20")
	take_multi_params()
	take_multi_params("sai", "ram", "karthikeyan", "balaji", "hari", "vasavi")
	take_multi_int(10, 20, 30)
	take_multi_int(10)
}
