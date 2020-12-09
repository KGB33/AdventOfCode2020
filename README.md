# Advent of Code 2020

My solutions to [AdventOfCode2020](https://adventofcode.com/2020https://adventofcode.com/2020)

I am using this problem set to practice Go. My notes are below.


# Notes

## Files

A major part of this challenge is reading input files.
An excellent stack overflow response can be found [here](https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go/16615559#16615559)

```Go
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open("/path/to/file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
```

## Logical Operators

Go has `&&`, `||` and, `!`. It does not have a XOR operator because
`bool XOR bool` == `bool != bool`

## String Parsing

### Split

You can split strings using a regex pattern.
For example, on day two we had to split a string in the following format:
```
"3-6 s: some_text" --> [3, 6, "s", "some_text"]
```

We can acheave this split using the regex `-| |: ` (*dash* OR *space* OR *colon* *space*).
Then converting the first two indexes to ints.

For day 4 I had to split a multiline string at the white space. For this use-case
`strings.Fields(s)` worked perfectly.

### Iteration/runes

You can iterate over the runes in a string using a simple for loop.

```Go
for index, rune := range string {
	// Stuff
}
```

You can also cast a string into a slice of runes.

```Go
slice_runes = []rune(string)
```

Converting runes to the integers they represent is a little more tricky.
Calling `int('3')` will return `51` the integer representation of a UTF-8 3.
To convert a rune to the value it represents use `int(r - '0')` where `r` is
any rune '0'->'9'. This works because the 0-9 are all sequential.
Such that `'0'` (UTF-8: 48) minus itself equals UTF-8: 0. See the table below.

| Rune | UTF-8 | int(r - '0') |
|:----:|:-----:|:------------:|
|  '0' |   48  |       0      |
|  '1' |   49  |       1      |
|  '2' |   50  |       2      |

# Loops

## Modular Logic

On day three we had a pattern that repeated infinitely in one direction.
As such to iterate over it in that direction we could just loop back around.
I.e `last_index + 1 == first_index`. An easy way to do this mathematically would
be by using the modulo operator.

Given some slice s, an integer i.

`i = i + 1 % len(s)` will always be in s.

# Switch statements

## if-else-if-else shorthand

You can use switch statements as a shorthand for extended if else-if else chains.

## One of

In python to check if a value, n, was one of several values we would do.

```Python
if n in (possible, values,):
	# do stuff
else:
	# do other stuff
```

In go you can use a switch statement.


```Go
switch n:
	case possible, values:
		// do stuff
	default:
		// do other stuff
```

# Sorting
You can sort splices using the `sort` package. There is a function for each type of
splice.

# DataTypes

## Maps

Maps are automatically passed by reference in Go. I used this in combination with the
set implementation described below to keep track of the possible bags in day 7 part 1.

## Structs

You can add a `String()` method to a struct that will automatically be called
when converting to a string. This is similar to `__str__` in python.

## Set

There is no `set` datatype in Go. However, the keys in a map must be unique,
so we can build a map where the keys are the elements of our "set" and the
values are booleans. Because the values are booleans, it makes asking if 
`x exists in y` easy. [source](https://yourbasic.org/golang/implement-set/)

```Go
set := make(map[string]bool) // New empty set
set["Foo"] = true            // Add
for k := range set {         // Loop
    fmt.Println(k)
}
delete(set, "Foo")    // Delete
size := len(set)      // Size
exists := set["Foo"]  // Membership
```

## Counter (?)

We can also use a map as a counter, Just like with sets we use the
key to be the value we care about, and its corresponding value is an integer
denoting its frequency.

# Misc

## Copy

One way to copy slices of structs is to use `append`.

```Go
copy := append([]Type{}, source...)
```
This creates a new slice of types, then appends the unpacked values from the source slice.

# Interesting Problem Details

## Day 05
An ordinary binary search didn't work due to indexing errors. Instead, rounding with
`math.Float` and `math.Ceil` gave the correct answers.


