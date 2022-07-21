package exhashes

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func Main() {
	s := "sha256 this string"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	s1 := "string to hash256"
	h1 := sha256.New()
	h1.Write([]byte(s1))
	fmt.Println(s1)
	fmt.Println(h1.Sum(nil))
	s2 := "string to hash512"
	h2 := sha512.New()
	h2.Write([]byte(s2))
	fmt.Println(s2)
	fmt.Println(h2.Sum(nil))

	s3 := "Закрепляем хэширование"
	h3 := sha256.New()
	h3.Write([]byte(s3))
	result := h3.Sum(nil)
	fmt.Println(result)
}
