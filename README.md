# 文心一言 Web (beta)
Baidu Ernie Pure Go CLI and API  
第三方百度文心一言命令行与接口  

低门槛，普通百度用户也可使用，填入 `bduss` 即可运行。

## 测试用 CLI 使用方法
在config.yaml 中填写你的百度 bduss，bduss可以从浏览器 `F12` 中获取
```shell
# 双击直接运行，或在shell中运行
./ernie
```

## API 使用方法
```go
package main

import (
	"github.com/r3inbowari/ernie"
	"fmt"
)

func main() {
	token := "你的百度 bduss"
	prompts := "你的问题"
	stream, _ := ernie.New(token).Query(prompts)
	for {
		select {
		case event := <-stream.Events:
			seg, err1 := ernie.ParseStreamSegment(event.Data())
			if err1 != nil {
				continue
			}
			if !seg.Empty() {
				fmt.Print(seg.Text())
			}
			if seg.IsCompleted() {
				return
			}
		}
	}
}

```

## 获取 bduss
![image](https://github.com/r3inbowari/ernie-go/assets/30739857/a026f677-385e-4c52-8217-991a28b0c690)


## 其它
![image](https://github.com/r3inbowari/ernie-go/assets/30739857/c76dfa79-7b07-4933-b442-1b2978891f74)


## 注意事项
Only for test, Please abide by Baidu User Terms. 请遵守百度用户条款，切勿滥用。 
