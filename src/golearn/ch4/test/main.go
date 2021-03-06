package main

import "fmt"

//
// var c = [...]int{1, 2, 3, 4}
//
// func main() {
// 	var a [3]int
// 	fmt.Println(a[0])
// 	fmt.Println(a[len(a)-1])
//
// 	for i, v := range a {
// 		fmt.Printf("%d %d\n", i, v)
// 	}
//
// 	b := [...]int{1, 2, 3}
//
// 	for i, v := range b {
// 		fmt.Printf("%d %d\n", i, v)
// 	}
//
// 	r := [...]int{99: -1}
// 	for i, v := range r {
// 		fmt.Printf("%d %d\n", i, v)
// 	}
// 	fmt.Println(r[0:3])
// }
//
// func main() {
// 	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
// 	Q2 := months[4:7]
// 	summer := months[6:9]
// 	fmt.Println(Q2)
// 	fmt.Println(summer)
//
// 	var runes []rune
// 	for _, r := range "Hello, 世界" {
// 		runes = append(runes, r)
// 	}
// 	fmt.Printf("%q\n", runes)
// }

// func main() {
// 	ages := make(map[string]int)
// 	fmt.Printf("%q\n", ages)
// 	ages = map[string]int{
// 		"alice":   11,
// 		"charlie": 12,
// 	}
// 	for key, value := range ages {
// 		fmt.Printf("%s:%d\n", key, value)
// 	}
// }

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func main() {
	map1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	map2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	map3 := map[string]int{
		"a": 1,
		"b": 0,
	}
	map4 := map[string]int{
		"a": 1,
		"c": 0,
	}
	fmt.Println(equal(map1, map2))
	fmt.Println(equal(map2, map3))
	fmt.Println(equal(map3, map4))
}
