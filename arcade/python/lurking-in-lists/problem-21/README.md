# List Multiplication

Assume that you come across the following piece of code:

`c = a * b`

Can it stand for:

1. list `a` repeated `b` times;
2. list `b` repeated `a` times?

## Solution

Both options are correct:

```python
>>> [1, 2] * 3
[1, 2, 1, 2, 1, 2]
>>> 3 * [1, 2]
[1, 2, 1, 2, 1, 2]
```
