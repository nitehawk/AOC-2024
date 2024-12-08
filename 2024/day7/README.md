Part 1:
With 2 operators to try in each space to make up the value, first thought is to just try all options and keep them in a slice as it goes.
Using: 292: 11 6 16 20 sample:
  Step 1:   11+6 => 17   11*6 => 66
  Step 2:   17+16 => 33  17*16 => 272  66+16 => 82  66*16 => 1056
  ...
Small optimization could be to ignore any value greater than the target for further operations
For each step, create new list with one operator first, then apply second operator in place and append the new list