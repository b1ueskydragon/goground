package main

// 	"abab", "bbcc" is false,
//	"aabb", "bbcc" is true.
func isIsomorphic(s string, t string) bool {
	// set
	xRange := toIdenticalList(s, make(map[int32]struct{}))
	yRange := toIdenticalList(t, make(map[int32]struct{}))

	if len(xRange) != len(yRange) {
		return false
	}

	// zip
	zipped := make(map[int32]int32)
	yList := toList(t)
	i := 0
	for _, x := range s {
		value := yList[i]
		if zipped[x] != 0 && zipped[x] != value {
			return false
		}
		zipped[x] = value
		i++
	}
	return true
}

func toList(str string) [] int32 {
	res := make([] int32, len(str))
	i := 0
	for _, k := range str {
		res[i] = k
		i++
	}
	return res
}

func toIdenticalList(str string, setLike map[int32]struct{}) [] int32 {
	for _, c := range str {
		setLike[c] = struct{}{}
	}
	keys := make([] int32, len(setLike))
	i := 0
	for k := range setLike {
		keys[i] = k
		i++
	}
	return keys
}
