### Упражнение 1.10

Найдите веб-сайт, который содержит большое количество данных. Исследуйте работу кеширования путём двухкратного запуска
```fetchall``` и сравнения времени запросов. Получите ли вы каждый раз одно и то же содержимое?
Измените ```fetchall``` так, чтобы вывод осуществлялся в файл и чтобы затем можно было его изучить.

Exercise 1.10: 
Find a web site that produces a large amount of data. Investigate caching by running
fetchall twice in succession to see whether the reported time changes much. Do you get the same
content each time? Modify fetchall to print its output to a file so it can be examined.


Результат двухкратного запуска fetchall:

Первый запуск:
```bash
$ ./fetchall https://youtube.com
1.56s  545594 https://youtube.com
1.56s elapseed
```

Второй запуск:
```bash
$ ./fetchall https://youtube.com
1.46s  531168 https://youtube.com
1.46s elapseed
```
