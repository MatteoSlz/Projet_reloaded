# Projet_reloaded

Pour commencer il vous faut un logiciel de code, de preference vscode. une fois sur vs code il faut installé le langage golang et l'outil git. Ensuite faire git clone https://github.com/MatteoSlz/Projet_reloaded

Lancement du programme :

Pour lancer le programme, il faut faire go run main.go "dossier de départ en .txt" "dossier qui va recevoir le nouveau texte.txt"

Explication du programme :

Le programme va prendre en entrée un texte et il va y appliquer plusieurs actions en fonction de ce même texte.

Tout d'abord, le programme va lire le fichier d'entrée en .txt et créer un tableau où il va ranger tous les mots séparés par un espace. Par exemple "bonjour (cap) je suis étudiant" va créer un tableau de string qui va contenir {"bonjour" , "(cap)", "je", "suis", "étudiant"}.

Puis notre programme va s'occuper de la ponctuation en l'indentant correctement.

Ensuite, on va parcourir ce tableau index par index et en fonction de la chaîne de caractères contenue à la position, alors on effectue telle action. Par exemple, au-dessus, lorsqu'on arrive à l'action "(cap)", alors on va appliquer une majuscule à la première lettre du mot précédent.
Bonjour va devenir "Bonjour" et "(cap)" va disparaître.

Si le programme tombe sur une action à faire et l'effectue alors l'index recule, ce qui lui permet de traiter plusieurs cas à la suite, par exemple "BONJOUR (low) (cap) Lilian"
va mettre BONJOUR en low puis va appliquer cap pour obtenir "Bonjour Lilian"

Fin du programme :

À la fin de notre programme, si tout le texte a été vérifié, alors on va lire notre tableau et on va stocker nos chaînes de caractères dans une variable 'resultat' puis on écrit resultat dans notre dossier .txt qui va recevoir le texte corrigé.