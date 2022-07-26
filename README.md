# NexterExam1

The concept to solve sequence problem is create a graph equation that pass 5 already known points 
(more than 1 graph equation can do) so the other points will be found.

For example:
if data set is [1, X, 8, 17, Y, Z, 78, 113] (they can be shuffled).

We already know 5 points (x,y) are
(1,1)
(3,8)
(4,17)
(7,78)
(8,113)

when x is index of sequence
     y is value of sequence
     
Then graph equation: Y = A(X^4) + B(X^3) + C(X^2) + D(X) + E 
when A, B, C, D, E are constant will be used for finding sequence relation.
Next, we will get a 5 equation with 5 parameter.

1   = A(1^4) + B(1^3) + C(1^2) + D(1) + E
8   = A(3^4) + B(3^3) + C(3^2) + D(3) + E
17  = A(4^4) + B(4^3) + C(4^2) + D(4) + E
78  = A(7^4) + B(7^3) + C(7^2) + D(7) + E
113 = A(8^4) + B(8^3) + C(8^2) + D(8) + E

Next, matrix will be used for solving these equation -> AX = B 
when A,B,X is matrix ("Gonum" library will be used for solving matrix equation).
Therefore we will get the value of A, B, C, D, E

find X value by passing x=2 to graph equation
find Y value by passing x=5 to graph equation
find Z value by passing x=6 to graph equation
