// An interface type specifies a collection of method prototypes.
// Each type has a method set associated with it.
//  1. For a non-interface type, its method set is the
//	prototype collection of all the methods (either
//  explicit or implicit ones) declared for it.
//	2. For an interface type, its method set is the method
//	prototype collection it specifies.
//
// We can view each interface value as a box to
// encapsulate a non-interface value. To box/encapsulate
// a non-interface value into an interface value, the type
// of the non-interface value must implement the type of
// the interface value.
// When a T value is converted (assigned) to an I value,
//	1. if type T is a non-interface type, then a copy of
//  the T value is boxed (or encapsulated) into the
//  resultant (or destination) I value. The time complexity
//	of the copy is O(n), where n is the size of copied T
//	value.
//	2. if type T is also an interface type, then a copy of
//  the value boxed in the T value is boxed (or encapsulated
//	) into the resultant (or destination) I value.
//  The standard Go compiler makes an optimazation here,
//  so the time complexity of the copy is O(1),
//	instead of O(n).
//	NOTE: The type information of the boxed value is also
//	stored in the resultant (or destination) interface value.
//
// The direct part of the dynamic value of an interface
// value is immutable, though we can replace the dynamic
// value of an interface value with another dynamic value.
//
// When an untyped value (except untype nil values) is
// assigned to a blank interface value, the untyped value
// will be first converted to its default type.
// (In other words, we can think the untyped value is
// deduced as a value of its default type.)
//
// An interface type can embed another named interface type.
// The final effect is the same as unfolding the method prototypes
// specified by the embedded interface type into the definition
// body of the embedding interface type.
package main

// Type I is a defined blank interface type.
// a blank interface type:
// If the method set of an arbrtrary type T,
// T may be an interface type or not, is a super set
// of the method set of an interface type I, then we
// say type T implements interface I.
//
// An interface type always implements itself.
// Two interface types with the same method set implement each other.
// any type implements any blank interface type.
// If a type T implements an interface type I, then
// any value of type T can be implicitly converted to
// type I.
type I interface{}

func main() {
	// An unnamed blank interface type.
	// Two unnamed interface types are identical if
	// their method sets are identical. Please note,
	// non-exported method names, which start with
	// lower-case letters, from different packages will
	// be always viewed as two different method names,
	// even if the two method names are the same in literal.
	var i interface{}
}
