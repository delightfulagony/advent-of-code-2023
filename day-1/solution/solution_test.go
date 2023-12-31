// author: delightfulagony<agony@disroot.org> year: 2023

package solution

import "testing"

func TestCalibrationValueParser(t *testing.T) {

	var input []string = []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	calibrationValue, err := CalibrationValueParser(input)

	if calibrationValue != 142 || err != nil {
		t.Fatalf(
			`CalibrationValueParser(%s) = %d, %s wanted %d`,
			input, calibrationValue, err, 142,
		)
	}

}

func TestPopulateReplaceOrdered(t *testing.T) {

	var input string = "eightwothree"

	var replaceMap map[string]rune = map[string]rune {
		"two": '2',
		"three": '3',
		"eight": '8',
	}


	var replaceOrdered []string = make([]string, len(input))
	for s, _ := range replaceMap {
		populateReplaceOrdered(input, s, replaceOrdered, 0)
	}

	if replaceOrdered[0] != "eight" || replaceOrdered[4] != "two" || replaceOrdered[7] != "three" {
		t.Fatalf(
			`populateReplaceOrdered(%s) = %s`,
			input, replaceOrdered,
		)
	}
}

func TestImprovedCalibrationValueParser(t *testing.T) {

	var input []string = []string{
		"threeight",
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	calibrationValue, err := ImprovedCalibrationValueParser(input)

	if calibrationValue != 281+38|| err != nil {
		t.Fatalf(
			`ImprovedCalibrationValueParser(%s) = %d, %s wanted %d`,
			input, calibrationValue, err, 281+38,
		)
	}

}

