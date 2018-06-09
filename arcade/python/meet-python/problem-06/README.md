# Modulus

## Solution

The floor division and modulo operators are connected by the following identity: `(x % y) == x - (x // y) * y`.

In order to solve the problem we need to check whether `x` is an integer or not using `type(x) is int` or `isinstance(x, int)`.
