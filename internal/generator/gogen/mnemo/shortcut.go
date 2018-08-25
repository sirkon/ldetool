package mnemo

func Make(value string) string {
	w := New()
	for _, r := range []rune(value) {
		w.WriteRune(r)
	}
	return w.String()
}
