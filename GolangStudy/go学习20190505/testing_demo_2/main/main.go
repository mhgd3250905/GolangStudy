package main

func addUpper(n int) int {
	res:=0
	for i := 0; i < n; i++ {
		res+=i
	}
	return res
}

func main() {

}

func getSub(i1 int, i2 int) int {
	return i1-i2
}