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

func TestImprovedCalibrationValueParser(t *testing.T) {

	var input []string = []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	calibrationValue, err := ImprovedCalibrationValueParser(input)

	if calibrationValue != 281 || err != nil {
		t.Fatalf(
			`ImprovedCalibrationValueParser(%s) = %d, %s wanted %d`,
			input, calibrationValue, err, 281,
		)
	}

}
