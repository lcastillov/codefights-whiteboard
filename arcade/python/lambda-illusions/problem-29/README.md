# Is Test Solvable

## Solve

We need to create a lambda expression that given an integer `n` return the sum of its digits.

We can convert `n` to string, then `map` every character to an integer and finally, `sum` all the digits.

```python
lambda n: sum(map(int, str(n)))
```

Lets see this step by step:

```python
>>> n = 1992
>>> str(n)  # convert n to string
'1992'
>>> list( map(int, str(n)) )  # maps every character to an integer
[1, 9, 9, 2]
>>> sum(map(int, str(n)))  # sum all the digits (ignore list conversion here)
21
```