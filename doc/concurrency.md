## CONCURRENCY NOTE

### chan

并发中chan的用法整理.

#### chan.1

下面一段代码, 是来自go的编译器部分, 词法和语法分析步骤, 该过程使用goroutine并发执行syntax.Parse. 
里面使用了sem chan来控制同时并发的数量
```
// 同时并发的数量
sem := make(chan struct{}, runtime.GOMAXPROCS(0)+10)
for _, filename := range filenames {
    p := &noder{
        basemap: make(map[*syntax.PosBase]*src.PosBase),
        err:     make(chan syntax.Error),
    }
    noders = append(noders, p)

    go func(filename string) {
        // 阻塞goroutine
        sem <- struct{}{}
        // 延迟释放sem的一个空间
        defer func() { <-sem }()
        defer close(p.err)
        base := syntax.NewFileBase(filename)

        f, err := os.Open(filename)
        if err != nil {
            p.error(syntax.Error{Msg: err.Error()})
            return
        }
        defer f.Close()

        p.file, _ = syntax.Parse(base, f, p.error, p.pragma, syntax.CheckBranches) // errors are tracked via p.error
    }(filename)
}
```