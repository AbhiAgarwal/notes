#### Strength reduction

Strength reduction looks for expressions involving a loop invariant and an induction variable. Some of those 
expressions can be simplified. For example, the multiplication of loop invariant c and induction variable.

```
c = 8;
for (i = 0; i < N; i++)
{
    y[i] = c * i;
}
```

can be replaced with successive weaker additions

```
c = 8;
k = 0;
for (i = 0; i < N; i++)
{
    y[i] = k;
    k = k + c;
}
```
