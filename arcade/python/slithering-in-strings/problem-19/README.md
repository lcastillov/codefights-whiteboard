# New Style Formatting

## Solution

We are going to take care of `%%` first, and then we will replace every `%<type>` where `type ::=  "b" | "c" | "d" | "e" | "E" | "f" | "F" | "g" | "G" | "n" | "o" | "s" | "x" | "X"` [[1]][format-specs].

The solution would be:

1. Replace every `'%%'` with `'{%}'`. We can do this because the original string doesn't contain curly brackets.
2. Replace every `'%<type>'` with `'{}'`.
3. Finally, replace `'{%}'` with `'%'`.

[format-specs]:https://docs.python.org/2/library/string.html#format-specification-mini-language
