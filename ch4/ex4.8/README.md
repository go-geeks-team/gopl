### Exercise 4.8.

Modify ```charcount``` to count letters, digits, and so on in their Unicode categories, 
using functions like ```unicode.IsLetter```.

---

### Упражнение 4.8.

Измените ```charcount```  так, чтобы программа подсчитывала количество букв, цифр и прочих категорий Unicode с использованием 
функции наподобие ```unicode.IsLetter```. 

---

Run:
```bash
$ go run ./charcount.go < dict.txt
```

Output:
```bash
$ go run ./charcount.go < dict.txt
graphic count   count
'8'     1
'd'     1
'D'     1
'a'     3
'2'     6
'l'     1
'1'     2
'A'     2
'n'     2
'i'     2
's'     1
' '     3
'0'     3
upper count     count
'D'     1
'A'     2
spaces  count
' '     3
'\n'    2
letter  count
'D'     1
'i'     2
'a'     3
's'     1
'A'     2
'l'     1
'n'     2
'd'     1
digit   count
'2'     6
'0'     3
'1'     2
'8'     1
rune    count
'D'     1
'a'     3
'1'     2
'l'     1
'n'     2
' '     3
'0'     3
'A'     2
'i'     2
'\n'    2
'8'     1
's'     1
'2'     6
'd'     1

len     count
'1'     2
'l'     1
'n'     2
'D'     1
'a'     3
'A'     2
' '     3
'0'     3
'8'     1
'i'     2
'\n'    2
'd'     1
's'     1
'2'     6

len     count
1       30
2       0
3       0
4       0
```