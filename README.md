# Embracing Iterators in GO

Появились в версии 1.23.

Итератор это просто функция, но с определенной сигнатурой. 

В самом простом и наглядном виде:

```go
func(yield func(int) bool)
```

Итератор принимает параметром функцию. 
Принято называть такой итератор как yield.

Итератор можно возвращать из функции. 
В этом случае, yield будет вызываться каждый раз, выбрасывая в range успешное значение: 

```go
import ("iter")

func main() {
	for v := range revertSliceInt([]int{1, 2, 3, 4, 5}) {
    // .. любая логика с v
  }
}

func revertSliceInt(sl []int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := len(sl) - 1; i >= 0; i-- {
			if !yield(sl[i]) {
				return
			}
		}
	}
}
```

a) Для работы итератора нужно импортировать библиотеку "iter"
b) Возвращаем из функции итератор (iter)
c) iter.Seq указывает, итератор какого типа будет использован (Seq|Seq2)
d) Создаётся условие, при котором yield прекращает работу
e) Как только все успешные значения будут возвращены, цикл прекращается

## Функции принимающие итератор

Описанная выше функция возвращает в каждый момент вызова итератор. 

Также можно описать функцию, которая этот итератор будет принимать и обрабатывать в соответствии с нужной логикой

```go
func PrintItems(s iter.Seq[int]) {
	for v := range s {
		// .. любая логика с v
	}
}
```

## Generic in Iterators

В примерах выше использовался slice с типом int. 

Но что, если нам нужно использовать слайсы из разных типов? 

В этом случае можно использовать дженерик:

```go
func main() {
  sl := []string{"One", "Two", "Three"}
  // Можно использовать теперь и slice int:
	// sl := []int{1, 2, 3, 4, 5}

	PrintItems(revertSliceInt(sl))
}

func PrintItems[V any](s iter.Seq[V]) {
	for v := range s {
		fmt.Println(v)
	}
}

func revertSliceInt[V any](sl []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := len(sl) - 1; i >= 0; i-- {
			if !yield(sl[i]) {
				return
			}
		}
	}
}
```

# Standard library Iterators in GO

Введение итераторов позволило добавить ряд новых API для slice && map

## slices iterator

Понадобится import {"slices"}

```go
func main() {
  s := []int{1, 2, 3, 4, 5}

  // ---
	for k, v := range slices.All(s) {
		fmt.Printf("%d: %d\n", k, v) // Output: 0: 1, 1: 2, 2: 3 ..
	}

  // ---
  for v := range slices.Values(s) {
		fmt.Println(v) // Output: 1, 2, 3 ..
	}

  // ---
  result := slices.Collect(slices.Values(s))
	fmt.Println(result) // [1,2,3,4,5]
}
```

## maps iterators

Понадобится import {"maps"}

```go
func main() {
	m := map[int]string{
		1: "first",
		2: "second",
		3: "third",
		4: "fourth",
	}

	for k, v := range maps.All(m) {
		fmt.Printf("%d: %s\n", k, v)
	}
}
```

Аналогично можно использолвать maps.Keys() и maps.Values()

# Дополнительная информация

[Iterators in Go](https://bitfieldconsulting.com/posts/iterators)
