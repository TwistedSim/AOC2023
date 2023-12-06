package util

import "fmt"

type Number interface {
    int64 | float64 | int
}

func Max[T Number](x, y T) T {
	if x < y {
		return y
	}
	return x
}

func Min[T Number](x, y T) T {
	if x > y {
		return y
	}
	return x
}

func Intersection[T comparable](pS ...[]T) (result []T) {
    hash := make(map[T]*int)
    for _, slice := range pS {
        duplicationHash := make(map[T]bool)
        for _, value := range slice {
            if _, isDup := duplicationHash[value]; !isDup { 
                if counter := hash[value]; counter != nil { 
                    if *counter++; *counter >= len(pS) {
                        result = append(result, value)
                    }
                } else {
                    i := 1
                    hash[value] = &i
                }
                duplicationHash[value] = true
            }
        }
    }
    return
}

func MinMax(array []int) (int, int) {
    var max int = array[0]
    var min int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return min, max
}

func Transpose(a, b []int) ([][]int, error) {

    if len(a) != len(b) {
        return nil, fmt.Errorf("zip: arguments must be of same length")
    }

    r := make([][]int, len(a))

    for i, e := range a {
        r[i] = []int{e, b[i]}
    }

    return r, nil
}