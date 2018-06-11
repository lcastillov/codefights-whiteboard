# Cat Walk

Replace all multiple space characters in a given string with single ones.

For `line = "def      m   e  gaDifficu     ltFun        ction(x):"`, the output should be `catWalk(line) = "def m e gaDifficu ltFun ction(x):"`.

## Solution

We can use `split` [[1]][split] and `join` [[2]][join] to solve this problem. The solution would looks like `" ".join(line.split())`. We should avoid passing a blank space to `split`, if we do this, then `line.split(" ")` will return empty strings for every two consecutive spaces in `line`.

Another way to solve this problem is using method `sub` [[3]][sub] in the `re` library: `re.sub("\w+", " ", line).strip()`. The previous statement will replace consecutive spaces with a single space. Note if the string starts or ends with spaces we must to remove them using `strip` [[4]][strip].

[split]:https://docs.python.org/3/library/stdtypes.html#str.split
[join]:https://docs.python.org/3/library/stdtypes.html#str.join
[sub]:https://docs.python.org/3/library/re.html#re.sub
[strip]:https://docs.python.org/3/library/stdtypes.html#str.strip
