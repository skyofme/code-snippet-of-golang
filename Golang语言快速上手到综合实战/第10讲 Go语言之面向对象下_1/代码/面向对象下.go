package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (person Person) getNameAndAge() (string, int) {
	return person.name, person.age
}

type Student struct {
	Person
	speciality string
}

func (student Student) getSpeciality() string {
	return student.speciality
}

type Animal interface {
	Fly()
	Run()
}

type Animal2 interface {
	Fly()
}

type Bird struct {
}

func (bird Bird) Fly() {
	fmt.Println("Bird is flying!!!!")
}

func (bird Bird) Run() {
	fmt.Println("Bird is running!!!!")
}

func main() {
	/*student := new(Student)
	student.name = "zhangsan"
	student.age = 26
	name, age := student.getNameAndAge()

	speciality := student.getSpeciality()

	fmt.Println(name, age)*/
	/*var animal Animal
	var animal2 Animal2

	bird := new(Bird)
	animal = bird
	animal2 = animal
	//
	animal2.Fly()*/
	//animal = bird //把类实例直接赋值给接口
	//animal2 = bird

	//animal2.Fly()
	//animal.Fly()
	//animal.Run()

	var v1 interface{}
	v1 = "zhangsan"

	switch v1.(type) {
	case int:
	case float32:
	case float64:
		fmt.Println("this is float64")
	case string:
		fmt.Println("this is string")
	}
}
