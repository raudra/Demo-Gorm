package config

func RaiseException(err error) {
	if err != nil {
		panic(err.Error())
	}
}