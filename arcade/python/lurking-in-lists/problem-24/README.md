# Remove Tasks

Given the list of task ids in your `toDo` list, remove each `k`<sup>th</sup> task from it and return the list of remaining tasks.

For `k = 3` and `toDo = [1237, 2847, 27485, 2947, 1, 247, 374827, 22]`,
the output should be
`removeTasks(k, toDo) = [1237, 2847, 2947, 1, 374827, 22]`.

## Solution

Start with the k-th element and then remove every `k` elements.

```python
del toDo[k-1::k]
```
