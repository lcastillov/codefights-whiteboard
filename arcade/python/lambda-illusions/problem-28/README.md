# Sort Students

Sort students by their suernames.

## Solution

Use a lambda expression that maps a name to a surname.

- `Jacky Mon Simonoff -> Simonoff`
- `John Smith -> Smith`
- `Kate -> Kate`

```python
lambda name: name.split()[-1]
```

We can use this lambda expression as argument for the `key` parameter in the builtin `sort` [[1]][sort] function. As you can read in the previous link, `sort` is stable.

[sort]:https://docs.python.org/3/library/stdtypes.html#list.sort
