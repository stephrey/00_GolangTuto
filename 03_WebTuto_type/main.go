package main

import "fmt" //librairie with stdio

/* comment */
func main() {
	var (
		b  bool
		s  string
		i  int   //64bit on a 64bits machine
		u  uint  //64bit on a 64bits machine
		u8 uint8 // 0- 255
		i8 int8  //-128 - 127
		f  float32
	)

	b = true
	s = "steph"
	i = -15
	u = 15
	u8 = 255
	i8 = -127
	f = 3.14159

	fmt.Println(b, s, i, u, u8, i8, f)
}

// main.main()
