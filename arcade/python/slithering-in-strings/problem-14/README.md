# Convert Tabs

Implement a function that, given a piece of `code` and a positive integer `x` will turn each tabulation character in code into `x` whitespace characters.

## Solution

To solve this problem we are going to use the string method `replace` [[1]][replace], `code.replace("\t", " " * x)`.

The multiplication of a string `s` and a positive integer `x` is the concatenation of `s`, `x` times.

[replace]:https://docs.python.org/3/library/stdtypes.html#str.replace