package main

func main() {

	var studenti map[string]Student
	studenti = make(map[string]Student)

	studenti["1"] = Student{ime: "Jože", priimek: "Pločnik", ocene: []int{9, 9, 9}}
	studenti["2"] = Student{"Blaž", "Pirc", []int{10, 10, 10}}
	studenti["3"] = Student{"Maj", "Zorko", []int{8, 7, 6}}

	DodajOceno(studenti, "2", 6)
	DodajOceno(studenti, "1", 5)
	DodajOceno(studenti, "1", 11)

	//fmt.Printf("%s povprečje: %0.2f\n", studenti["2"].ime, povprecje(studenti, "2"))
	//fmt.Printf("%s povprečje: %0.2f\n", studenti["1"].ime, povprecje(studenti, "1"))
	//fmt.Printf("Neko povprečje: %0.2f\n", povprecje(studenti, "4"))

	IzpisRedovalnice(studenti)
	IzpisiKoncniUspeh(studenti)
}
