# Efficient Comparison

You would like to write a function that takes integer numbers `x`, `y`, `L` and `R` as parameters and returns `True` if `x^y` lies in the interval `(L, R]` and `False` otherwise. You're considering several ways to write a conditional statement inside this function:

1. `if L < x**y <= R:`
2. `if x**y > L and x**y <= R:`
3. `if x**y in range(L + 1, R + 1):`

What option would be the most efficient in terms of execution time?

## Solution

The first option is the most efficient. The second option calculates `x**y` twice and the third option creates an unnecessary array from `L` to `R`.