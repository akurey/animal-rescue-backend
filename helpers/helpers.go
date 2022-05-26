package helpers

// manage error
func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
