package Sorter

import (
	"sort"
	"github.com/Inkeliz/InkDir/Types"
)

type data struct {
	Key   string
	Value int64
}

type dataList []data

func Sort(info Types.Info) dataList {

	p := make(dataList, len(info))
	i := 0

	for key, value := range info {
		p[i] = data{key, value}
		i++
	}

	sort.Sort(p)

	return p
}

func (d dataList) Len() int {
	return len(d)
}
func (d dataList) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d dataList) Less(i, j int) bool {
	return d[i].Value > d[j].Value
}

func Sum(info Types.Info) int64 {
	var total int64

	for _, value := range info {
		total += value
	}

	return total
}
