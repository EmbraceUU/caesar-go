# QUESTIONS

> 记录一下工作中碰到的问题, 以前没有系统的记录过, 今后慢慢积累

## 数据库相关

### mysql update 一直报语法错误

那个傻逼用了关键字做字段名. 

### gorm sql 执行情况与预期不相符

通过结构体变量更新字段值, gorm库会忽略零值字段。就是字段值等于0, nil, "", false这些值会被忽略掉，不会更新。

打开gorm调试, 可以在获得db的时候, 打开log

```
db.LogMode(true)
```

调用sql时, 在db后面加上debug模式

```gotemplate
db.Debug().Model()...
```

### mongo 根据id查询不到数据

参数需要使用ObjectId类型

```gotemplate
db.collection.findOne({
    '_id': ObjectId('5d4e946888f74c4d4b876dec')
})
```

## 并发相关

### map 并发读写冲突

#### 频繁读, 少量的写

可以使用channel将map的读写分离, 专门有一个read map, 可以频繁读
将要写入的数据, 压入到channel中, 定时从channel中读出来写到一个write map中, 将write map一次性覆盖掉read map

#### 频繁写, 少量的读

直接对一个map读写锁, 暂时没想到会产生什么数据上的问题

#### 读写都频繁

使用sync.Map

[文章1](https://juejin.im/post/5c9b93a6e51d4579e94a9477)

[文章2](https://studygolang.com/articles/23370)