# go-reloaded

Outil en ligne de commande, écrit en Go, qui corrige automatiquement un texte.
Il lit un fichier d'entrée, applique des règles de correction, puis écrit le résultat dans un fichier de sortie.

## Règles appliquées

| Règle | Exemple | Résultat |
|---|---|---|
| `(hex)` : hexadécimal → décimal | `1E (hex) files` | `30 files` |
| `(bin)` : binaire → décimal | `10 (bin) years` | `2 years` |
| `(up)` : majuscules | `go (up)` | `GO` |
| `(low)` : minuscules | `SHOUTING (low)` | `shouting` |
| `(cap)` : première lettre en majuscule | `bridge (cap)` | `Bridge` |
| `(up, n)`, `(low, n)`, `(cap, n)` : sur les n mots précédents | `so exciting (up, 2)` | `SO EXCITING` |
| Ponctuation `. , ! ? : ;` collée au mot précédent | `there ,and` | `there, and` |
| Groupes de ponctuation | `thinking ...` | `thinking...` |
| Apostrophes collées aux mots qu'elles entourent | `' awesome '` | `'awesome'` |
| `a` → `an` devant une voyelle ou un h | `a hour` | `an hour` |

## Prérequis

- Go 1.25 ou plus récent
- Aucune bibliothèque externe : uniquement la bibliothèque standard

## Récupérer le projet

```bash
git clone https://github.com/8r1c3-333/go-reloaded.git
cd go-reloaded
```

## Compiler

```bash
go build -o go-reloaded .
```

## Exécuter

```bash
go run . <fichier_entree> <fichier_sortie>
```

ou, après compilation :

```bash
./go-reloaded <fichier_entree> <fichier_sortie>
```

- Le fichier d'entrée doit exister. Le fichier de sortie est créé, ou écrasé s'il existe déjà.
- Avec un mauvais nombre d'arguments, le programme affiche `usage: go run . <in> <out>` et s'arrête (code 1).
- Si le fichier d'entrée est introuvable, il affiche l'erreur et s'arrête (code 1).
- Les retours à la ligne du fichier d'entrée sont conservés.

## Lancer les tests

```bash
go test ./...
```

Pour voir le détail :

```bash
go test -v ./...
```

Les tests couvrent chaque règle, les cas de référence du sujet et des cas particuliers : marqueur en début de ligne, `n` plus grand que le nombre de mots, `n = 0`, accents, conversion impossible, apostrophe à l'intérieur d'un mot, plusieurs lignes, texte vide.

## Exemple

`sample.txt` :

```
it (cap) was the worst of times (up) , it was
Simply add 42 (hex) and 10 (bin)
bearing a untold story ,what do you think ?
Ready set go (up).
This is so exciting (up, 2)
I am exactly how they describe me: ' awesome '
```

```bash
go run . sample.txt result.txt
cat result.txt
```

`result.txt` :

```
It was the worst of TIMES, it was
Simply add 66 and 2
bearing an untold story, what do you think?
Ready set GO.
This is SO EXCITING
I am exactly how they describe me: 'awesome'
```

## Organisation du code

| Fichier | Rôle |
|---|---|
| `main.go` | Vérifie les arguments, lit le fichier d'entrée, appelle `Process`, écrit le fichier de sortie |
| `process.go` | `Process` : applique toutes les règles, ligne par ligne |
| `process_test.go` | Tests automatiques de `Process` |

Ordre des traitements pour chaque ligne : marqueurs → ponctuation → apostrophes → articles.

## Limite connue

Un nombre décimal comme `3.14` devient `3. 14`, car la règle de ponctuation ajoute un espace après le point.