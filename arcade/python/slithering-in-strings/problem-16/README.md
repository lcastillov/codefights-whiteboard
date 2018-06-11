# Is Word Palindrome

Given a `word`, check whether it is a palindrome or not. A string is considered to be a palindrome if it reads the same in both directions.

## Solution

This is a "must-to-know-trick" in python, you can reverse a string using `word[::-1]`.

You can read a bit more about slicing in [this](https://stackoverflow.com/questions/509211/understanding-pythons-slice-notation) Stack Overflow question.

Another way to solve this problem is using `join` and `reversed` as follow: `word == "".join(reversed(word))`.