# Simple syntax

- Division is not integer division by default
	- 3/3 == 1.0
- Can do integer divsion using 3 `div` 3 == 1
	- it works for any function that takes 2 arguments
	- if you had a function `let f a b = show a ++ show b` then you could do 3 `f` 4.
- Boolean are primitives: True, False
- not True -> not is a function that takes one value
- "" == String, '' == Character
- Concatenation is ++, "Hello " ++ "World"
- A string is a list of characters, "This is a String" !! 0 == 'T'
- Comments: -- for single line, {- -} for multi-line
- List: [1, 2, 3, 4, 5] == [1..5]
	- [1..] !! 999 == 1000
- [1..5] ++ [1..4]
	- + is only defined for numbers
- Adding something to list you could do 5:[1..2]
- List operations: head [1..5], tail [1..5], init [1..5], last [1..5]