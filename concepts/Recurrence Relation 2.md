# Recurrence Relation (T(n) = T(n-1)+n) #2

## Recursive algorithm

```go
func Test(n int) { //T(n)
    if n > 0 { // 1
        for i := 0; i < n; i++ { // n+1
           fmt.Printf("%d", n); //n
        }
        Test(n-1); //T(n-1)
    }
}

```

So that

```text
T(n) = T(n-1) + 2n + 2
```

We can simplify by taking the asymptotic notation of (2n + 2) so that T(n) becomes

```text
T(n) = T(n-1) + n (for n > 0)
```

## Solve T(n) with recursion tree

```text
            T(n)
            / \
           n   T(n-1)
                 / \
              n-1   T(n-1)
                     / \
                  n-2   T(n-1)
                         .
                         .
                         .
                        T(2)
                        / \
                       2   T(1)
                            / \
                           1   T(0)
```

From the tree we add the time taken by each call:

$$0+1+2+...+(n-1)+n $$
so that
$$T(n) = {n+(n+1)\over2}$$

We can conclude that the time taken by the algorithm is:  
$$ \theta(n^2) $$
