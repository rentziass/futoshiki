# Futoshiki

This is a library to generate games of [Futoshiki](https://en.wikipedia.org/wiki/Futoshiki). Right it outputs the full board, with all the numbers and all the More/Less signs.

Example:

```golang
package main

import "github.com/rentziass/futoshiki"

func main() {
	f := futoshiki.GenerateFutoshiki(7)
	f.Print()
}
```

Will output something like:

```
4 > 1 < 7 > 3 > 2 < 5 < 6
⋁   ⋀   ⋁   ⋀   ⋀   ⋀   ⋁
2 < 6 > 4 < 5 > 3 < 7 > 1
⋀   ⋁   ⋁   ⋁   ⋀   ⋁   ⋀
6 > 5 > 2 > 1 < 4 > 3 < 7
⋁   ⋁   ⋁   ⋀   ⋀   ⋁   ⋁
5 > 4 > 1 < 7 > 6 > 2 < 3
⋀   ⋁   ⋀   ⋁   ⋁   ⋁   ⋀
7 > 2 < 3 < 6 > 5 > 1 < 4
⋁   ⋀   ⋀   ⋁   ⋁   ⋀   ⋁
3 < 7 > 5 > 4 > 1 < 6 > 2
⋁   ⋁   ⋀   ⋁   ⋀   ⋁   ⋀
1 < 3 < 6 > 2 < 7 > 4 < 5
```
