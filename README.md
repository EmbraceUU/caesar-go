# caesar-go

## arithmetic

算法部分, 日常学习算法的记录

## geek-learning

这部分是学习极客时间的课, 做一些笔记。

## knowledge-points

这部分是一些知识点, 或者积累的一些问题点, 不是系统性的整理笔记

## learning-notes

这部分是整理《go学习笔记》的知识点, 以及一些相关的demo。

## doc

这部分是写的一些心得或者感想之类的。

## design

这部分是设计模式相关的学习整理。

## utils

这部分是慢慢积累的一些工具类和工具接口, 以及一些api的用法案例。

## pkg-example

常用第三方库的example

## go-example

go语法相关的example

## 优化CPU总结

1. sync.Map, 在读写频繁的情况下使用

2. 使用channel将map读写分离

3. 将所有的getXXXMap去掉, 换成以切片或者具体值的形式返回

```txt
将所有getMap接口去掉, 如果需要读取, 直接写函数, 返回结构体, 因为结构体都是copy,
```

4. 慎用pool, 会将cpu飙高

5. 不做重复的工作, 想明白再写

6. 必要时, 可以用空间换取时间, 用内存换取cpu

7. goroutine中是否有不必要的goroutine

8. clear map的时候也要加锁

## 开发注意事项

### 日志

1. 查询数据接口, 如果有nil或者len为0的情况, 打印时需要将入参一同打印, 便于查找问题

### 公共库

1. 不可以修改公共库

2. 公共库不可以调用业务模块的代码, 更不能通过业务模块获取数据或者读取缓存