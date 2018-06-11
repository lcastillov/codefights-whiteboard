# Two Teams

Given an array of integers, calculate the difference between elements in even positions and elements in odd positions.

`[1, 11, 13, 6, 14] => (1 + 13 + 14) - (11 + 6) = 11`

## Solution

We can slice the array to get all its even/odd-positionated elements.

```python
>>> a = [1, 11, 13, 6, 14]
>>> a[::2]
[1, 13, 14]
>>> a[1::2]
[11, 6]
>>> sum(a[::2]) - sum(a[1::2])
11
```
