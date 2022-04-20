# Prime

Simple module for testing generating prime numbers

## File Structure

``` console
├── generator   (Generate a list of primes)
│   ├── count   (Generate a given number of primes from 0)
│   └── domain  (Generate all primes in a given domain)
├── primality   (Check if an integer is prime)
├── tool        (Tools for reading, writing and managing lists of prime numbers)
```

## Primality > HybridCache Optimisation

For checking the large values the cache should be large enough so that it is inclusive of most numbers, larger numbers that rely on checking for possible prime factors are significantly slower. This is allowable for checking outliers but if a large number of numbers fall into this category performance is going to be limited. Avoid at all cost checking numbers which are greater than the largest number in the cache squared. These will revert to checking all possible factors greater than the larest number in the cache squared exhaustively and is thus very slow.

> Please note it is also optimal for the cache length to be equal to 2^n as the cache using a binary search.

Cache Size = 2^n

> Please note the above applies when this function is being called sequentially and thus the cache is in L3 or faster memory. Occasional calls have not been tested by are likely to lead to unpredictable behaviour and performance due to cache missed and the additional latency of loading from main memory
