package gonet_utils

import (
	"math/rand"
	"os"
)

func RandString(n int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTVXYZ" +
		"abcdefghijklmnopqrstvxyz" +
		"0123456789"

	ret := ""
	for i := 0; i < n; i++ {
		ind := rand.Int31n(int32(len(charset)))
		ret = ret + string(charset[ind])
	}

	return ret
}

func PathExists(p string) (exists bool, err error) {
	_, err = os.Stat(p)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
