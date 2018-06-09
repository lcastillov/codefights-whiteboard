# Language Differences

Consider the following two codes:

Python

```python
def division(x, y):
    return x // y
```

Java

```java
int division(int x, int y) {
    return x / y;
}
```

You noticed that the functions aren't quite the same: they won't produce the same result for some valid values of `x` and `y`. Which ones?

1. `x = -10, y = -3`
2. `x = 15, y = -4`
3. `x = -8, y = 2`
4. `x = 5, y = 10`
5. `x = 17, y = 13`

## Solution

The `//` in python is the _floor division_ [[1]][floor-division], so `x // y` will be the greatest integer less than or equal to `x / y`. The real division for `x = 15` and `y = -4` is `x / y = -3.75`, thus `x // y = -4`. The `/` operator is a "classic" division in Java, C++, C# and other languages.

[floor-division]:https://www.python.org/dev/peps/pep-0238/
