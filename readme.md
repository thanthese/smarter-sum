Extremely simple utility app that ignores any garbage and sums all numbers
from stdin. For example:

    $ echo "1 2 3 hey that's not a number 4 5 6" | smarter-sum
    21

It's smart about maintaining precision, too.

    $ echo '$12.43 $100 $2200' | smarter-sum
    2312.43

It will even keep your commas for you, should you decide to use them.

    $ echo '1,234 5,278, 1,000,000.123' | smarter-sum
    1,006,512.123

The golden rule is that numbers must be whitespace delimited. More specifically, `$`s are removed, `,` are left in numbers, and everything else is reduced to whitespace. Those numbers that remain are added together.

### Motivation

I find myself wanting to add numbers together in vim *a lot*. This utility
makes that as easy as selecting a range and `:!smarter-sum`.

### Installation

    $ go get https://github.com/thanthese/smarter-sum
