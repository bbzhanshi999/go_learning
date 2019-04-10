package main

import (
	"fmt"
	"sort"
)

func main() {
	drinkMap := map[string]string{"CocoCola": "可口可乐", "BesiCola": "百事可乐", "Sprite": "雪碧", "Mirinda": "美年达", "Hedy": "七喜", "Soda": "苏打水"}

	for k, v := range drinkMap {
		fmt.Printf("%v,", drinkMap[k])
		fmt.Printf("drinking in En:%s,in Chinese:%s\n", k, v)
	}

	sorArr := make([]string, len(drinkMap))

	i := 0

	for k, _ := range drinkMap {
		sorArr[i] = k
		i++
	}

	sort.Strings(sorArr)

	for _, v := range sorArr {
		fmt.Printf("%v /", v)
	}

}
