package solution

import(
	"fmt"
	"strconv"
	"strings"
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

func ImprovedCalibrationValueParser(input []string) (calibrationValue int,
		err error) {

	var replaceMap map[string]string = map[string]string {
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}

	for i, str := range input {

		var replaceOrdered []string = make([]string, len(str))

		for s, _ := range replaceMap {
			if i := strings.Index(str, s); i != -1 {
				replaceOrdered[i] = s
			}
		}

		for _, s := range replaceOrdered {
			if s != "" {
				str = strings.ReplaceAll(str, s, replaceMap[s])
			}
		}

		input[i] = str
	}

	return CalibrationValueParser(input)

}
