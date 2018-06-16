# Multiplication Table

Construct a multiplication table. Given an integer `n`, return the multiplication table of size `n × n`.

For `n = 5`, the output should be

<pre>
multiplicationTable(n) = [[1, 2,  3,  4,  5 ], 
                          [2, 4,  6,  8,  10], 
                          [3, 6,  9,  12, 15], 
                          [4, 8,  12, 16, 20], 
                          [5, 10, 15, 20, 25]]
</pre>

## Solution

Using _list comprehensions_ iterate for every row `i` in the range `[1, n]` and create a new array for every `i` with the correct products:

```python
>>> n = 3
>>> [[i*j for j in range(1, n + 1)] for i in range(1, n+1)]
[[1, 2, 3], [2, 4, 6], [3, 6, 9]]
```
