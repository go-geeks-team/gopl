### Упражнение 1.5

Измените палитру программы lissajous так, чтобы изображение было зелёного цвета на чёрном фоне,
чтобы быть более похожим на экран осциллографа. Чтобы создать веб-цвет #RRGGBB, воспользуйтесь инструкцией 
color.RGBA{0xRR, 0xGG, 0xBB, 0xff}, в которой каждая пара шестнодцатеричных цифр представляет яркость красного,
зелёного, и синего компонентов пиксеоя.

```bash
$ go run ./lissajous.go > out0.png
```