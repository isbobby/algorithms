# Matrix Multiplication
Multiplying a $M \times N$ matrix `A` and $N \times P$ matrix `B` gives a $M \times P$ matrix `C`. Where $C^{i}_{j}$ is obtained by the dot product of $i$ th row in `A` and $j$ th column in `B`.

$$C^{i}_{j} = \sum_{k=1}^{n} A_{ik} \times B_{kj}$$

## Simple Multiplication
The above sequence is easily translated into a triply nested for loop. And the complexity of this operation will be $ \Omega(N^3)$ (lower bound is $N^3$).

## Divide and Conquer
We can also apply divide and conquer philosophy in performing matrix multiplication.

The division step will divide a matrix C into 4 different quadrants - $\{ C_{11}, C_{12}, C_{21}, C_{22} \}$.

Likewise, the input matrix A and B are also divided into the following
$$\{ A_{11}, A_{12}, A_{21}, A_{22}, B_{11}, B_{12}, B_{21}, B_{22} \}$$

With this, we define the recurrence cases to be

1. $C_{11} = F(A_{11}, B_{11}) + F(A_{12}, B_{21})$
2. $C_{12} = F(A_{11}, B_{12}) + F(A_{12}, B_{22})$
3. $C_{21} = F(A_{21}, B_{11}) + F(A_{22}, B_{21})$
4. $C_{22} = F(A_{21}, B_{12}) + F(A_{22}, B_{22})$

And the base case can be defined as
```
n = A.rows
if n == 1
	c11 = a11 * b11 // just a scalar multiplication for [1]x[1]
	return c11 
```

The partial results will be stored in $C_{ij}$ and returned to the call from previous layer. Eventually, $C$ will be constructed and returned.

We can evaluate the time complexity by assigning time complexity to each step. First, we assume the partitioning of the matrix takes constant time. Then, we reach the base case check where we perform a single multiplication, this will take constant time as well.

Now, we will partition (constant time) and make 8 different recursive calls. Each recursive call will multiplies two $n/2$ matrices, which contributes $T(n/2)$ to the overall running time. The time taken for these calls will be $8T(n/2)$.

After the recursive calls return, we will then sum them up. Each of the matrix returned from the recursive call will contain $n^2/4$ entries, and for four additions it will take $n^2$ time.

```
def matrix_multiplication(A, B):
	n = A.rows
	c = matrix.New(n,n) # C is an N by N matrix
	if n == 1: # O(1)
		return A * B # 1 by 1 multiplication 

	else:
		A11,A12,A21,A22 = partition(A)
		B11,B12,B21,B22 = partition(B)
		# partition takes O(1) time

		# each call deals with 2 n/2 matrix, so 8(n/2)
		# the result of a matrix multiplication call is n/4 matrix
		# matrix addition of two n/4 matrix is n^2 / 4 time 
		C11 = matrix_multiplication(A11, B11) + matrix_multiplication(A12, B21)
		C12 = matrix_multiplication(A11, B12) + matrix_multiplication(A12, B22)
		C21 = matrix_multiplication(A21, B11) + matrix_multiplication(A22, B21)
		C22 = matrix_multiplication(A21, B12) + matrix_multiplication(A22, B22)

	C = combine(C11,C12,C21,C22)
	return C
```

Combining everything gives us the recurrence relation of (for $n>1$)
$$
\begin{equation}
\begin{aligned}
T(n) = &\ \Theta(1)+8(T/2) +\Theta(n^2) \\
= &\ 8T(n/2) + \Theta(n^2)
\end{aligned}
\end{equation}
$$

According to master method, this algorithm also results in $\Theta(n^3)$. This is not faster than the naive method.