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
