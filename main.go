package main

func main() {
	fl_pth := "C:/Users/dkh18/text_data/SouthParkData-master/SouthParkData-master/All-seasons.csv"
	// f, err := os.Open(fl_pth)
	// if err != nil {
	// 	log.Fatal()
	// }
	// defer f.Close()

	// b1 := make([]byte, 5)
	// n1, err := f.Read(b1)
	// if err != nil {
	// 	log.Fatal()
	// }

	// fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	Ts_op_suffixarray(fl_pth, []byte(`Cool`))
	readoneFile(fl_pth)
}
