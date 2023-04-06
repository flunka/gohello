package das

// import "fmt"
import "math"

func LinearSearch(haystack []int, needle int) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}

func BinarySearch(haystack []int, needle int) bool {
	start_index := 0
	end_index := len(haystack) - 1
	for i := (end_index-start_index)/2 + start_index; start_index < end_index; i = (end_index-start_index)/2 + start_index {
		value := haystack[i]
		if value == needle {
			return true
		} else if value < needle {
			start_index = i + 1
		} else {
			end_index = i - 1
		}
	}
	return haystack[start_index] == needle

}

func CrystalBallSearch(breaks []bool) int {
	step := int(math.Sqrt(float64(len(breaks))))
	i := step
	for ; i < len(breaks); i += step {
		if breaks[i] {
			break
		}
	}

	i -= step

	for j := i; j < i+step; j++ {
		if breaks[j] {
			return j
		}
	}
	return -1
}
