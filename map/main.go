package main

import (
	"fmt"
)

func checkMapNum() {
	// m := make(map[int]string)
	// m[1] = "US"
	// m[2] = "JAPAN"
	// m[3] = "CHINA"

	m := map[int]string{1: "US", 2: "JAPAN", 3: "CIHNA"}

	fmt.Println(m)
}

func checkMapStr() {
	m := make(map[string]string)
	m["TANAKA"] = "TARO"
	m["SUZUKI"] = "ICHIRO"
	m["TANAKA"] = "YUKO"

	fmt.Println(m)
}

func checkMapArray() {

	// m := map[int][]int{
	// 	1: []int{1},
	// 	2: []int{1, 2, 3},
	// 	3: []int{1, 2, 3, 4, 5},
	// }

	m := map[int][]int{
		1: {1},
		2: {1, 2, 3},
		3: {1, 2, 3, 4, 5},
	}

	fmt.Println(m)
}

func checkMapInMap() {

	m := map[int]map[string]string{
		1: {"JP": "円"},
		2: {"US": "ドル"},
		3: {"CN": "元"},
	}

	fmt.Println(m)
}

func checkMapBoolean() {

	m := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}

	s1, ok1 := m[1]
	s4, ok4 := m[4]

	fmt.Println(m)
	fmt.Println(s1)
	fmt.Println(ok1)
	fmt.Println(s4)
	fmt.Println(ok4)

	fmt.Println(len(m))
	delete(m, 2)
	fmt.Println(len(m))
}

func checkForMap() {

	m := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}

	for i, v := range m {
		fmt.Println(i, v)
	}
}

func main() {
	// checkMapNum()
	// checkMapStr()
	// checkMapArray()
	// checkMapInMap()
	// checkMapBoolean()
	checkForMap()
}
