##Golang 进度条

###引用

require github.com/Qianud/golang-gressbar v1.0.0

###主程序引用

bar := Init()
	bar.NewOption(0,8)

	for i:=1;i<=8;i++{
		bar.Play(int64(i))
		time.Sleep(100*time.Millisecond)
	}

	bar.Finish() 