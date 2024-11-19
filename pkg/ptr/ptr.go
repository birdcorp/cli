package ptr

func GetStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
