### Exercise 4.6.

Write an  in-place function that squashes each run of adjacent Unicode spaces (see unicode.IsSpace) in a UTF-8-encoded
[]byte slice into a single ASCII space.

---

### Упражнение 4.6.

Напишите функцию, которая без выделения дополнительной памяти преобразует последовательности смежных 
пробельных символов Unicode (см. unicode.IsSpace) в срезе []byte в кодировке UTF-8 в один пробел ASCII.

---

https://github.com/go-geeks-team/gopl/blob/c10229bf7606fe6f0f11c746f2c6f87801bb8053/ch4/ex4.6/strutil/stringutil.go#L1-L16