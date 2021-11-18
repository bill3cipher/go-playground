package main

import "fmt"

func main() {

	fmt.Println("============for loop============")
	//for loop
	for i := 0; i <= 10; i++ {
		fmt.Println("i=", i)
	}

	//while loop
	fmt.Println("============while loop============")
	j := 0
	for j <= 10 {
		fmt.Println("j=", j)
		j++
	}

	//infinity loop
	fmt.Println("============infinity loop============")
	k := 0
	for {
		k++
		fmt.Println("k=", k)
		if k == 10 {
			break
		}
	}

	n := sumOfFirst(3)
	println(n)

	boo1 := isPrime(1)
	println(boo1)
	boo2 := isPrime(2)
	println(boo2)
	boo3 := isPrime(3)
	println(boo3)
	boo4 := isPrime(4)
	println(boo4)
}

// func sumOfFirst ใช้ for ในการบวกค่า โดยค่าที่นำมาบวกจะลดลงทีละ 1 จนถึง 0
func sumOfFirst(n int) int {
	sum := 0
	for i := n; i > 0; i-- {
		sum = sum + i
	}
	return sum
}

// func isPrime ใช้ for ในการเช็คว่าไม่สามารถหาค่าระหว่าง 1 (ตั้งแต่ 2 ขึ้นไป) จนถึงตัวมันเองมาหารได้ลงตัว
func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
