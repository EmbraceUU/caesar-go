# EXERCISES IN GO

## basis exercises

**定义一个包内全局字符串变量**

A. var str string

B. str := ""

C. str = ""

D. var str = ""

(AD)

**通过指针变量 p 访问其成员变量 name**

A. p.name

B. (*p).name

C. (&p).name

D. p->name

(AB)

**关于接口和类的说法，下面说法正确的是**

A. 一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口

B. 实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才合理

C. 类实现接口时，需要导入接口所在的包

D. 接口由使用方按自身需求来定义，使用方无需关心是否有其他模块定义过类似的接口

(ABD)

**关于init函数**

A. 一个包中，可以包含多个init函数

B. 程序编译时，先执行导入包的init函数，再执行本包内的init函数

C. main包中，不能有init函数

D. init函数可以被其他函数调用

(AB)

**关于循环语句**

A. 循环语句既支持for关键字，也支持while和do-while

B. 关键字for的基本使用方法与C/C++中没有任何差异

C. for循环支持continue和break来控制循环，但是它提供了一个更高级的break，可以选择中断哪一个循环

D. for循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量 

(CD)

**函数调用正确的方式**

```gotemplate
func add(args... int) int {
	sum :=0
	for _,arg := range args {
		sum += arg
	}
	return sum
}
```

A. add(1, 2)

B. add(1, 3, 7)

C. add([]int{1, 2})

D. add([]int{1, 3, 7}...)

(ABD)

