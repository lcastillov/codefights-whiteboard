# Basic Conversion

Given a number `n` (as string) and a base `x`, converts the number from base `x` to base `16`.

## Solution

First, we convert `n` to base `10` using `int(n, x)` [[1]][int], then we can use one of the following alternatives:

* `hex(int(n, x))[2:]` [[2]][hex]
* `'%x' % int(n, x)`
* `'{0:x}'.format(int(n, x))`
* `format(int(n, x), 'x')` [[3]][format]

The function `hex` convert an integer number to a lowercase hexadecimal string prefixed with `"0x"`.

[int]:https://docs.python.org/3/library/functions.html#int

[hex]:https://docs.python.org/3/library/functions.html#hex

[format]:https://docs.python.org/3/library/functions.html#format