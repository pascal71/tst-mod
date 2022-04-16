package tstmod

// Calculate Fac of n recursively

func Fac(n uint64) uint64 {

	if n == 0 {
		return 1
	}

	return n * Fac(n-1)
}

// Calculate Fib nr recursively

func Fib(n uint64) uint64 {

	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)

}
