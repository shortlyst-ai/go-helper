package helper

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestVal(t *testing.T) {
	testValStr(t)
	testValInt(t)
	testValFloat(t)
	testValBool(t)
	testValTime(t)
}

func testValStr(t *testing.T) {
	testCases := map[string]struct {
		Input  *string
		Output string
	}{
		"ShouldReturnEmptyString_WhenInputIsNil": {
			Input:  nil,
			Output: "",
		},
		"ShouldReturnEmptyString_WhenInputIsEmptyString": {
			Input:  Pointer(""),
			Output: "",
		},
		"ShouldReturnString_WhenInputIsNotNil": {
			Input:  Pointer("SomeValue"),
			Output: "SomeValue",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := Val(testCase.Input)
			require.Equal(t, testCase.Output, actual)
		})
	}
}

func testValInt(t *testing.T) {
	testCases := map[string]struct {
		Input  *int64
		Output int64
	}{
		"ShouldReturn0_WhenInputIsNil": {
			Input:  nil,
			Output: int64(0),
		},
		"ShouldReturn0_WhenInput0": {
			Input:  Pointer[int64](0),
			Output: int64(0),
		},
		"ShouldReturnInt64_WhenInputIsNotNil": {
			Input:  Pointer[int64](432),
			Output: int64(432),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := Val(testCase.Input)
			require.Equal(t, testCase.Output, actual)
		})
	}
}

func testValFloat(t *testing.T) {
	testCases := map[string]struct {
		Input  *float64
		Output float64
	}{
		"ShouldReturn0_WhenInputIsNil": {
			Input:  nil,
			Output: float64(0),
		},
		"ShouldReturn0_WhenInput0": {
			Input:  Pointer[float64](0),
			Output: float64(0),
		},
		"ShouldReturnFloat_WhenInputIsNotNil": {
			Input:  Pointer(43.2),
			Output: 43.2,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := Val(testCase.Input)
			require.Equal(t, testCase.Output, actual)
		})
	}
}

func testValBool(t *testing.T) {
	testCases := map[string]struct {
		Input  *bool
		Output bool
	}{
		"ShouldReturnFalse_WhenInputIsNil": {
			Input:  nil,
			Output: false,
		},
		"ShouldReturnBool_WhenInputIsNotNil": {
			Input:  Pointer(true),
			Output: true,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := Val(testCase.Input)
			require.Equal(t, testCase.Output, actual)
		})
	}
}

func testValTime(t *testing.T) {
	now := time.Now()
	testCases := map[string]struct {
		Input  *time.Time
		Output time.Time
	}{
		"ShouldReturnZero_WhenInputIsNil": {
			Input:  nil,
			Output: time.Time{},
		},
		"ShouldReturnTime_WhenInputIsNotNil": {
			Input:  &now,
			Output: now,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := Val(testCase.Input)
			require.Equal(t, testCase.Output, actual)
		})
	}
}

func TestToString(t *testing.T) {
	t.Run("ShouldReturnString", func(t *testing.T) {
		var stringNil *int64
		var jsonNumber json.Number = "json number"

		require.Equal(t, ToString(Pointer("stringPointer")), Pointer("stringPointer"))
		require.Equal(t, ToString(Pointer(0.111)), Pointer("0.111"))
		require.Equal(t, ToString(Pointer(111)), Pointer("111"))
		require.Equal(t, ToString(Pointer(false)), Pointer("false"))
		require.Equal(t, ToString(Pointer(jsonNumber)), Pointer("json number"))
		require.Equal(t, ToString(stringNil), Pointer(""))
	})
}

func TestToInt64(t *testing.T) {
	t.Run("ShouldReturnInt64", func(t *testing.T) {
		var int64Nil *int64
		var jsonNumber json.Number = "25"

		require.Equal(t, ToInt64(Pointer("64")), Pointer[int64](64))
		require.Equal(t, ToInt64(Pointer(0.111)), Pointer[int64](0))
		require.Equal(t, ToInt64(Pointer(111)), Pointer[int64](111))
		require.Equal(t, ToInt64(Pointer(false)), Pointer[int64](0))
		require.Equal(t, ToInt64(Pointer(jsonNumber)), Pointer[int64](25))
		require.Equal(t, ToInt64(int64Nil), int64Nil)
	})
}

func TestEqualPointerValue(t *testing.T) {
	testEqualValueStr(t)
	testEqualValueInt64(t)
	testEqualValueFloat64(t)
	testEqualValueBool(t)
	testEqualValueTime(t)
}

