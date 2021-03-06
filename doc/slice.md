## slice

### 数据结构

* compile的Slice类型，只有elem一个属性
* runtime的SliceHeader类型，包括: 
    * Data: 指向数组的指针
    * Len: 当前切片的长度
    * Cap: 切片的容量, 对应数组的大小
* 数组在编译期间已经基本固定
* 切片在运行时动态的

### 初始化

有三种方式初始化切片，在编译和运行阶段的方式也各不相同。
* `arr[0:3] or slice[0:3]`使用下标
* `slice := []int{1, 2, 3}`字面量
* `slice := make([]int, 3)`使用make关键字

### 访问优化

编译期间会对访问操作做出优化
* len()和cap()操作, 在一些情况下会直接转换成切片的len和cap;
* 通过index的访问, 会直接转换成内存地址。

### Q&A

#### 修改新的切片，原切片的数据会变化吗 ？

原切片会变化，因为新切片初始化时，没有copy底层数组，只是创建一个指向原数组的结构体。

#### 切片和数组的区别是什么，为什么`var a []int; a=append(a, 1, 2, 3, 4)`是切片的扩容 ？难道数组时静态的，扩容和append只属于切片吗 ？

#### `append(s, 1)`不能编译, 为什么编译阶段还要区分`append(s,1)和s=append(s,1)` ?