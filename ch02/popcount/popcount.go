package popcount

var pc[256] byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	res := 0
	for i := uint(0); i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}

//func main() {
//	for _, arg := range os.Args[1:] {
//		val, err := strconv.ParseUint(arg, 10, 64)
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "Invalid argument '%s'\n", arg)
//			os.Exit(2)
//		}
//
//		pc := PopCount(val)
//		fmt.Printf("%d (%b): %d\n", val, val, pc)
//	}
//}