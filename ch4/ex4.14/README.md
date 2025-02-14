### Exercise 4.14:

Create a web server that queries GitHub once and then allows navigation of the list of bug reports,
milestones, and users.

---

### Упражнение 4.14:

Создайте веб-сервер, однократно запрашивающий информацию у GitHub, а затем позволяющий выполнять навигацию по списку
сообщений об ошибках, контрольных точек и пользователей.

---

Solution (решение):

```bash
$ go run issues repo:golang/go is:open json decoder
95 тем:
#71475  oiweiwei encoding/json: improve decoder alloc count
#71618  romshark image/jpeg: improve decoder performance by ~12%
#48298     dsnet encoding/json: add Decoder.DisallowDuplicateFields
#71497     dsnet proposal: encoding/json/v2: new API for encoding/json
#56733 rolandsho proposal: encoding/json/jsontext: add WithByteLimit and
#69449   gazerro encoding/json: Decoder.Token does not return an error f
#42571     dsnet encoding/json: clarify Decoder.InputOffset semantics
#5901        rsc encoding/json: allow per-Encoder/per-Decoder registrati
#11046     kurin encoding/json: Decoder internally buffers full input
#58649 nabokihms encoding/json: show nested fields path if DisallowUnkno
#6647    btracey x/pkgsite: display type kind of each named type
#67525 mateusz83 encoding/json: don't silently ignore errors from (*Deco
#36225     dsnet encoding/json: the Decoder.Decode API lends itself to m
#29035    jaswdr proposal: encoding/json: add error var to compare  the
#61627    nabice x/tools/gopls: feature: CLI syntax for renaming by iden
#34543  maxatome encoding/json: Unmarshal & json.(*Decoder).Token report
#26946    deuill encoding/json: clarify what happens when unmarshaling i
#32779       rsc encoding/json: memoize strings during decode
#40128  rogpeppe proposal: encoding/json: garbage-free reading of tokens
#59053   joerdav proposal: encoding/json: add a generic Decode function
#40982   Segflow encoding/json: use different error type for unknown fie
#40127  rogpeppe encoding/json: add Encoder.EncodeToken method
#16212 josharian encoding/json: do all reflect work before decoding
#65691  Merovius encoding/xml: Decoder does not reject xml-ProcInst prec
#41144 alvaroale encoding/json: Unmarshaler breaks DisallowUnknownFields
#14750 cyberphon encoding/json: parser ignores the case of member names
#56332    gansvv encoding/json: clearer error message for boolean like p
#64847 zephyrtro encoding/json: UnmarshalJSON methods of embedded fields
#43513 Alexander encoding/json: add line number to SyntaxError
#22752  buyology proposal: encoding/json: add access to the underlying d
```

https://api.github.com/search/issues?q={query}{&page,per_page,sort,order}
