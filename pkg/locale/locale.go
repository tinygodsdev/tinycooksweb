package locale

const (
	En = "en"
	Ru = "ru"
)

func List() []string {
	return []string{
		En,
		Ru,
	}
}

func Default() string {
	return Ru
}

func IsValid(loc string) bool {
	switch loc {
	case En, Ru:
		return true
	}
	return false
}
