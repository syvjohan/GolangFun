package checksum

import (
		"crypto/md5"
		"fmt"
		"encoding/hex"
)

func MD5(text string) (result string){
	hash := md5.New()
	
	hash.Write([]byte(text))
	result = hex.EncodeToString(hash.Sum(nil))

	return result
}
