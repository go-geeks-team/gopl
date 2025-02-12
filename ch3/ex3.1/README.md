### Упражнение 3.1

Если функция f возвращает значение float64, не являющееся конечным, SVG-файл содержит неверные элементы <polygon>
(хотя многие визуализаторы SVG успешно обрабатывают эту ситуацию). Измените программу так, 
чтобы некорректные многоугольники были опущены.

### Exercise 3.1

If the function f returns a non-finity float64 value, the SVG file will contain invalid <polygon> elements
(although many SVG renderers handle this gracefully). Modify the program to skip invalid polygons.