# Word Power

## Solution

For this problem we need to create a dictionary `num` that maps every letter to its position:

- `a -> 1`
- `b -> 2`
- ...
- `z -> 26`

### 1. Using `dict` constructor

```python
num = dict((c, ord(c)-96) for c in "abcdefghijklmnopqrstuvwxyz")
```

### 2. Using `{}` constructor

```python
num = {c: ord(c)-96 for c in "abcdefghijklmnopqrstuvwxyz"}
```

### 3. Using `dict` and `zip`

```python
num = dict(zip("abcdefghijklmnopqrstuvwxyz", range(1, 27)))
```
