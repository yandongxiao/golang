package main

func ExampleNestedComposite() {
	// If a composite literal nested many other composite literals,
	// then those nested composited literals can simplified to the form {...}.
	var heads = []*[4]byte{
		&[4]byte{'P', 'N', 'G', ' '},
		&[4]byte{'G', 'I', 'F', ' '},
		&[4]byte{'J', 'P', 'E', 'G'},
	}

	var heads2 = []*[4]byte{
		{'P', 'N', 'G', ' '},
		{'G', 'I', 'F', ' '},
		{'J', 'P', 'E', 'G'},
	}

	_ = heads
	_ = heads2
	// Output:
	//
}

func ExampleNestedComposite2() {
	type LangCategory struct {
		dynamic bool
		strong  bool
	}
	// A value of map type whose key type is a struct type and
	// whose element type is another map type "map[string]int".
	_ = map[LangCategory]map[string]int{
		LangCategory{true, true}: map[string]int{
			"Python": 1991,
			"Erlang": 1986,
		},
		LangCategory{true, false}: map[string]int{
			"JavaScript": 1995,
		},
		LangCategory{false, true}: map[string]int{
			"Go":   2009,
			"Rust": 2010,
		},
		LangCategory{false, false}: map[string]int{
			"C": 1972,
		},
	}
	_ = map[LangCategory]map[string]int{
		{true, true}: {
			"Python": 1991,
			"Erlang": 1986,
		},
		{true, false}: {
			"JavaScript": 1995,
		},
		{false, true}: {
			"Go":   2009,
			"Rust": 2010,
		},
		{false, false}: {
			"C": 1972,
		},
	}
	// Output:
	//
}
