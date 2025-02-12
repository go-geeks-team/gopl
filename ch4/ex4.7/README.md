### Exercise 4.7:

Modify reverse to reverse the characters of a []byte slice that represents a UTF-8-encoded string, in-place.
Can you do it without allocating new memory?

---

### Упражнение 4.7:

Перепешите функцию reverse так, чтобы она без выделения дополнительной памяти обращала последовательность символов
среза ```[]byte```, который представляет строку в кодировке UTF-8. Сможете ли вы обойтись без выделения новой памяти? 

---

Test Output:

```bash
$ go test -v
=== RUN   TestReverseOld
    rev_test.go:13: [5 4 3 2 1 0]
--- PASS: TestReverseOld (0.00s)
=== RUN   TestReverseNew
    rev_test.go:24: [5 4 3 2 1 0]
--- PASS: TestReverseNew (0.00s)
PASS
ok      rev     0.002s
```