func testEqualValueStr(t *testing.T) {
	testCases := map[string]struct {
		Param1         *string
		Param2         *string
		ExpectedOutput bool
	}{
		"ShouldReturnTrue_WhenInputIsNil": {
			Param1:         nil,
			Param2:         nil,
			ExpectedOutput: true,
		},
		"ShouldReturnTrue_WhenInputIsEmptyString": {
			Param1:         Pointer(""),
			Param2:         Pointer(""),
			ExpectedOutput: true,
		},
		"ShouldReturnFalse_WhenInputIsDifferent": {
			Param1:         Pointer("text1"),
			Param2:         Pointer("text2"),
			ExpectedOutput: false,
		},
		"ShouldReturnFalse_WhenOneOfTheInputIsNil": {
			Param1:         nil,
			Param2:         Pointer("text2"),
			ExpectedOutput: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := EqualPointerValue(testCase.Param1, testCase.Param2)
			require.Equal(t, testCase.ExpectedOutput, actual)
		})
	}
}

func testEqualValueInt64(t *testing.T) {
	testCases := map[string]struct {
		Param1         *int64
		Param2         *int64
		ExpectedOutput bool
	}{
		"ShouldReturnTrue_WhenInputIsNil": {
			Param1:         nil,
			Param2:         nil,
			ExpectedOutput: true,
		},
		"ShouldReturnTrue_WhenInputIs0": {
			Param1:         Pointer[int64](0),
			Param2:         Pointer[int64](0),
			ExpectedOutput: true,
		},
		"ShouldReturnFalse_WhenInputIsDifferent": {
			Param1:         Pointer[int64](64),
			Param2:         Pointer[int64](63),
			ExpectedOutput: false,
		},
		"ShouldReturnFalse_WhenOneOfTheInputIsNil": {
			Param1:         nil,
			Param2:         Pointer[int64](64),
			ExpectedOutput: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := EqualPointerValue(testCase.Param1, testCase.Param2)
			require.Equal(t, testCase.ExpectedOutput, actual)
		})
	}
}

func testEqualValueFloat64(t *testing.T) {
	testCases := map[string]struct {
		Param1         *float64
		Param2         *float64
		ExpectedOutput bool
	}{
		"ShouldReturnTrue_WhenInputIsNil": {
			Param1:         nil,
			Param2:         nil,
			ExpectedOutput: true,
		},
		"ShouldReturnTrue_WhenInputIs0": {
			Param1:         Pointer[float64](0),
			Param2:         Pointer[float64](0),
			ExpectedOutput: true,
		},
		"ShouldReturnFalse_WhenInputIsDifferent": {
			Param1:         Pointer(64.32),
			Param2:         Pointer(32.64),
			ExpectedOutput: false,
		},
		"ShouldReturnFalse_WhenOneOfTheInputIsNil": {
			Param1:         nil,
			Param2:         Pointer[float64](64),
			ExpectedOutput: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := EqualPointerValue(testCase.Param1, testCase.Param2)
			require.Equal(t, testCase.ExpectedOutput, actual)
		})
	}
}

func testEqualValueBool(t *testing.T) {
	testCases := map[string]struct {
		Param1         *bool
		Param2         *bool
		ExpectedOutput bool
	}{
		"ShouldReturnTrue_WhenInputIsNil": {
			Param1:         nil,
			Param2:         nil,
			ExpectedOutput: true,
		},
		"ShouldReturnTrue_WhenInputIsTrue": {
			Param1:         Pointer(true),
			Param2:         Pointer(true),
			ExpectedOutput: true,
		},
		"ShouldReturnFalse_WhenInputIsDifferent": {
			Param1:         Pointer(false),
			Param2:         Pointer(true),
			ExpectedOutput: false,
		},
		"ShouldReturnFalse_WhenOneOfTheInputIsNil": {
			Param1:         nil,
			Param2:         Pointer(true),
			ExpectedOutput: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := EqualPointerValue(testCase.Param1, testCase.Param2)
			require.Equal(t, testCase.ExpectedOutput, actual)
		})
	}
}

func testEqualValueTime(t *testing.T) {
	now := time.Now()
	testCases := map[string]struct {
		Param1         *time.Time
		Param2         *time.Time
		ExpectedOutput bool
	}{
		"ShouldReturnTrue_WhenInputIsNil": {
			Param1:         nil,
			Param2:         nil,
			ExpectedOutput: true,
		},
		"ShouldReturnTrue_WhenInputIsSame": {
			Param1:         Pointer(now),
			Param2:         Pointer(now),
			ExpectedOutput: true,
		},
		"ShouldReturnFalse_WhenInputIsDifferent": {
			Param1:         Pointer(now),
			Param2:         Pointer(time.Now()),
			ExpectedOutput: false,
		},
		"ShouldReturnFalse_WhenOneOfTheInputIsNil": {
			Param1:         nil,
			Param2:         Pointer(now),
			ExpectedOutput: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := EqualPointerValue(testCase.Param1, testCase.Param2)
			require.Equal(t, testCase.ExpectedOutput, actual)
		})
	}
}
