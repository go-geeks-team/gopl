### Exercise 4.13:

The JSON-based web service of the Open Movie Database lets you search https://omdbapi.com/ for a movie by name
and download its poster image. Write a tool poster that downloads the poster image for the movie named on the 
command line.

---

### Упражнение 4.13:

Веб-служба Open Movie Database https://omdbapi.com/ на базе JSON позволяет выполнять поиск фильма по названию и загружать его афишу. 
Напишите инструмент poster, который загружает афишу фильма по переданному в командной строке названию. 

---

Example1 (download movie poster):
```bash
$ ./poster -apikey="<API_KEY>" -title="titanik"
```

---

Example2 (download movie poster):
```bash
$ ./poster -apikey="<API_KEY>" -title="The Mask"
```

---

Example3 (download movie poster):
```bash
$ ./poster -apikey="<API_KEY>" -title="2001: A Space Odyssey"
```
