package main 

func use(_ int) {}

func main() {
	b := 1  // b : declared but not used

	var c int 
	d := c 		// d : declared but not used

	e := 1
	use(e)

	f := 1
	_ = f    // technically used for an explicit no identifier (_) appraoch useful if you want to tell the reader that you want to ignore the value
}

