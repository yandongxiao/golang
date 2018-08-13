/*
   Except when printed using the verbs %T and %p, special formatting
   considerations apply for operands that implement certain interfaces. In
   order of application:

   1. If the operand is a reflect.Value, the operand is replaced by the
   concrete value that it holds, and printing continues with the next rule.

   2. If an operand implements the Formatter interface, it will be invoked.
   Formatter provides fine control of formatting.

   3. If the %v verb is used with the # flag (%#v) and the operand
   implements the GoStringer interface, that will be invoked.

   If the format (which is implicitly %v for Println etc.) is valid for a
   string (%s %q %v %x %X), the following two rules apply:

   4. If an operand implements the error interface, the Error method will
   be invoked to convert the object to a string, which will then be
   formatted as required by the verb (if any).

   5. If an operand implements method String() string, that method will be
   invoked to convert the object to a string, which will then be formatted
   as required by the verb (if any).

   For compound operands such as slices and structs, the format applies to
   the elements of each operand, recursively, not to the operand as a
   whole. Thus %q will quote each element of a slice of strings, and %6.2f
   will control formatting for each element of a floating-point array.
*/
