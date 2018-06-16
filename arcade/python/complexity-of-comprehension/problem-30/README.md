# Create Spiral Matrix

The aim of this problem is to create a 2D matrix of `NxN`.

## Solution

We need to iterate `n` times and for every iteration create a row. This can be done using _list comprehensions_ [[1]][lc] as follow:

```python
res = [[0] * n for _ in range(n)]
```

[lc]:https://docs.python.org/3/tutorial/datastructures.html#list-comprehensions