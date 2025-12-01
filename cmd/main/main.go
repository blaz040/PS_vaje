package main

import (
	"context"
	"log"
	"os"

	redovalnica "github.com/blaz040/PS_vaje/pckg/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "Redovalnica",
		Usage: "Shows the basic Redovalnica usage",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "The minimal number of grades to pass",
				Value: 1,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "The minimal possible grade",
				Value: 0,
			},

			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "The maximal possible grade",
				Value: 10,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")
			runRedovalnica(stOcen, minOcena, maxOcena)
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func runRedovalnica(stOcen int, minOcena int, maxOcena int) {

	redovalnica.Init(stOcen, minOcena, maxOcena)

	studenti := make(map[string]redovalnica.Student)

	studenti["1"] = redovalnica.Student{Ime: "Jože", Priimek: "Pločnik", Ocene: []int{9, 9, 9}}
	studenti["2"] = redovalnica.Student{Ime: "Blaž", Priimek: "Pirc", Ocene: []int{10, 10, 10}}
	studenti["3"] = redovalnica.Student{Ime: "Maj", Priimek: "Zorko", Ocene: []int{8, 7, 6}}

	redovalnica.DodajOceno(studenti, "2", 6)
	redovalnica.DodajOceno(studenti, "1", 5)
	redovalnica.DodajOceno(studenti, "1", 11)

	redovalnica.IzpisRedovalnice(studenti)
	redovalnica.IzpisiKoncniUspeh(studenti)

}
