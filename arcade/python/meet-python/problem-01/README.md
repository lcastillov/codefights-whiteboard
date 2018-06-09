# Collections Truthness

What will be the value of `res` after the following snippet is executed:

```python
xs = [()]
res = [False] * 2
if xs:
    res[0] = True
if xs[0]:
    res[1] = True
```

## Solution

This quiz expects from you to know the implicit conversion from list to boolean. A list/tuple will evaluate to `False` only if it's empty.

In the snippet, `xs` is a non-empty list and `xs[0]` is an empty tuple. So `res` will be `[True, False]` at the end.
