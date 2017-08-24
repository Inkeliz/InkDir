package Display

import (
	"fmt"
	"math"
	"strings"
	"github.com/Inkeliz/InkDir/Types"
	"github.com/Inkeliz/InkDir/Sorter"
	"strconv"
	"os"
)

var (
	alreadyLoaded float64
)

const (
	_        = iota
	KB int64 = 1 << (10 * iota)
	MB
	GB
	TB
)

func Error(err error) {
	if err != nil {
		createLine("Error")
		createHorizontalDivisor()
		createLine(err.Error())
		os.Exit(1)
	}
}

func Loading(qnt float64) {
	alreadyLoaded += 1

	loadedRepeat := int(math.Floor(alreadyLoaded * 100 / qnt))
	fmt.Print("\r" + strings.Repeat("=", loadedRepeat) + strings.Repeat("_", 100-loadedRepeat))
}

func Result(info Types.Info) {
	createLine("Name", "Size")
	createHorizontalDivisor()

	for _, data := range Sorter.Sort(info) {
		createLine(data.Key, createSize(data.Value), strconv.FormatInt(data.Value, 10))
	}


	createHorizontalDivisor()
	total := Sorter.Sum(info)
	createLine("Total", createSize(total), strconv.FormatInt(total, 10))
}

func Install(){
	createLine("Installation", )
	createHorizontalDivisor()

	createLine("The installation is finished successfully.")
	createLine("Note: You may need to restart your computer to take effect.")
}

func Path(path string) {
	createLine("Path", path)
}

func Time(time string) {
	createLine("Time", time)
}
// ----------------------------------------
// CREATE
// ----------------------------------------

func createName(text string) string {
	return StrpadRight(Strimwidth(text, 67, "..."), " ", 70)
}

func createSize(value int64) string {
	var Size string

	switch {
	case value > TB:
		Size = strconv.FormatInt(value/TB, 10) + "MB"
	case value > GB:
		Size = strconv.FormatInt(value/GB, 10) + "GB"
	case value > MB:
		Size = strconv.FormatInt(value/MB, 10) + "MB"
	case value > KB:
		Size = strconv.FormatInt(value/KB, 10) + "KB"
	default:
		Size = strconv.FormatInt(value, 10) + "B"
	}

	return StrpadRight(Size, " ", 10)
}

// ----------------------------------------
// DIVIDERS
// ----------------------------------------

func createLine(left string, right ...string) {
	if len(right) != 0{
		left = createName(left)
	}

	fmt.Println(left, strings.Join(right, " "))
}

func createHorizontalDivisor() {
	fmt.Println(strings.Repeat("-", 100))
}
