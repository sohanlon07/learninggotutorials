package main

func main() {
	// type person struct {
	// 	name string
	// 	age int
	// 	pet string
	// }

	var person struct {
		name string
		age int
		pet string
	}

	person.name = "bob"
	person.age = 50
	person.pet = "dog"

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	println(person.name, person.age, person.pet)
	println(pet.name, pet.kind)

	// var fred person
	// bob := person{}
	// julia := person{
	// 	"Julia",
	// 	40,
	// 	"cat",
	// }

	// beth := person{
	// 	age: 30,
	// 	name: "Beth",
	// }

	// bob.name = "Bob"
	// fmt.Println(beth.name)
	// fmt.Println(fred.name)
	// fmt.Println(julia.name, julia.age, julia.pet)
}
