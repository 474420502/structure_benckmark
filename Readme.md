# Benchmark

- replace ns/op|B/op|allocs/op
- dup means Duplicate 

## AVL

| scene                     | count   | ns/op | B/op | allocs/op |
|:--------------------------|:-------:|:-----:|:---:|:--:|
| BenchmarkRemove/10-8      | 1000000 | 83.6  | 8   | 1  |
| BenchmarkRemove/100-8     | 1000000 | 96.6  | 8   | 1  |
| BenchmarkRemove/1000-8    | 1000000 |  172  | 8   | 1  |
| BenchmarkRemove/10000-8   | 1000000 |  236  | 8   | 1  |
| BenchmarkRemove/100000-8  | 1000000 | 343   | 8   | 1  |
| BenchmarkRemove/1000000-8 | 5000000 | 1058  | 8   | 1  |
| BenchmarkGet/10-8         | 1000000 | 60.6  | 8   | 1  |
| BenchmarkGet/100-8        | 1000000 | 81.7  | 8   | 1  |
| BenchmarkGet/1000-8       | 1000000 | 158   | 8   | 1  |
| BenchmarkGet/10000-8      | 1000000 | 234   | 8   | 1  |
| BenchmarkGet/100000-8     | 1000000 | 355   | 8   | 1  |
| BenchmarkGet/1000000-8    | 5000000 | 1132  | 8   | 1  |
| BenchmarkPut/10-8         | 1000000 | 89.1  | 56  | 2  |
| BenchmarkPut/100-8        | 1000000 | 92.4  | 56  | 2  |
| BenchmarkPut/1000-8       | 1000000 | 142   | 56  | 2  |
| BenchmarkPut/10000-8      | 1000000 | 202   | 56  | 2  |
| BenchmarkPut/100000-8     | 1000000 | 313   | 56  | 2  |
| BenchmarkPut/1000000-8    | 5000000 | 888   | 56  | 2  |

## AVLDUP

| scene                     | count   | ns/op | B/op | allocs/op |
|:--------------------------|:-------:|:-----:|:---:|:--:|
| BenchmarkRemove/10-8      | 1000000 | 50.4  | 8   | 1  |
| BenchmarkRemove/100-8     | 1000000 | 55.0  | 8   | 1  |
| BenchmarkRemove/1000-8    | 1000000 |  114  | 8   | 1  |
| BenchmarkRemove/10000-8   | 1000000 |  171  | 8   | 1  |
| BenchmarkRemove/100000-8  | 1000000 | 269   | 8   | 1  |
| BenchmarkRemove/1000000-8 | 5000000 | 884   | 8   | 1  |
| BenchmarkGet/10-8         | 1000000 | 25.7  | 8   | 1  |
| BenchmarkGet/100-8        | 1000000 | 39.0  | 8   | 1  |
| BenchmarkGet/1000-8       | 1000000 | 96.3  | 8   | 1  |
| BenchmarkGet/10000-8      | 1000000 | 153   | 8   | 1  |
| BenchmarkGet/100000-8     | 1000000 | 250   | 8   | 1  |
| BenchmarkGet/1000000-8    | 5000000 | 848   | 8   | 1  |
| BenchmarkPut/10-8         | 1000000 | 89.7  | 56  | 2  |
| BenchmarkPut/100-8        | 1000000 | 96.3  | 56  | 2  |
| BenchmarkPut/1000-8       | 1000000 | 154   | 56  | 2  |
| BenchmarkPut/10000-8      | 1000000 | 211   | 56  | 2  |
| BenchmarkPut/100000-8     | 1000000 | 321   | 56  | 2  |
| BenchmarkPut/1000000-8    | 5000000 | 905   | 56  | 2  |

## AVLKEY

