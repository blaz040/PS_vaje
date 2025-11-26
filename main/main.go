package main

import (
	redovalnica "github.com/blaz040/PS_vaje/pckg/redovalnica"
)

func main() {

	var studenti map[string]redovalnica.Student
	studenti = make(map[string]redovalnica.Student)

	studenti["1"] = redovalnica.Student{"Jože", "Pločnik", []int{9, 9, 9}}
	studenti["2"] = redovalnica.Student{"Blaž", "Pirc", []int{10, 10, 10}}
	studenti["3"] = redovalnica.Student{"Maj", "Zorko", []int{8, 7, 6}}

	redovalnica.DodajOceno(studenti, "2", 6)
	redovalnica.DodajOceno(studenti, "1", 5)
	redovalnica.DodajOceno(studenti, "1", 11)

	//fmt.Printf("%s povprečje: %0.2f\n", studenti["2"].ime, povprecje(studenti, "2"))
	//fmt.Printf("%s povprečje: %0.2f\n", studenti["1"].ime, povprecje(studenti, "1"))
	//fmt.Printf("Neko povprečje: %0.2f\n", povprecje(studenti, "4"))

	redovalnica.IzpisRedovalnice(studenti)
	redovalnica.IzpisiKoncniUspeh(studenti)
}
