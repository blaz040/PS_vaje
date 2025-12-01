package redovalnica

import "fmt"

// Structure defining one instance of a Student
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

var stOcen int = 1
var minOcena int = 1
var maxOcena int = 10

// Initiallizes the parameters for redovalnica
func Init(m_stOcen int, m_minOcena int, m_maxOcena int) {
	stOcen = m_stOcen
	minOcena = m_minOcena
	maxOcena = m_maxOcena
}

// adds grade to student map with key VpisnaStevilka
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Študenta ni na seznamu :(")
		return
	}
	if minOcena <= ocena && ocena <= maxOcena {
		student.Ocene = append(student.Ocene, ocena)
		studenti[vpisnaStevilka] = student

	} else {
		fmt.Println("Ocena ni v pravem rangu")
	}

}

// average grade of a student with key
func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Študent ne obstaja")
		return -1.0
	}
	var sum float64
	for _, v := range student.Ocene {
		if v < 6 {
			return 0.0
		}
		sum += float64(v)
	}
	sum /= float64(len(student.Ocene))

	return sum
}

// prints all students in Redovalnica
func IzpisRedovalnice(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for v := range studenti {
		var student = studenti[v]
		fmt.Printf("%s - %s %s: %d\n", v, student.Ime, student.Priimek, student.Ocene)
	}

}

// prints the final grade of all stundets in map
func IzpisiKoncniUspeh(studenti map[string]Student) {
	for v := range studenti {
		var student = studenti[v]
		var povp_ocena = povprecje(studenti, v)
		var uspeh = "Povprečen študent"

		switch {
		case len(student.Ocene) < stOcen:
			uspeh = "Neuspešen študent"
		case povp_ocena >= 9:
			uspeh = "Odločen študent!"
		case povp_ocena < 6:
			uspeh = "Neuspešen študent"
		}
		fmt.Printf("%s %s: povprečna ocena %0.1f -> %s\n", student.Ime, student.Priimek, povp_ocena, uspeh)
	}
}
