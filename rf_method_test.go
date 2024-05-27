package main

import (
	"testing"
)

var fl_pth string = "C:/Users/dkh18/text_data/SouthParkData-master/SouthParkData-master/All-seasons.csv"
var search_txt string = "Cool"

func TestGetonefile(t *testing.T) {
	tst_data := getonefile(fl_pth)

	if tst_data == nil {
		t.Error("Failed No file was opened")
	}

}

func BenchmarkGetoneFile(b *testing.B) {
	getonefile(fl_pth)
}

func TestTs_op_suffixarray(t *testing.T) {
	offsets := Ts_op_suffixarray(fl_pth, []byte(search_txt))

	if len(offsets) != 86 {
		t.Error("Incorrect return value from Ts_op_suffixarray value should be 86")
	}
}

func BenchmarkTs_op_suffixarray(b *testing.B) {
	Ts_op_suffixarray(fl_pth, []byte(search_txt))
}

func TestReadoneFile(t *testing.T) {
	rtrnval := readoneFile(fl_pth)

	if rtrnval != 84 {
		t.Error("Recived the Wrong Return value should be 84")
	}
}

func BenchmarkReadoneFile(b *testing.B) {
	readoneFile(fl_pth)

}
