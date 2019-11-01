package libraries

import "crypto/sha1"
import "fmt"

func MakeSHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	sha := fmt.Sprintf("%x", bs)

	return sha
}
