# Pascals Triangle

Compute Pascal's triangle up to a given number of rows.

In Pascal's Triangle each number is computed by adding the numbers to
the right and left of the current position in the previous row.

```plain
    1
   1 1
  1 2 1
 1 3 3 1
1 4 6 4 1
# ... etc
```

## Running the tests

To run the tests run the command `go test` from within the exercise directory.

If the test suite contains benchmarks, you can run these with the `-bench`
flag:

    go test -bench .

Keep in mind that each reviewer will run benchmarks on a different machine, with
different specs, so the results from these benchmark tests may vary.

## Further information

For more detailed information about the Go track, including how to get help if
you're having trouble, please visit the exercism.io [Go language page](http://exercism.io/languages/go/about).

## Source

Pascal's Triangle at Wolfram Math World [http://mathworld.wolfram.com/PascalsTriangle.html](http://mathworld.wolfram.com/PascalsTriangle.html)

## Submitting Incomplete Problems
It's possible to submit an incomplete solution so you can see how others have completed the exercise.

