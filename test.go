package main

import "fmt"

func main()  {
	mainMapB := make(map[string]map[string]string)
	//内部容器必须再次初始化才能使用
	subMapB := make(map[string]string)
	subMapB["B_Key_1"] = "B_SubValue_1"
	subMapB["B_Key_2"] = "B_SubValue_2"
	mainMapB["MapB"] = subMapB
	fmt.Println("\nMultityMapB")

	fmt.Println(mainMapB)

	for keyB, valB := range mainMapB {
		for subKeyB, subValB := range valB {
			fmt.Printf("mapName=%s	Key=%s	Value=%s\n", keyB, subKeyB, subValB)
		}
	}
}


