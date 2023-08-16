package main

import "fmt"

// Number Declares a Number interface type to use as a type constraint. Now, when you want to constrain a type parameter
// to either int64 or float64, you can use this Number type constraint instead of writing out int64 | float64.
type Number interface {
	int64 | float64
}

func main() {

	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	// Sum using two functions, i.e without generics
	fmt.Printf("Non-Generic Sums: %v and %v\n", SumInts(ints), SumFloats(floats))

	// Specify type arguments – the type names in square brackets – to be clear about the types that should replace
	// type parameters in the function you’re calling.
	fmt.Printf("Non-Generic Sums: %v and %v\n", SumNums[string, int64](ints), SumNums[string, float64](floats))

	// You can omit type arguments in calling code when the Go compiler can infer the types you want to use.
	// The compiler infers type arguments from the types of function arguments.
	//Note that this isn’t always possible. For example, if you needed to call a generic function that had no arguments,
	// you would need to include the type arguments in the function call.
	// Call the generic function, omitting the type arguments.
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n", SumNums(ints), SumNums(floats))

	// Call SumNumbers with each map, printing the sum from the values of each.
	fmt.Printf("Generic Sums with Constraint: %v and %v\n", SumNumbers(ints), SumNumbers(floats))
}

// SumInts takes a map of string to int64 values.
// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats takes a map of string to float64 values.
// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumNums Declare a function with two type parameters (inside the square brackets), K and V, and one argument
// that uses the type parameters, m of type map[K]V. The function returns a value of type V.
// Specify for the K type parameter the type constraint comparable. Intended specifically for cases like these,
// the comparable constraint is predeclared in Go. It allows any type whose values may be used as an operand of the
// comparison operators == and !=. Go requires that map keys be comparable.
// So declaring K as comparable is necessary so you can use K as the key in the map variable.
// It also ensures that calling code uses an allowable type for map keys.
func SumNums[K comparable, V int64 | float64](m map[K]V) V {
	// Specify for the V type parameter a constraint that is a union of two types: int64 and float64.
	// Using | specifies a union of the two types, meaning that this constraint allows either type.
	// Either type will be permitted by the compiler as an argument in the calling code.
	// Specify that the m argument is of type map[K]V, where K and V are the types already specified for the
	// type parameters.
	// Note that we know map[K]V is a valid map type because K is a comparable type.
	// If we hadn’t declared K comparable, the compiler would reject the reference to map[K]V.
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

// SumNumbers Declare a generic function with the same logic as the generic function you declared previously, but with
// the new interface type instead of the union as the type constraint. As before, you use the type parameters for the
// argument and return types.
// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
