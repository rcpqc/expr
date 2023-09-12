# expr
一个基于go/ast的轻量级右值表达式推断引擎

## 安装与使用
```go get -u github.com/rcpqc/expr@latest```

```
package main

import (
	"fmt"

	"github.com/rcpqc/expr"
)

func main() {
	ex, _ := expr.Comp(`a+b-c*d/2`)
	val, _ := expr.Eval(ex, expr.Vars{"a": 1.2, "b": 3.4, "c": 0.7, "d": 5.3})
	fmt.Print(val)
}
```

## 语法支持
由于使用Golang原生的抽象语法树库作表达式解析，所以表达式支持的语法为标准Golang语法的子集，目前支持以下语法特征：

- 常数  
  123, "123", 123.0, true
- 变量  
  a, b, c
- 一元运算  
  !a, -b
- 二元运算  
  a+b, a-b, a*b, a/b, a%b, a>b, a>=b, a<b, a<=b, a==b, a!=b, a&&b, a||b
- 括号  
  (a+b)*(a-b)
- 索引  
  a[b], a[b:c], a["b"]
- 选择器  
  a.b
- 函数调用  
  a(b)

## 内建函数
- 类型转换  
  int(), uint(), int64(), uint64(), float32(), flaot64() ...
- 数学函数  
  abs(), ceil(), round(), floor(), sin(), cos(), log(), sigmoid() ...
- 字符串函数  
  sfmt(), split(), stoi(), stof(), sfind() ...
- 时间函数  
  tnow(), tfmt(), tprs() ...
- 哈希函数  
  md5(), sha1(), sha256() ...
- golang函数  
  len(), cap() ...

## 例子
可参考[单测](/test)