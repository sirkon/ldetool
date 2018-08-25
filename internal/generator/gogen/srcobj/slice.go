package srcobj

// SliceFrom generates {rest}[index:] expression
func SliceFrom(expr, from Source) Source {
	return Index{
		Src:   expr,
		Index: OperatorColon(from, Raw("")),
	}
}

// SliceTo generates {rest}[:index] expression
func SliceTo(expr, to Source) Source {
	return Index{
		Src:   expr,
		Index: OperatorColon(Raw(""), to),
	}
}

func Slice(expr, from, to Source) Source {
	return Index{
		Src:   expr,
		Index: OperatorColon(from, to),
	}
}