| scene                     | count   | ns/op | B/op | allocs/op |
|:--------------------------|:-------:|:-----:|:---:|:--:|
| BenchmarkRemove/10-8      | 1000000 | 82.5  | 8   | 1  |
| BenchmarkRemove/100-8     | 1000000 | 95.0  | 8   | 1  |
| BenchmarkRemove/1000-8    | 1000000 |  178  | 8   | 1  |
| BenchmarkRemove/10000-8   | 1000000 |  243  | 8   | 1  |
| BenchmarkRemove/100000-8  | 1000000 | 389   | 8   | 1  |
| BenchmarkRemove/1000000-8 | 5000000 | 1154  | 8   | 1  |
| BenchmarkGet/10-8         | 1000000 | 60.7  | 8   | 1  |
| BenchmarkGet/100-8        | 1000000 | 84.3  | 8   | 1  |
| BenchmarkGet/1000-8       | 1000000 | 173   | 8   | 1  |
| BenchmarkGet/10000-8      | 1000000 | 237   | 8   | 1  |
| BenchmarkGet/100000-8     | 1000000 | 405   | 8   | 1  |
| BenchmarkGet/1000000-8    | 5000000 | 1221  | 8   | 1  |
| BenchmarkPut/10-8         | 1000000 | 109   | 80  | 3  |
| BenchmarkPut/100-8        | 1000000 | 105   | 80  | 3  |
| BenchmarkPut/1000-8       | 1000000 | 155   | 80  | 3  |
| BenchmarkPut/10000-8      | 1000000 | 226   | 80  | 3  |
| BenchmarkPut/100000-8     | 1000000 | 350   | 80  | 3  |
| BenchmarkPut/1000000-8    | 5000000 | 997   | 80  | 3  |

## AVLKEYDUP

| scene                     | count   | ns/op | B/op | allocs/op |
|:--------------------------|:-------:|:-----:|:---:|:--:|
| BenchmarkRemove/10-8      | 1000000 | 49.4  | 8   | 1  |
| BenchmarkRemove/100-8     | 1000000 | 52.4  | 8   | 1  |
| BenchmarkRemove/1000-8    | 1000000 |  115  | 8   | 1  |
| BenchmarkRemove/10000-8   | 1000000 |  175  | 8   | 1  |
| BenchmarkRemove/100000-8  | 1000000 | 313   | 8   | 1  |
| BenchmarkRemove/1000000-8 | 5000000 | 966   | 8   | 1  |
| BenchmarkGet/10-8         | 1000000 | 26.0  | 8   | 1  |
| BenchmarkGet/100-8        | 1000000 | 39.3  | 8   | 1  |
| BenchmarkGet/1000-8       | 1000000 | 98.1  | 8   | 1  |
| BenchmarkGet/10000-8      | 1000000 | 161   | 8   | 1  |
| BenchmarkGet/100000-8     | 1000000 | 287   | 8   | 1  |
| BenchmarkGet/1000000-8    | 5000000 | 962   | 8   | 1  |
| BenchmarkPut/10-8         | 1000000 | 104   | 80  | 3  |
| BenchmarkPut/100-8        | 1000000 | 103   | 80  | 3  |
| BenchmarkPut/1000-8       | 1000000 | 159   | 80  | 3  |
| BenchmarkPut/10000-8      | 1000000 | 229   | 80  | 3  |
| BenchmarkPut/100000-8     | 1000000 | 365   | 80  | 3  |
| BenchmarkPut/1000000-8    | 5000000 | 1042  | 80  | 3  |

## SkipList

- change the label of Set to Put

