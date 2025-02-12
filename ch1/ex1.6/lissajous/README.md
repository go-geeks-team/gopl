### Упражнение 1.6.

Измените программу lissajous так, чтобы она генерировала изображения разных цветов, добавляя в палитру palette больше значений,
а затем выводя их путём изменения третьего аргумента функции SetColorIndex некоторым нетривиальным образом.


```bash
$ go run ./lissajous.go
```

<img src="./1.png" alt="1">
<img src="./2.png" alt="2">
<img src="./3.png" alt="3">


Справка по функции ```DirFS```:
DirFS returns a file system (an fs.FS) for the tree of files rooted at the directory dir.
Note that DirFS("/ prefix") only guarantees that the Open calls it makes to the operating system will begin with "/ prefix": DirFS("/ prefix").Open("file") is the same as os.Open("/ prefix/ file"). So if / prefix/ file is a symbolic link pointing outside the / prefix tree, then using DirFS does not stop the access any more than using os.Open does. Additionally, the root of the fs.FS returned for a relative path, DirFS("prefix"), will be affected by later calls to Chdir. DirFS is therefore not a general substitute for a chroot-style security mechanism when the directory tree contains arbitrary content.
The directory dir must not be "".
The result implements io/ fs.StatFS, io/ fs.ReadFileFS and io/ fs.ReadDirFS.