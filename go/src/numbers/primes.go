package numbers

// Primes generates a list a list of prime numbers up to n numbers, where n must be 1 or larger
func Primes(n int) []int {
	// prime numbers have only two positive divisors (themselves, and 1)
	primes := make([]int, n)
	primes[0] = 2

	count := 1
	current := 3
	for count < n {
		divisors := primes[0:count]
		divisorFound := false
		for _, divisor := range divisors {
			if current%divisor == 0 {
				divisorFound = true
				break
			}
		}
		if !divisorFound {
			primes[count] = current
			count++
		}
		current++
	}
	return primes
}
