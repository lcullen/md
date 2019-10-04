1.数据采集方式
    [郝林](http://wiki.jikexueyuan.com/project/go-command-tutorial/0.12.html)
    任何以go tool开头的Go命令内部指向的特殊工具都被保存在目录GOROOT/pkg/tool/GOOS_$GOARCH/中
    1. 直接 import _ "net/http/pprof"
        初始化http handler 可以查看 http serve and listen code
2.[数据展示方式](https://segmentfault.com/a/1190000016412013)
    指标：flat flat% cum cum% sum  
         函数运行相关 
         占用cpu
         父函数运行相关    
3. 如何用prometheus 做 goroutine 监控