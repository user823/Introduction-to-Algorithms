package main 

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println(RandomSelect([]int{3, 2, 9, 0, 7, 5, 4, 8, 6, 1}, 0, 9, 3))
}

func RandomSelect(A []int, p int, r int, i int) int {
    if p >= r {
        return A[p]
    }
    q := RandomPartition(A, p, r)
    l := q - p + 1
    if l == i {
        return A[q]
    } else if l < i {
        return RandomSelect(A, q+1, r, i)
    } else {
        return RandomSelect(A, p, q-1, i-l)
    }
}

func RandomPartition(A []int, p int, q int) int {
    index := rand.Intn(q - p + 1) + p
    metric := A[index]
    A[p], A[index] = A[index], A[p]
    for p < q {
        for A[q] >= metric && p < q{
            q--
        }
        A[p] = A[q]
        for A[p] <= metric && p < q {
            p++
        }
        A[q] = A[p]
    }
    A[p] = metric
    return p
}
