# Permutation Chiper

## Solution

`password.translate(table)` [[1]][translate] will map every character from password to another character according to `table`. The table must be an object that implements indexing via `__getitem__()`.

`str` has a `maketrans` [[2]][maketrans] method that help us in creating the `table`: `str.maketrans("abcdefghijklmnopqrstuvwxyz", key)`. The `maketrans` method is going to create a dictionary from unicode ordinals to unicode ordinals given two strings of the same length. Then `translate` will use `table` to create a new string from `pasword`, but using `table.get(ord(password[i]), ord(password[i])` for position `i` instead.

For instance:

```python
>>> str.maketrans('abc', 'xyz')
{97: 120, 98: 121, 99: 122}
```

We can use `string.ascii_lowercase` instead of `"abcdefghijklmnopqrstuvwxyz"`.

[translate]:https://docs.python.org/3/library/stdtypes.html#str.translate
[maketrans]:https://docs.python.org/3/library/stdtypes.html#str.maketrans