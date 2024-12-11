Turns out that this blink ramps up in a hurry.   The line of stones exceeds 100 million stones in a meer 40ish blinks and just keeps scaling out of control.
It also turns out that keeping the entire list of stones in memory becomes impossible.

The solution?   Don't.

To mitigate memory issues, just run the process 5 blinks on the first stone, then iterate through the output a stone at a time for another 5 blinks each, and so forth.    This allows us to eventually work through all the stones, though it will take a while.

Originally part a just kept the entire list in memory and was fine.   Converting Part a to use this 5 at a time process, it takes 28,124 runs of the five blink process.   Part b may approach over 1 trillion executions of the five blink - TBD