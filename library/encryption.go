package library

import (
	"QAPI/logger"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/alexsasharegan/dotenv"
)

func Encode(data string) (string, error) {
	if err := dotenv.Load(); err != nil {
		logger.Log.Err(err).Msg("Gagal Load .env:")
		return "", err
	}

	key := []byte(os.Getenv("APP_KEY"))

	// Enkripsi menggunakan AES
	encrypted, err := _encryptAES([]byte(data), key)
	if err != nil {
		logger.Log.Err(err).Msg("Gagal Encrypt AES")
		return "", err
	}

	// Konversi hasil enkripsi menjadi base64
	fmt.Println("String awal:", data)
	fmt.Println("Hasil aes:", encrypted)
	result := base64.StdEncoding.EncodeToString(encrypted)
	fmt.Println("Hasil enkripsi:", result)
	return result, nil
}

func Decode(data string) (result *string) {
	if err := dotenv.Load(); err != nil {
		logger.Log.Err(err).Msg("Gagal Load .env:")
		result = nil
		return
	}

	// Dekripsi dari base64 ke byte array
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		logger.Log.Err(err).Msg("Gagal Decode Base64")
		result = nil
		return
	}

	key := []byte(os.Getenv("APP_KEY"))

	// Dekripsi menggunakan AES
	decrypted, err := _decryptAES(decoded, key)
	if err != nil {
		logger.Log.Err(err).Msg("Gagal Decrypt AES")
		result = nil
		return
	}

	*result = string(decrypted)
	return
}

func _encryptAES(plainText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Log.Err(err).Msg("Gagal Buat Cipher")
		return nil, err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		logger.Log.Err(err).Msg("Gagal Membaca Cipher Text")
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return cipherText, nil
}

func _decryptAES(cipherText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Log.Err(err).Msg("Gagal Buat Cipher")
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext is too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}
