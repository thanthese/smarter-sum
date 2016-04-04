Extremely simple utility app that ignores any garbage and sums all numbers from `STDIN`. For example:

    $ echo "1 2 3 hey that's not a number 4 5 6" | smarter-sum
    21

It's smart about maintaining precision.

    $ echo '12.43 100 2200' | smarter-sum
    2312.43

It'll keep your commas for you if you use them.

    $ echo '1,234 5,278, 1,000,000.123' | smarter-sum
    1,006,512.123

Negatives are no big thing.

    $ echo '100 -10' | smarter-sum
    90

And it will play nicely with dollar signs.

    $ echo '10.1 -$50' | smarter-sum
    -$39.90

The golden rule is that numbers must be whitespace delimited. Everything that's not a number is reduced to whitespace and the remainders are added together.

### Motivation

I find myself wanting to add numbers together in vim *a lot*. This utility makes that as easy as selecting a range and `:!smarter-sum`.

I wrote the original [smart_sum](https://github.com/thanthese/smart_sum) in python, but it broke when I updated to python 3. That was annoying, so here we are.

### Installation

    $ go get https://github.com/thanthese/smarter-sum
