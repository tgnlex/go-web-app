package utils

func LoadEnv() {
	err := env
	if err != nil {
		logger("Error Loading .env file.")
	}
}
