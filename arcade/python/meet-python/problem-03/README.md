# Special Conditional

One of the given statements doesn't work the same way others do. Which one?

1. `a == (not b)`
2. `a == not b`
3. `not (a == b)`
4. `not a == b`

Assume that `a` and `b` are boolean variables.

## Solution

`not` has a lower priority than non-Boolean operators, so `not a == b` is interpreted as `not (a == b)`, and `a == not b` is a syntax error [[1]][oolean-operations-and-or-not]. So, the second statement doesn't work the same way others do.

[oolean-operations-and-or-not]: https://docs.python.org/3/library/stdtypes.html#boolean-operations-and-or-not
