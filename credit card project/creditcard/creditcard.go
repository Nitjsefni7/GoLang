package creditcard

import ( 
	"fmt"
	"strings"
	"errors"
	"strconv"
)
func TakeCCNumber() ([]int, error) {
	fmt.Println("Enter Your Credit Card number:")
	var ccNum string

	fmt.Scanln(&ccNum)
	digitsCC, err := convertUserInput(ccNum)
	return digitsCC, err
}

func convertUserInput(s string) ([]int, error) {
	sSlice := strings.Split(s, "")
	if len(sSlice) != 16 {
		return nil, errors.New("You have to type 16digit cc number, without spaces")
	}
	digitSlice := []int{}
	for _, v := range sSlice {
		if digit, err := strconv.Atoi(v); err != nil {
			return nil, errors.New("You have to type only digits")
		} else {
			digitSlice = append(digitSlice, digit)
		}
	}

	return digitSlice, nil
}

func LuhnSumAlg(digits []int) bool {
	sum := 0
	isSecond := false
	for i := len(digits)-1; i >= 0; i-- {
		if isSecond {
			digits[i] *= 2
		}
		sum += digits[i] / 10
		sum += digits[i] % 10
		isSecond = !isSecond
	}

	return (sum % 10 == 0)
}