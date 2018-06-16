def wordPower(word):
    num = {c: ord(c)-96 for c in "abcdefghijklmnopqrstuvwxyz"}
    return sum([num[ch] for ch in word])
