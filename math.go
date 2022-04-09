package main

import (
	"fmt"
)

var primes = []uint64{2, 3, 5, 7, 11, 13, 17, 19}

func main() {

	fmt.Println(MDC(24, 15))
	fmt.Println(MDC(11, 50))
	fmt.Println(MDC(25, 100))
	fmt.Println(MDC(300, 192))
	fmt.Println(next_prime(true))
	fmt.Println(next_prime(true))
	fmt.Println(next_prime(true))
	fmt.Println(next_prime(true))
	fmt.Println(next_prime(true))

	prime_after(1000)

	fmt.Println(primes[len(primes)-1])

	for i := 29; i < 100; i += 3 {
		fmt.Printf("%d is prime: ", i)
		is_prime := "false"
		if is_prime_bf(uint64(i)) {
			is_prime = "true"
		}
		fmt.Println(is_prime)
	}

	for i := 800; i < 810; i += 1 {
		fmt.Printf("%d is prime: ", i)
		is_prime := "false"
		if is_prime_bf(uint64(i)) {
			is_prime = "true"
		}
		fmt.Println(is_prime)
	}

	print_before(797)
	print_before(800)
	print_before(1999)
	print_before(2000)
	print_before(1999)
	print_before(2000)
	print_before(2001)
	print_before(2002)
	print_before(2003)
	print_before(2004)
	print_before(2005)
	print_before(0)
	print_before(1)
	print_before(2)
	print_before(3)
	print_before(4)
	print_before(5)

	print_after(800)
	print_after(809)
	print_after(3000)
	print_after(3001)

	print_after(0)
	print_after(1)
	print_after(2)
	print_after(3)
	print_after(4)
	print_after(5)
	print_after(3000)
	print_after(3001)
	print_after(3002)

}

func print_before(n uint64) {
	fmt.Printf("prime before %d: %d\n", n, prime_before(n))
}

func print_after(n uint64) {
	fmt.Printf("prime after %d: %d\n", n, prime_after(n))
}

func next_prime(add bool) uint64 {
	l := len(primes)
	next := primes[l-1] + 2
	found := false
	for !found {
		i := 0
		last := primes[i]
		last2 := last * last
		for ; last2 <= next; last2 = last * last {
			r := next % last
			if r == 0 {
				next += 2
				break
			}
			i++
			last = primes[i]
		}
		found = last2 > next
	}
	if add {
		primes = append(primes, next)
	}
	return next
}

func prime_after(max uint64) uint64 {
	if max < 2 {
		return 2
	}
	p := len(primes) - 1
	last := primes[p]
	if max >= last {
		for max >= last {
			last = next_prime(true)
		}
	} else {
		p_max := p
		p_min := 0
		for {
			p1 := p_min + (p_max-p_min)/2
			p2 := p1 + 1
			prime1 := primes[p1]
			prime2 := primes[p2]
			if max >= prime1 {
				if max < prime2 {
					return prime2
				}
				p_min = p1
			} else {
				p_max = p1
			}
		}
	}
	return last
}

func is_prime_bf(n uint64) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	next := func(i *uint64) {
		if *i == 2 {
			*i = 3
		} else {
			*i = *i + 2
		}
	}

	for i := uint64(2); n > i*i; next(&i) {
		if n%i == 0 {
			return false
		}
	}
	return true

}

func prime_before(max uint64) uint64 {
	if max < 3 {
		return 0
	}
	p := len(primes) - 1
	last := primes[p]
	if max <= last {
		p_max := p
		p_min := 0
		for {
			p1 := p_min + (p_max-p_min)/2
			p2 := p1 + 1
			prime1 := primes[p1]
			prime2 := primes[p2]
			if max <= prime2 {
				if max > prime1 {
					return prime1
				}
				p_max = p1
			} else {
				p_min = p1
			}
		}
	}
	next := next_prime(false)
	for next < max {
		primes = append(primes, next)
		next = next_prime(false)
	}
	return primes[len(primes)-1]
}

func MDC(a uint64, b uint64) uint64 {
	if b > a {
		c := b
		b = a
		a = c
	}
	r := a % b
	for r > 0 {
		a = b
		b = r
		r = a % b
	}
	return b
}

func MMC(a uint64, b uint64) uint64 {
	return a * b / MDC(a, b)
}
