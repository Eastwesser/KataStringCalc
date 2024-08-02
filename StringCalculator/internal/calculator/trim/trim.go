package trim

// TrimStringsAfter40 Если строка длиннее 40 символов, обрезает её до 40 символов и добавляет '...'
func TrimStringsAfter40(s string) string {
	if len(s) > 40 {
		return s[:40] + "..."
	}
	return s
}
