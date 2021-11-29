package main

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}
