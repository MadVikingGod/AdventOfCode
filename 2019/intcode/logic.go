package intcode

func add(r1, r2 int) int {
	return r1 + r2
}
func mul(r1, r2 int) int {
	return r1 * r2
}
func less(r1, r2 int) int {
	if r1 < r2 {
		return 1
	}
	return 0
}
func equal(r1, r2 int) int {
	if r1 == r2 {
		return 1
	}
	return 0
}
