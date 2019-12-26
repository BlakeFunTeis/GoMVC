package libraries

import "golang.org/x/crypto/bcrypt"

func MakeBcrypt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		//fmt.Println(err)
	}
	//fmt.Println(hash)

	return string(hash)
}

func IsCheckBcrypt(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return true
	}

	return false
}