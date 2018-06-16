# Construct Shell

A 2D list lst of size `2 * n - 1` is called a shell if it is filled with zeros and has the following configuration:

- `lst[0]` contains one element;
- `lst[1]` contains two elements;
- ...
- `lst[n - 2]` contains `n - 1` elements;
- `lst[n - 1]` contains `n` elements;
- `lst[n]` contains `n - 1` elements;
- ...
- `lst[2 * n - 3]` contains two elements;
- `lst[2 * n - 2]` contains one element.

Given an integer `n`, return a shell list.

## Solution

We can construct the lenghts for every list in the shell using `range` [[1]][range] method in two parts:

```python
list(range(1, n)) + list(range(n, 0, -1)
```

Then we can solve the problem:

```python
def constructShell(n):
    return [[0] * l for l in list(range(1, n)) + list(range(n, 0, -1))]
```

### A shorter solution

Let `n` be equal to `3` consider the lenghts of the shell and the range `[-n+1, +n-1]`:

<pre>
[ 1,  2, 3, 2, 1]  # lengths
[-2, -1, 0, 1, 2]  # (-n+1) .. (+n-1)
</pre>

We can convert every value `i` in the range `[-n+1, +n-1]` into the corresponding length using the formula `n - abs(i)`. With this in mind we can get an even shorter solution for the problem:

```python
def constructShell(n):
    return [[0] * (n-abs(i)) for i in range(-n+1, n)]
```


[range]:https://docs.python.org/3/library/functions.html#func-range