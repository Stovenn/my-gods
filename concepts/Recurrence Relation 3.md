# Recurrence Relation (T(n) = T(n-1)+1) #3

## Recursive algorithm

```go
func Test(n int) { //T(n)
    if n > 0 { // 1

        Test(n-1); //T(n-1)
        Test(n-1); //T(n-1)
    }
}

```

So that

```text
T(n) = 2T(n-1) + 1
```

## Solve T(n) with recursion tree

```text
                T(n)
            /     |     \
          1   T(n-1)      T(n-1)
          /    |    \       /   |   \
        1  T(n-2)  T(n-2) 1 T(n-2)  T(n-2)
        /   |   \  /  |  \  /  |  \ /  |  \
    1  T(n-3) T(n-3) same    same     same        
       
```
Each calls increase by the power of 2:  

So we have :
$$ 1+2^1+2^2+2^3+... +2^k = 2^{k+1}-1$$
We assume that k - n = 0  
so k = n  
We can then write  
$$ T(n)= 2^{n+1}-1$$

We can conclude that the time taken by the algorithm is: 
$$ \theta(2^n) $$
