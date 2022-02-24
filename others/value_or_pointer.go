package others

import (
	"fmt"
)

func ValueOrPointer() {

	fmt.Println("")
	var i int = 20
	fmt.Println("int")
	fmt.Println(_int(i))
	fmt.Println(i)

	fmt.Println("")
	var f float64 = 3.14
	fmt.Println("float")
	fmt.Println(_float(f))
	fmt.Println(f)

	fmt.Println("")
	var b bool = true
	fmt.Println("bool")
	fmt.Println(_bool(b))
	fmt.Println(b)

	fmt.Println("")
	var s string = "start!"
	fmt.Println("string")
	fmt.Println(_string(s))
	fmt.Println(s)

	fmt.Println("")
	ai := []int{1, 2, 3, 4}
	fmt.Println("_array")
	fmt.Println(_array(ai)[0])
	fmt.Println(ai[0])

	fmt.Println("")
	u := User{Name: "Yanun", Age: 26}
	fmt.Println("struct")
	fmt.Println(_struct(u).Name)
	fmt.Println(u.Name)

	fmt.Println("")
	u1 := User{Name: "Yanun", Age: 26}
	u2 := User{Name: "Vin", Age: 21}
	us := []User{u1, u2}
	fmt.Println("struct_arr")
	fmt.Println(_struct_arr(us)[0].Name)
	fmt.Println(us[0].Name)

}

func _int(p int) int {
	p = 100
	return p
}

func _float(p float64) float64 {
	p = 0.35
	return p
}

func _bool(p bool) bool {
	p = false
	return p
}

func _string(p string) string {
	p = "changed!"
	return p
}

func _array(p []int) []int {
	p[0] = 314158
	return p
}

func _struct(u User) User {
	u.Name = "Blank"
	return u
}

func _struct_arr(u []User) []User {
	u[0].Name = "Blank"
	return u
}

type User struct {
	Name string
	Age  int
}
