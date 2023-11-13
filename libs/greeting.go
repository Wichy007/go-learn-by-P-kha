package libs

func Login(username string, password string) bool {
	if username == "wichy" && password == "password" {
		return true
	}
	return false
}

func GreetingText() string {
	return "welcome"
}