| scene                     | count   | ns/op | B/op | allocs/op |
|:--------------------------|:-------:|:-----:|:---:|:--:|
| BenchmarkRemove/10-8      | 1000000 |  157  | 8   | 1  |
| BenchmarkRemove/100-8     | 1000000 |  156  | 8   | 1  |
| BenchmarkRemove/1000-8    | 1000000 |  184  | 8   | 1  |
| BenchmarkRemove/10000-8   | 1000000 |  298  | 8   | 1  |
| BenchmarkRemove/100000-8  | 1000000 | 610   | 8   | 1  |
| BenchmarkRemove/1000000-8 | 5000000 | 2140  | 8   | 1  |
| BenchmarkGet/10-8         | 1000000 | 97.1  | 8   | 1  |
| BenchmarkGet/100-8        | 1000000 | 113   | 8   | 1  |
| BenchmarkGet/1000-8       | 1000000 | 176   | 8   | 1  |
| BenchmarkGet/10000-8      | 1000000 | 351   | 8   | 1  |
| BenchmarkGet/100000-8     | 1000000 | 902   | 8   | 1  |
| BenchmarkGet/1000000-8    | 5000000 | 2852  | 8   | 1  |
| BenchmarkPut/10-8         | 1000000 | 311   | 91  | 4  |
| BenchmarkPut/100-8        | 1000000 | 268   | 91  | 4  |
| BenchmarkPut/1000-8       | 1000000 | 294   | 91  | 4  |
| BenchmarkPut/10000-8      | 1000000 | 435   | 91  | 4  |
| BenchmarkPut/100000-8     | 1000000 | 1020  | 91  | 4  |
| BenchmarkPut/1000000-8    | 5000000 | 2561  | 91  | 4  |

## VBT

- the tree is designed by me
- it is compared with skiplist
- it is for priority queue
- Index, Range
- i think SBT is better than it!

| scene                     | count   | ns/op | B/op | allocs/op |
|:--------------------------|:-------:|:-----:|:---:|:--:|
| BenchmarkRemove/10-8      | 1000000 | 90.3  | 8   | 1  |
| BenchmarkRemove/100-8     | 1000000 |  111  | 8   | 1  |
| BenchmarkRemove/1000-8    | 1000000 |  197  | 8   | 1  |
| BenchmarkRemove/10000-8   | 1000000 |  271  | 8   | 1  |
| BenchmarkRemove/100000-8  | 1000000 | 434   | 8   | 1  |
| BenchmarkRemove/1000000-8 | 5000000 | 1230  | 8   | 1  |
| BenchmarkGet/10-8         | 1000000 | 65.5  | 8   | 1  |
| BenchmarkGet/100-8        | 1000000 | 99.7  | 8   | 1  |
| BenchmarkGet/1000-8       | 1000000 | 169   | 8   | 1  |
| BenchmarkGet/10000-8      | 1000000 | 238   | 8   | 1  |
| BenchmarkGet/100000-8     | 1000000 | 404   | 8   | 1  |
| BenchmarkGet/1000000-8    | 5000000 | 1226  | 8   | 1  |
| BenchmarkPut/10-8         | 1000000 | 89.2  | 56  | 2  |
| BenchmarkPut/100-8        | 1000000 | 92.4  | 56  | 2  |
| BenchmarkPut/1000-8       | 1000000 | 150   | 56  | 2  |
| BenchmarkPut/10000-8      | 1000000 | 211   | 56  | 2  |
| BenchmarkPut/100000-8     | 1000000 | 326   | 56  | 2  |
| BenchmarkPut/1000000-8    | 5000000 | 887   | 56  | 2  |

## VBTKEY

| scene                     | count   | ns/op | B/op | allocs/op |
|:--------------------------|:-------:|:-----:|:---:|:--:|
| BenchmarkRemove/10-8      | 1000000 | 84.1  | 8   | 1  |
| BenchmarkRemove/100-8     | 1000000 |  105  | 8   | 1  |
| BenchmarkRemove/1000-8    | 1000000 |  191  | 8   | 1  |
| BenchmarkRemove/10000-8   | 1000000 |  261  | 8   | 1  |
| BenchmarkRemove/100000-8  | 1000000 | 410   | 8   | 1  |
| BenchmarkRemove/1000000-8 | 5000000 | 1194  | 8   | 1  |
| BenchmarkGet/10-8         | 1000000 | 62.2  | 8   | 1  |
| BenchmarkGet/100-8        | 1000000 | 85.4  | 8   | 1  |
| BenchmarkGet/1000-8       | 1000000 | 162   | 8   | 1  |
| BenchmarkGet/10000-8      | 1000000 | 238   | 8   | 1  |
| BenchmarkGet/100000-8     | 1000000 | 399   | 8   | 1  |
| BenchmarkGet/1000000-8    | 5000000 | 1224  | 8   | 1  |
| BenchmarkPut/10-8         | 1000000 | 101   | 80  | 3  |
| BenchmarkPut/100-8        | 1000000 | 103   | 80  | 3  |
| BenchmarkPut/1000-8       | 1000000 | 152   | 80  | 3  |
| BenchmarkPut/10000-8      | 1000000 | 222   | 80  | 3  |
| BenchmarkPut/100000-8     | 1000000 | 326   | 80  | 3  |
| BenchmarkPut/1000000-8    | 5000000 | 933   | 80  | 3  |

