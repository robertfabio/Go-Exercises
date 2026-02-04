package main

func main() {
	i := 0
	for j := 0; j < 10000; j++ {
		go func() {
			i++
		}()
	}
	println(i)
}

func changeNumber(i *int, number int) {
	*i = number
}
