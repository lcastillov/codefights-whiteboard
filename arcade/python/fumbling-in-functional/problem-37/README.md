# College Courses

## Solution

Use `filter` [[1]][filter] function to solve this problem.

### Example # 1

Filter even integers.

```python
>>> l = [1, 3, 4, 5, 6, 8, 11]
>>> list(filter(lambda n: n % 2 == 0, l))
[4, 6, 8]
```

### Example # 2

Filter words with length at least `5`.

```python
>>> l = ["Never", "argue", "with", "the", "data"]
>>> list(filter(lambda s: len(s) >= 5, l))
['Never', 'argue']
```

### Example # 3

Filter strings consisting of digits only.

```python
>>> l = ["I'm", "26", "years", "old"]
>>> list(filter(str.isdigit, l))
["26"]
```

[filter]:https://docs.python.org/3/library/functions.html#filter