## PriorityQueue

- it is compared with skiplist

| scene                     | count   | ns/op | B/op | allocs/op |
|:--------------------------|:-------:|:-----:|:---:|:--:|
| BenchmarkRemove/10-8      | 1000000 | 93.1  | 8   | 1  |
| BenchmarkRemove/100-8     | 1000000 |  109  | 8   | 1  |
| BenchmarkRemove/1000-8    | 1000000 |  196  | 8   | 1  |
| BenchmarkRemove/10000-8   | 1000000 |  269  | 8   | 1  |
| BenchmarkRemove/100000-8  | 1000000 | 401   | 8   | 1  |
| BenchmarkRemove/1000000-8 | 5000000 | 1187  | 8   | 1  |
| BenchmarkGet/10-8         | 1000000 | 62.8  | 8   | 1  |
| BenchmarkGet/100-8        | 1000000 | 84.3  | 8   | 1  |
| BenchmarkGet/1000-8       | 1000000 | 165   | 8   | 1  |
| BenchmarkGet/10000-8      | 1000000 | 235   | 8   | 1  |
| BenchmarkGet/100000-8     | 1000000 | 341   | 8   | 1  |
| BenchmarkGet/1000000-8    | 5000000 | 1189  | 8   | 1  |
| BenchmarkPut/10-8         | 1000000 | 90.8  | 56  | 2  |
| BenchmarkPut/100-8        | 1000000 | 89.7  | 56  | 2  |
| BenchmarkPut/1000-8       | 1000000 | 146   | 56  | 2  |
| BenchmarkPut/10000-8      | 1000000 | 201   | 56  | 2  |
| BenchmarkPut/100000-8     | 1000000 | 300   | 56  | 2  |
| BenchmarkPut/1000000-8    | 5000000 | 854   | 56  | 2  |

## PriorityQueueKey

- it is compared with skiplist

| scene                     | count   | ns/op | B/op | allocs/op |
|:--------------------------|:-------:|:-----:|:---:|:--:|
| BenchmarkRemove/10-8      | 1000000 | 93.7  | 8   | 1  |
| BenchmarkRemove/100-8     | 1000000 |  108  | 8   | 1  |
| BenchmarkRemove/1000-8    | 1000000 |  198  | 8   | 1  |
| BenchmarkRemove/10000-8   | 1000000 |  276  | 8   | 1  |
| BenchmarkRemove/100000-8  | 1000000 | 483   | 8   | 1  |
| BenchmarkRemove/1000000-8 | 5000000 | 1271  | 8   | 1  |
| BenchmarkGet/10-8         | 1000000 | 62.9  | 8   | 1  |
| BenchmarkGet/100-8        | 1000000 | 86.0  | 8   | 1  |
| BenchmarkGet/1000-8       | 1000000 | 161   | 8   | 1  |
| BenchmarkGet/10000-8      | 1000000 | 239   | 8   | 1  |
| BenchmarkGet/100000-8     | 1000000 | 406   | 8   | 1  |
| BenchmarkGet/1000000-8    | 5000000 | 1252  | 8   | 1  |
| BenchmarkPut/10-8         | 1000000 | 112   | 80  | 3  |
| BenchmarkPut/100-8        | 1000000 | 109   | 80  | 3  |
| BenchmarkPut/1000-8       | 1000000 | 157   | 80  | 3  |
| BenchmarkPut/10000-8      | 1000000 | 221   | 80  | 3  |
| BenchmarkPut/100000-8     | 1000000 | 334   | 80  | 3  |
| BenchmarkPut/1000000-8    | 5000000 | 928   | 80  | 3  |