package structure_benchmark

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"

	randomdata "github.com/Pallinder/go-randomdata"
)

const BenchmarkMinSize = 1000000
const NumberMax = 50000000

var BenchmarkSizeList = []int{10, 100, 1000, 10000, 100000, 1000000}

func TestSave(t *testing.T) {
	_, err := os.Stat("./Readme.md")
	if err == nil {
		for _, size := range BenchmarkSizeList {
			createData(size)
		}
	}
}

func createData(csize int) {

	fname := spew.Sprintf("benchmark%d.data", csize)

	f, err := os.OpenFile(fname, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
	}

	var l []int
	m := make(map[int]int)
	for i := 0; len(l) < csize; i++ {
		v := randomdata.Number(0, NumberMax)
		if _, ok := m[v]; !ok {
			m[v] = v
			l = append(l, v)
		}
	}

	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	encoder.Encode(l)
	lbytes := result.Bytes()
	f.Write(lbytes)
}

func LoadTestData(csize int) []int {
	fname := spew.Sprintf("../benchmark%d.data", csize)
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Println(err)
	}
	var l []int
	decoder := gob.NewDecoder(bytes.NewReader(data))
	decoder.Decode(&l)
	return l
}
