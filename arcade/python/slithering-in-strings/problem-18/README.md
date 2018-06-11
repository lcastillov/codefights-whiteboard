# Competitive Eating

## Solution

This problem can be solved in two steps:

1. Convert `t` from float to string with precision `precision`.
2. Center the converted number inside `width` spaces.

For `t = 3.1415`, `width = 10` and `precision = 2` the output should be

```python
competitiveEating(t, width, precision) = "   3.14   "
```

If we want to round `t` to `2` digits, for instance, we can do it at least in two ways:

- `'{:.2f}'.format(t)` (_new-style_ formatting [[1]][format])
- `'%.2f' % t` (_old-style_ formatting)

In order to center a string inside `width` spaces we can use the _new-style_ formatting or the `str` builtin method `center` [[2]][center].

- `'{:^10}'.format("3.14")`
- `"3.14".center(10)`

As we want to use a variable precision and width, we can combine _old_ and _new-style_ formatting. Note that we can round and center at the same time using _new-style_ formatting `'{:^<width>.<precision>f}'`.

```python
def competitiveEating(t, width, precision):
    return ('{:^%d.%df}' % (width, precision)).format(t)
```

A shorter solution using _new-style_ formatting only:

```python
def competitiveEating(t, width, precision):
    return '{:^{}.{}f}'.format(t, width, precision)
```

`t` is assigned to the outer curly braces, `width` to the inner first `{}` braces and `precision` to the second one.

[format]:https://www.python.org/dev/peps/pep-3101/
[center]:https://docs.python.org/3/library/stdtypes.html#str.center