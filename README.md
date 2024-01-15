# push-swap


**Project push-swap aims to sort the numbers given based on a certain** `**Algo**`**, and then print them**

Here is the file structure

```

.
├── Algo
│   └── Algorithem.go
│   └── AppendNumbers.go
│   └── FindValues.go
│   └── FullSort.go
│   └── SmallSorta&B.go
│   └── Stract.go
├── checker
│   └── main.go
├── pushSwap
│   └── main.go
├── audit.sh
├── build.sh
├── go.mod
├── Makefile
└── README.md
```

_**To Run this project**_

1.  `make Push` after cloning and navigating for ./push-swap or `go run . "numbers` after `cd pushSwap`
2.  `make Checek` after cloning and navigating for ./checker or `go run . "numbers` after `cd checker`
2.  `make Both` after cloning and navigating for buliding both of them
3.  `./"nameOfTheBot" "numbers"` (if u used the first method).
4.  use `| wc -l` after the output to show the count of the moves.

**bulid:**
```
`go build -o checker main.go`
`go build -o push-swap main.go`
```

**These are the instructions that you can use to sort the stack:**

- `pa` push the top first element of stack `b` to stack `a`
- `pb` push the top first element of stack `a` to stack `b`
- `sa` swap first 2 elements of stack `a`
- `sb` swap first 2 elements of stack `b`
- `ss` execute `sa` and `sb`
- `ra` rotate stack `a` (shift up all elements of stack `a` by 1, the first element becomes the last one)
- `rb` rotate stack `b`
- `rr` execute `ra` and `rb`
- `rra` reverse rotate `a` (shift down all elements of stack `a` by 1, the last element becomes the first one)
- `rrb` reverse rotate `b`
- `rrr` execute `rra` and `rrb`

Project was Written in **Golang** using only allowed libraies

**Authors**

1.  Fatima Abbas **fatabbas**
2.  Mariam Abbas **maabbas**
3.  Ruqaya Helal **rhelal**
