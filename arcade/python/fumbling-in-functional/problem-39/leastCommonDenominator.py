from fractions import gcd
from functools import reduce

def leastCommonDenominator(denominators):
    return reduce(lambda n, m: n*m/gcd(n,m), denominators, 1)
