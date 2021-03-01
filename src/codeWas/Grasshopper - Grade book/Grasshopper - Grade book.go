package kata

// GetGrade 分级
func GetGrade(a, b, c int) rune {
	score := (a + b + c) / 3
	grade := 'F'
	switch {
	case score >= 90:
		grade = 'A'
	case 80 <= score:
		grade = 'B'
	case 70 <= score:
		grade = 'C'
	case 60 <= score:
		grade = 'D'
	}
	return grade
}
