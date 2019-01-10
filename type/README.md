# types

Up to now (Go 1.11), Go has 26 kinds of types. (包括unsafe types)

## Basic Types

- Built-in string type: string.
- Built-in boolean type: bool.
- Built-in numeric types:
    - int8, uint8 (byte), int16, uint16, int32 (rune), uint32, int64, uint64, int, uint, uinptr.
    - float32, float64.
    - complex64, complex128.

## Composite Types

In Go, the form T{...}, where T must be a type literal or a type name,
is called as a composite literal and is used as the value literals of
some kinds of types, including struct types and the container types introduced later.

- pointer types - C pointer alike.
- struct types - C struct alike.
- function types - functions are first-class types in Go.
- container types:
    - array types - fixed-length container types.
    - slice type - dynamic-length and dynamic-capacity container types.
    - map types - maps are associative arrays (or dictionaries). The standard Go compiler implements maps as hashtables.
- channel types - channels are used to synchronize data among goroutines (the green threads in Go).
- interface types - interfaces play a key role in reflection and polymorphism.

## Type Definitions

Type Definitions 包括： type declaration 和 type alias declaration.

### Type Declaration

语法: `type identifiers type_literal`

```
// Define a solo new type.
type NewTypeName SourceType

// Define multiple new types together.
type (
	NewTypeName1 SourceType1
	NewTypeName2 SourceType2
)
```

### Type alias declaration

```
type (
	Name = string
	Age  = int
)

type table = map[string]int
type Table = map[Name]Age
```

## Defined Types vs. Non-Defined Types

1. A defined type is a type defined in a type definition or an alias of another defined type.
2. All basic types are defined. ==> []string被称为type literal， 不算是defined types.
3. A non-defined type must be a composite type.
```
type A []string     # 依据规则1, A是defined type
type B = A          # 依据规则1，B是defined type
type C = []string   # 依据规则2，[]string是non-defined types, 所以C是non-defined types
```

## Named Types vs. Unnamed Types

An unnamed type must be a composite type.

```
// a的类型是A（Named types）
type A []string
var a A

// b的类型是[]string(Unnamed Types)
var b []string
```

## Underlying Types

In Go, each type has an underlying type. Rules:

1. for built-in basic types, the underlying types are themselves.
2. the underlying type of unsafe.Pointer is itself.
3. the underlying types of an unnamed type, which must be a composite type, is itself.
4. in a type declaration, the new declared type and the source type have the same underlying type. 
5. How to trace to the underlying type for a given user declared type? The rule is, when a built-in basic type, unsafe.Pointer or an unnamed type is met, the tracing will be stopped. 

```
// The underlying types of the following ones are both int.
type (
	MyInt int
	Age   MyInt
)

// The following new types have **different(不同的)** underlying types.
type (
	IntSlice = []int   // underlying type is []int
	MyIntSlice []MyInt // underlying type is []MyInt
	AgeSlice   []Age   // underlying type is []Age
)

// The underlying types of Ages and AgeSlice are both []Age.
type Ages AgeSlice
```

## Values

1. An instance of a type is called a value.
2. Each type has a zero value
3. There are several kinds of value representation forms in code, including literals, named constants, variables and expressions.

## Value Parts

1. In Go, each of such values has a direct part
2. some of them have one or more indirect parts
3. The indirect underlying parts of a value are referenced by its direct part through pointers.

## Value Sizes

1. the number of bytes occupied by the direct part of the value is called the size of the value.
2. All values of the same type have the same value size.
3. We can use the Sizeof function in the unsafe standard package to get the size of any value.

## Dynamic Type And Dynamic Value Of An Interface Value

1. Interface values are the values whose types are interface types.
2. Each interface value can box a non-interface value in it.
3. The type of the dynamic value is called the dynamic type of the interface value.
4. The value boxed in an interface value is called the dynamic value of the interface value.
