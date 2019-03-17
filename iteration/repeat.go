package iteration

func Repeat(arg string, repeatCount int) string{
	result :=""
	for i:=0; i<repeatCount;i++  {
		result += arg
	}
	return result
}