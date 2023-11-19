package sqldumptocsv_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	sqldumptocsv "github.com/Piorosen/sqldump-to-csv"
)

func TestParse(t *testing.T) {
	// TODO
	read, err := os.Open("aaa.sql")
	if err != nil {
		t.Error(err, "not found aaa.sql")
	}
	defer read.Close()

	file, err := ioutil.ReadAll(read)
	if err != nil {
		t.Error(err, "not readable aaa.sql")
	}

	fmt.Printf("%s", sqldumptocsv.Parse(file))
}