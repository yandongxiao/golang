# reflect

package reflect 有众多类型，包括：

Method，SelectCase，SelectDir，SliceHeader，StringHeader，StructField，StructTag，
Type，Value等。关系如下：

![relation](./relation.png)

总结如下：

1. Type和Value是最重要的两个数据结构，尤其是Value
2. reflect.Value可以对内置类型，chan、map、struct等各种类型进行操作，如SetString, SetInt, Slice, MapKeys等。
3. reflect.Type的一大作用是给出了数据类型的模型，通过XxxOf()函数即可创建出该类型的一个新的实例。例如：```ArrayOf(count int, elem Type)```, ```FuncOf(in, out []Type, variadic bool) Type```, ```StructOf(fields []StructField) Type```等结构。

>
> StructOf用于动态构造struct类型对象。