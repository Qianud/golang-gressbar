package gobar

import "fmt"

type Bar struct {
	totalValue int64
	currValue  int64
	graph      string
	rate       string
}

//初始化该方法
func Init() *Bar {
	bar := new(Bar)
	return bar
}

//进度条主程序
func (bar *Bar) NewOption(start, totalValue int64) *Bar {
	bar.currValue = start
	bar.totalValue = totalValue
	if bar.graph == "" {
		bar.graph = "█"
	}
	return bar
}

//可定制进度条样式
func (bar *Bar) NewOptionWithGraph(start, total int64, graph string) {
	bar.graph = graph
	bar.NewOption(start, total)
}

/**

 */
func (bar *Bar) Play(value int64) {
	val := float64(bar.totalValue) / 50
	prePercent := int32(float64(bar.currValue) / val)
	nowPercent := int32(float64(value) / val)
	for i := prePercent + 1; i <= nowPercent; i++ {
		bar.rate += bar.graph
	}
	bar.currValue = value

	fmt.Printf("\r[%-50s]%0.2f%%   	%8d/%d", bar.rate, float64(bar.currValue)/float64(bar.totalValue)*100,
		bar.currValue, bar.totalValue)
}

/**
完成调用
*/
func (bar *Bar) Finish() {
	val := float64(bar.totalValue) / 50
	prePercent := int32(float64(bar.currValue) / val)
	for i := prePercent + 1; i <= 50; i++ {
		bar.rate += bar.graph
	}
	bar.currValue = bar.totalValue
	fmt.Printf("\r[%-50s]%0.2f%%   	%8d/%d", bar.rate, float64(bar.currValue)/float64(bar.totalValue)*100,
		bar.currValue, bar.totalValue)
	fmt.Println()
}

func (bar *Bar) Stop() {
	fmt.Println()
}
