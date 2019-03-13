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
	var animal Animal
	var animal2 Animal2

	bird := new(Bird)
	//animal = bird
	//animal2 = bird

	//1
	//animal = animal2
	//2
	animal = bird
	animal2 = animal
	animal2.Fly()
	//animal2.Fly()

	//animal.Fly()
	//animal.Run()

}
