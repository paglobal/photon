package main

func remove[C any](s []C, i int) []C {
	s[i] = s[len(s)-1]

	return s[:len(s)-1]
}
