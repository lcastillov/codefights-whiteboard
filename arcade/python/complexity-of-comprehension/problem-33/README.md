# Cool Pairs

## Solution

First, we are going to generate all possible pairs using _list comprehensions_.

```python
>>> a = [1, 2, 3]
>>> b = [4, 5, 6]
>>> [(i, j) for i in a for j in b]
[(1, 4), (1, 5), (1, 6), (2, 4), (2, 5), (2, 6), (3, 4), (3, 5), (3, 6)]
```

Then, we are going to filter all those pairs who meets the criteria `(i * j) % (i + j) == 0` and store the sum `i + j`.

```python
>>> a = [1, 2, 3]
>>> b = [4, 5, 6]
>>> [(i, j) for i in a for j in b if (i * j) % (i + j) == 0]
[(3, 6)]
>>> [i + j for i in a for j in b if (i * j) % (i + j) == 0]
[9]
```

We can use now the set constructor `{}` [[1]][set] to avoid repeated sums:

```python
def coolPairs(a, b):
    uniqueSums = {i+j for i in a for j in b if (i * j) % (i + j) == 0}
    return len(uniqueSums)
```

[set]:https://docs.python.org/3/tutorial/datastructures.html#sets
