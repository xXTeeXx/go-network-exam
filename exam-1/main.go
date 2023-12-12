package main

import "fmt"

func main() {
	// รับข้อมูล Username และ Password จากคีย์บอร์ด
	var username string
	var password string

	fmt.Print("ชื่อผู้ใช้: ")
	fmt.Scanln(&username)
	fmt.Print("รหัสผ่าน: ")
	fmt.Scanln(&password)

	// เรียกใช้ฟังก์ชัน LogIn
	LogIn(username, password)

	// ตัวอย่างการใช้งานฟังก์ชัน Add และ Minus
	resultAdd := Add(5, 3)
	resultMinus := Minus(8, 3)

	// พิมพ์ผลลัพธ์
	fmt.Printf("ผลบวก: %d\n", resultAdd)
	fmt.Printf("ผลลบ: %d\n", resultMinus)
}

func HelloFunction() string {
	return "สวัสดี"
}

func Add(a int, b int) int {
	return a + b
}

func Minus(a int, b int) int {
	return a - b
}

func UserLogin(username string, password string) bool {
	if username == "admin" && password == "password" {
		return true
	} else {
		return false
	}
}

// ใช้งานฟังก์ชัน UserLogin
func LogIn(username string, password string) {
	if UserLogin(username, password) {
		fmt.Println("เข้าสู่ระบบสำเร็จ")
	} else {
		fmt.Println("เข้าสู่ระบบล้มเหลว")
	}
}
