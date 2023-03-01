# Recurrence Relation (T(n) = T(n-1)+1)

## Recursive algorithm
```
void Test(int n) {
    if (n>0) {
        printf("\d",n);
        Test(n-1);
    }
}

Test(3)
```
For n = 3 the program will call Test function N+1 time and Print a value N times
so f(n) = n+1 and O(n) = n

## Recurrence relation
The standard function name for recurrence relations is T(n)  
Coming back to the previous algorithm, let's consider T(n) = Test(int n), we print a value (1 unit of time) and call Test(n-1) so that
```
T(n) = T(n-1) + 1
```

| N   | T(n)    |
| --- | --------|
| = 0 | 1       |
| > 0 | T(n-1)+1|


## Substitution method

To solve T(n) = T(n-1)+1 we can use the subsitution method by finding T(n-1).  
So that
```  
T(n) = T(n-1)+1
```
becomes
```
T(n-1) = T(n-2)+1
```
T(n) can now be written
```
[T(n-2) +1]+1 == T(n-2)+2
```
We continue to substitute for k times so that we have T(n) = T(n-k)+k

Assuming n-k = 0 then n = k
```
T(n) = T(n-n) +n
T(n) = T(0) +n
T(n) = 1+n
```
order of θ(n)