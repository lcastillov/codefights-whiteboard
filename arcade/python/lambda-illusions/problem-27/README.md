# Get Points

## Solution

We need to create a lambda expression that takes 2 parameters `i` and `ans`, the index of the question and whether the question was correct or not. The lambda expression must return `i+1` (1-based) if the answer to the `i`<sup>th</sup> question was correct or `-p` otherwise.

You can use _Conditional Expression_ [[1]][conditional-expression] to solve this problem easily:

```python
questionPoints = lambda i, ans: (i+1) if ans else -p
```

[conditional-expression]:https://www.python.org/dev/peps/pep-0308/