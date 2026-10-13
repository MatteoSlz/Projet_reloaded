package main 
import (
	"fmt"
	"os"
)
	func main() {

	texte := os.Args[1]
	contenu, err := os.ReadFile(texte)
	if err != nil{
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}
	fmt.Println(string(contenu))
	}