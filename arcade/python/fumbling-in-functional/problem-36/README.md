# Fix Result

## Solution

Use `map` [[1]][map] function to solve this problem.

### Example # 1

Convert every character to its ordinal.

```python
>>> s = "Hello"
>>> list(map(ord, s))
[72, 101, 108, 108, 111]
```

### Example # 2

Given a list of strings, get their lengths.

```python
>>> l = ["Never", "argue", "with", "the", "data"]
>>> list(map(len, l))
[5, 5, 4, 3, 4]
```

### Example # 3

Custom function that takes an integer `n` and return `2`<sup>`n`</sup>.

```python
>>> def pow2(n):
...    return 2**n
>>> l = [0, 1, 2, 3]
>>> list(map(pow2, l))
[1, 2, 4, 8]
```

[map]:https://docs.python.org/3/library/functions.html#map
