package main

import "fmt"

func main() {
	// Variadic function
	// ฟังก์ชั่นที่รับค่ากี่ตัวก็ได้ในฟังก์ชั่นเห็นจะเห็นเป็นslice
	s := []string{"x", "y", "z"}
	variadicString("a", "1", "b", "2", "c")

	//ส่งsliceเข้าไปที่ variadic function
	variadicString(s...)
}

func variadicString(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
