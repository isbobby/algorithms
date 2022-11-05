# Matrix Multiplication
Multiplying a $M \times N$ matrix `A` and $N \times P$ matrix `B` gives a $M \times P$ matrix `C`. Where $C^{i}_{j}$ is obtained by the dot product of $i$ th row in `A` and $j$ th column in `B`.

$$C^{i}_{j} = \sum_{k=1}^{n} A_{ik} \times B_{kj}$$

## Simple Multiplication
The above sequence is easily translated into a triply nested for loop. And the complexity of this operation will be $ \Omega(N^3)$ (lower bound is $N^3$).

## Divide and Conquer
We can also apply divide and conquer philosophy in performing matrix multiplication. For started, we have a $N \times N$ matrix where $N$ is a power of 2.

We can recursively divide the matrix by breaking them into four quadrants by $N \div 2$. We will end up with four quadrants. We then compute 