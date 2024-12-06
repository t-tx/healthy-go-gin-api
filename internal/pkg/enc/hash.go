package enc

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/big"

	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	saltByteSize   = 24
	hashByteSize   = 24
	minIterations  = 1000
	maxIterations  = 2000
	iterationIndex = 0
	saltIndex      = 1
	pbkdf2Index    = 2
	delimiter      = ":"
)

func Create(value string) (result string, err error) {
	salt, err := salt()
	if err != nil {
		return result, err
	}
	pbkdf2Iterations, err := randomIntInRange(minIterations, maxIterations)
	if err != nil {
		return result, err
	}
	hash := pbkdf2.Key([]byte(value), salt, pbkdf2Iterations, hashByteSize, sha1.New)
	result = strconv.Itoa(pbkdf2Iterations) + delimiter +
		base64.StdEncoding.EncodeToString(salt) + delimiter +
		base64.StdEncoding.EncodeToString(hash)

	return
}

func Check(value string, correctHash string) (result bool, err error) {
	split := strings.Split(correctHash, delimiter)
	iterations, err := strconv.Atoi(split[iterationIndex])
	if err != nil {
		return result, err
	}

	saltKey, _ := base64.StdEncoding.DecodeString(split[saltIndex])
	hash, _ := base64.StdEncoding.DecodeString(split[pbkdf2Index])
	testHash := pbkdf2.Key([]byte(value), saltKey, iterations, len(hash), sha1.New)
	result = bytes.Equal(hash, testHash)
	return
}

func salt() (result []byte, err error) {
	salt := make([]byte, saltByteSize)
	if _, err := rand.Read(salt); err != nil {
		return result, err
	}
	return salt, nil
}
func randomIntInRange(min, max int) (int, error) {
	if min > max {
		return 0, fmt.Errorf("min cannot be greater than max")
	}
	diff := max - min
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(diff+1)))
	if err != nil {
		return 0, err
	}
	return int(nBig.Int64()) + min, nil
}
