package season

import "errors"

func Season(month int) (string, error) {
	switch month {
	case 3:
	case 4:
	case 5:
		return "春", nil
	case 6:
	case 7:
	case 8:
		return "夏", nil
	case 9:
	case 10:
	case 11:
		return "秋", nil
	case 12:
	case 1:
	case 2:
		return "冬", nil

	}
	return "", errors.New("invalid month value")
}
