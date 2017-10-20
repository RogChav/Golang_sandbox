package main

func main() {
	// can be defined without writing out the fist or last name variable but the first and last name are assumed to be written in the same order as it expects them to come in in the initiation of type, which is firstName then lastName

	// roger := person{firstName: "Roger", lastName: "Chavez"}
	// fmt.Println(roger)

	// a third way to

	roger := person{
		firstName: "Roger",
		lastName:  "Chavez",
		contactInfo: contactInfo{
			email:   "roger@rogerchavez.com",
			zipCode: 92691,
		},
	}
	rogerPointer := &roger
	rogerPointer.updateName("Rogelio")
	roger.print()
}
