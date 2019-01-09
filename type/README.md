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

## Type alias declaration

