# Print List

Implement a function that, given a list lst, will return it as a string as follows: `"This is your list: lst"`.

For `lst = [1, 2, 3, 4, 5]`, the output should be `printList(lst) = "This is your list: [1, 2, 3, 4, 5]"`.

## Solution

The string representation of a list in python is exactly what we want:

```python
>>> a = [1, 2, 3, 4, 5]
>>> print ("This is your list: %s" % a)  # option 1
This is your list: [1, 2, 3, 4, 5]
>>> print ("This is your list: {}".format(a))  # option 2
This is your list: [1, 2, 3, 4, 5]
>>> print ("This is your list: " + str(a))  # option 3
This is your list: [1, 2, 3, 4, 5]
```