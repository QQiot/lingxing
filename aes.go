package lingxing

import (
	"bytes"
	"crypto/aes"
)

type AesTool struct {
	key       []byte
	blockSize int
}

func NewAesTool(key []byte, blockSize int) AesTool {
	return AesTool{blockSize: blockSize, key: key}
}

func (t *AesTool) padding(src []byte) []byte {
	padding := t.blockSize - len(src)%t.blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func (t *AesTool) unPadding(src []byte) []byte {
	length := len(src)
	up := int(src[length-1])
	return src[:(length - up)]
}

func (t *AesTool) ECBEncrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher(t.key)
	if err != nil {
		return nil, err
	}

	src = t.padding(src)
	encryptData := make([]byte, len(src))

	for bs, be := 0, block.BlockSize(); bs < len(src); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Encrypt(encryptData[bs:be], src[bs:be])
	}
	return encryptData, nil
}

func (t *AesTool) ECBDecrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher(t.key)
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(src))
	size := block.BlockSize()
	for bs, be := 0, size; bs < len(src); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], src[bs:be])
	}
	return t.unPadding(decrypted), nil
}
