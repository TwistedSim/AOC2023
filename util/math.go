package util

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

func Transpose(array [][]int) [][]int {
    r := make([][]int, len(array[0]))
    for i := 0; i < len(r); i++ {
        r[i] = make([]int, len(array))
        
		for j := 0; j < len(r[i]); j++ {
			r[i][j] = array[j][i]
		}
        
	}
    return r
}

func Sum[T Number](list []T) (result T) {
	for _, value := range list {
		result += value
	}
	return
}


func GCD(a, b int) int {
	for b != 0 {
		b, a = a % b, b
	}
	return a
}

func LCM(a, b int, integers ...int) int {
    result := a * b / GCD(a, b)
    for i := 0; i < len(integers); i++ {
        result = LCM(result, integers[i])
    }
    return result
}