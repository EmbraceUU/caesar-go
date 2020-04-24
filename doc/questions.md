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
