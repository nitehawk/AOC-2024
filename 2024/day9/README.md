Represent disk after reading as a slice of ints representing each block
    -1 for free blocks

After reading, traverse the slice from both ends until meeting in the middle.
* Find the last block with a file
* Find the first free block
* Move block to first free space and mark last block free

Once compacted, walk through the array adding each product to the sum