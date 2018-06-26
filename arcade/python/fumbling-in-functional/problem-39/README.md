# Least Common Denominator

For the given list of `denominators`, find the least common denominator by finding their LCM.

For `denominators = [2, 3, 4, 5, 6]`, the output should be `leastCommonDenominator(denominators) = 60`.

## Solution

We need to solve this problem using the function `reduce` [[1]][reduce].

According to the `reduce` documentation:

`functools.reduce(function, iterable[, initializer])`

Apply _function_ of two arguments cumulatively to the items of _sequence_, from left to right, so as to reduce the sequence to a single value.

```python
reduce(lambda n, m: n*m/gcd(n,m), denominators, 1)
```

### Example # 1

Calculate the sum of the elements in a list.

```python
>>> from functools import reduce
>>> l = [1, 2, 3]
>>> reduce(lambda acc, u: acc+u, l, 0)
6
```

[reduce]:https://docs.python.org/3/library/functools.html#functools.reduce