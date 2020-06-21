# solutions

## 解题思路记录

### 递归 recursion

以自相似方法重复事务的过程.

```gotemplate
// 方法
func recursion(i int) int {
    if i==1 {
        return i
    }
    // 调用自身
    return i*recursion(i-1)
}
```

### 迭代 iteration

重复反馈过程的活动, 每一次迭代的结果作为下一次迭代的初始值.

循环变量/循环体/循环终止条件

```gotemplate
// 循环变量
var a int
// 循环体
// 循环终止条件
for i:=0; i<=100; i++ {
    a += i
}
return a
```

这种迭代是线性迭代, 时间会随着i的增加而增加.