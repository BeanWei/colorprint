# 在windows终端中输出颜色字体

## 测试
```
$ go run test.go
package main

import (
	"github.com/BeanWei/colorprint"
)

func main() {

	colorprint.Print("Hello\n", colorprint.Deepblue)
	colorprint.Print("Hello\n", colorprint.Green)
	colorprint.Print("Hello\n", colorprint.Skyblue)
	colorprint.Print("Hello\n", colorprint.Red)
	colorprint.Print("Hello\n", colorprint.Purple)
	colorprint.Print("Hello\n", colorprint.Yellow)
	colorprint.Print("Hello\n", colorprint.White)
}
```

## 效果
![Alt text](https://github.com/BeanWei/colorprint/blob/master/Screenshot/test.PNG)