# Chess Teams

## Solution

We can use `zip` [[1]][zip] function to solve this problem.

```python
>>> list(zip([1, 2, 3], [4, 5, 6]))
[(1, 4), (2, 5), (3, 6)]
```

__note__: In python2 `zip` returns a list, however, we need to convert the output of `zip` to a list in python3 because it returns a generator instead.

[zip]:https://docs.python.org/3/library/functions.html#zip