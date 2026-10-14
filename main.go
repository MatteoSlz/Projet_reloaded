package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Converthex(nbr string) string {
	hex := "0123456789ABCDEF"
	resultat := 0
	for _, valeur := range nbr {
		if valeur >= 'a' && valeur <= 'f' {
			valeur -= 32
		}
		for i, valeurHex := range hex {
			if valeur == valeurHex {
				resultat = resultat*16 + i
				break
			}
		}
	}
	return strconv.Itoa(resultat)
}

func Convertbin(nbr string) string {
	resultat := 0
	tableau := strings.Split(nbr, "")

	for i := 0; i < len(tableau); i++ {
		if tableau[i] == "1" {
			resultat += int(math.Pow(2, float64(len(tableau)-1-i)))
		}
	}
	return strconv.Itoa(resultat)
}

func ToLower(s string) string {
	char := ""
	for _, r := range s {
		if (r <= 90) && (r >= 65) {
			char = char + string(int(r)+32)
		} else {
			char = char + string(r)
		}
	}
	return char
}

func ToUpper(s string) string {
	char := ""
	for _, r := range s {
		if (r <= 122) && (r >= 97) {
			char = char + string(int(r)-32)
		} else {
			char = char + string(r)
		}
	}
	return char
}

func ToCap(s string) string {
	char := ""
	modif := true
	for _, r := range s {
		if (r <= 122) && (r >= 97) && (modif) {
			char = char + string(int(r)-32)
			modif = false
		} else {
			char = char + string(r)
		}
	}
	return char
}

func main() {
	apostrophe_debut := false 
	action_effectuer := false
	count := ""
	nouveau_contenu := []byte{}
	texte := os.Args[1]

	contenu_byte, err := os.ReadFile(texte)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	contenu_string := string(contenu_byte)
	var tableau []string
	mot := ""

	for _, s := range contenu_string {
		if s != ' ' && s != '\n' && s != '\t' {
			mot = mot + string(s)
		} else {
			if mot != "" {
				tableau = append(tableau, mot)
				mot = ""
			}
		}
	}
	if mot != "" {
		tableau = append(tableau, mot)
	}

	for i := 1; i < len(tableau); i++ {
		action_effectuer = false
		if len(tableau[i]) > 0 && tableau[i][0] == ',' {
			tableau[i-1] = tableau[i-1] + ","
			count = ""
			for k := 1; k < len(tableau[i]); k++ {
				count = count + string(tableau[i][k])
			}
			tableau[i] = count
			action_effectuer = true

		}
		if len(tableau[i]) > 0 && tableau[i][0] == ';' {
			tableau[i-1] = tableau[i-1] + ";"
			count = ""
			for k := 1; k < len(tableau[i]); k++ {
				count = count + string(tableau[i][k])
			}
			tableau[i] = count
			action_effectuer = true
		}
		if len(tableau[i]) > 0 && tableau[i][0] == '.' {
			tableau[i-1] = tableau[i-1] + "."
			count = ""
			for k := 1; k < len(tableau[i]); k++ {
				count = count + string(tableau[i][k])
			}
			tableau[i] = count
			action_effectuer = true
		}
		if len(tableau[i]) > 0 && tableau[i][0] == ':' {
			tableau[i-1] = tableau[i-1] + ":"
			count = ""
			for k := 1; k < len(tableau[i]); k++ {
				count = count + string(tableau[i][k])
			}
			tableau[i] = count
			action_effectuer = true
		}
		if len(tableau[i]) > 0 && tableau[i][0] == '!' {
			tableau[i-1] = tableau[i-1] + "!"
			count = ""
			for k := 1; k < len(tableau[i]); k++ {
				count = count + string(tableau[i][k])
			}
			tableau[i] = count
			action_effectuer = true
		}
		if len(tableau[i]) > 0 && tableau[i][0] == '?' {
			tableau[i-1] = tableau[i-1] + "?"
			count = ""
			for k := 1; k < len(tableau[i]); k++ {
				count = count + string(tableau[i][k])
			}
			tableau[i] = count
			action_effectuer = true
		}
		if tableau[i] == "," {
			tableau[i-1] = tableau[i-1] + ","
			tableau[i] = ""
			action_effectuer = true
		}
		if tableau[i] == "?" {
			tableau[i-1] = tableau[i-1] + "?"
			tableau[i] = ""
			action_effectuer = true
		}
		if tableau[i] == "!" {
			tableau[i-1] = tableau[i-1] + "!"
			tableau[i] = ""
			action_effectuer = true
		}
		if tableau[i] == ";" {
			tableau[i-1] = tableau[i-1] + ";"
			tableau[i] = ""
			action_effectuer = true
		}
		if tableau[i] == "." {
			tableau[i-1] = tableau[i-1] + "."
			tableau[i] = ""
			action_effectuer = true
		}
		if tableau[i] == ":" {
			tableau[i-1] = tableau[i-1] + ":"
			tableau[i] = ""
			action_effectuer = true
		}

		j := i - 1
		if j >= 0 && tableau[j] == "" {
			j--
		}

		mot_precedent := j

		if tableau[i] == "(hex)" {
			tableau[mot_precedent] = Converthex(tableau[mot_precedent])
			tableau[i] = ""
			action_effectuer = true
		}
		if tableau[i] == "(bin)" {
			tableau[mot_precedent] = Convertbin(tableau[mot_precedent])
			tableau[i] = ""
			action_effectuer = true
		}
		if tableau[i] == "(up)" {
			tableau[mot_precedent] = ToUpper(tableau[mot_precedent])
			tableau[i] = ""
			action_effectuer = true
		}
		if tableau[i] == "(low)" {
			tableau[mot_precedent] = ToLower(tableau[mot_precedent])
			tableau[i] = ""
			action_effectuer = true
		}
		if tableau[i] == "(cap)" {
			tableau[mot_precedent] = ToCap(tableau[mot_precedent])
			tableau[i] = ""
			action_effectuer = true
		}
		if action_effectuer {
			i = i - 2
			if i < 0 {
				i = 0
			}
		}
	}
	resultat := ""
	for _, mot := range tableau {
		if mot != "" {
			if resultat != "" {
				resultat += " "
			}
			resultat += mot
			nouveau_contenu = []byte(resultat)
		}
	}
	os.WriteFile(os.Args[2], nouveau_contenu, 0644)
}