# Get Commit

Remove `'0'`, `'?'`, `'+'`, and `'!'` from the beginning of a string.

## Solution

Use `lstrip` [[1]][lstrip] method: `s.lstrip("0?+!")`.

[lstrip]:https://docs.python.org/3/library/stdtypes.html#str.lstrip