package helper

func PanicIE(err error) {
	if err != nil {
		panic(err)
	}
}
