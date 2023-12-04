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