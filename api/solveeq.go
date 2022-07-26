package api

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

//Error handling
type myerr struct {
	origin  string
	message error
}

func (me *myerr) Error() string {
	return fmt.Sprintf("api.%s()|%s", me.origin, me.message)
}

func Getresult(amx, bmx []float64, px, py, pz int) (float64, float64, float64, error) {
	orig := "Getresult"

	adj, err := solvematrixeq(amx, bmx)
	if err != nil {
		return -1, -1, -1, &myerr{
			origin:  orig,
			message: err,
		}
	} else {
		a := adj[0]
		b := adj[1]
		c := adj[2]
		d := adj[3]
		e := adj[4]

		fmt.Printf("Relation of sequence: Result = [%.6f(Index^4)]+[%.6f(Index^3)]+[%.6f(Index^2)]+[%.6f(Index)]+[%.6f]\n", a, b, c, d, e)
		return finddotfromgraph(adj, px),
			finddotfromgraph(adj, py),
			finddotfromgraph(adj, pz),
			nil
	}
}

//Solving 5 variable equation x = (a^-1)b by Cramer's rule
func solvematrixeq(a, b []float64) ([]float64, error) {
	orig := "Solvematrixeq"

	result := []float64{}

	//Create matrix (assume dimension of a is 5x5 and b is 5x1)
	amatrix := mat.NewDense(5, 5, a)
	bmatrix := mat.NewDense(5, 1, b)

	var x mat.Dense
	err := x.Solve(amatrix, bmatrix)
	if err != nil {
		return result, &myerr{
			origin:  orig,
			message: err,
		}
	}

	sol1 := mat.Row(nil, 0, &x)[0]
	sol2 := mat.Row(nil, 1, &x)[0]
	sol3 := mat.Row(nil, 2, &x)[0]
	sol4 := mat.Row(nil, 3, &x)[0]
	sol5 := mat.Row(nil, 4, &x)[0]

	result = []float64{sol1, sol2, sol3, sol4, sol5}
	return result, nil
}

//Find Y location on graph Y = A(X^4)+B(X^3)+C(X^2)+D(X)+E eqation
func finddotfromgraph(adj []float64, post int) float64 {

	y := adj[0]*(math.Pow(float64(post), 4)) +
		adj[1]*(math.Pow(float64(post), 3)) +
		adj[2]*(math.Pow(float64(post), 2)) +
		adj[3]*(math.Pow(float64(post), 1)) +
		adj[4]

	return y
}
