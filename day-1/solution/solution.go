package solution

import(
	"fmt"
	"strconv"
	"errors"
)

func CalibrationValueParser(input []string) (calibrationValue int, err error) {

	calibrationValue = 0
	var first, last rune
	for i, str := range input {
		first, last = ' ', ' '
		for _, run := range str {
			if run > 47 && run < 58 {
				if first == ' ' {
					first = run
				}
				last = run
			}
		}
		if first == ' ' {
			return 0, errors.New(fmt.Sprintf("no numbers found in line %d", i))
		}
		lineValue, err := strconv.Atoi(string(first)+string(last))
		if err != nil {
			return 0, errors.New(fmt.Sprintf("conversion error: %s", err))
		}
		calibrationValue += lineValue
	}
	return calibrationValue, nil

}
