package gel

// None produces an empty Text Node.
func None() View {
	return Text("")
}

// Maybe check if val is nil, a view or a viewable, and if so returns
// a View based on the val, but if it's none of these then it returns
// an empty Text node.
func Maybe(val interface{}) View {
	if val == nil {
		return None()
	}
	view, ok := val.(View)
	if ok {
		return view
	}
	s, ok := val.(string)
	if ok {
		return Text(s)
	}
	return None()
}

// Default takes two values, should the first be unconvertible to a Viewable
// or a view it will then attempt to convert the second.  Nil is not convertible
// so in cases where the first is nil, the second will be used if possible.
// Should they both be nil the View will be of an empty Text node.
func Default(val interface{}, def interface{}) View {
	if val == nil {
		return Maybe(def)
	}
	switch v := val.(type) {
	case View, string:
		return Maybe(v)
	default:
		return Maybe(def)
	}
}
