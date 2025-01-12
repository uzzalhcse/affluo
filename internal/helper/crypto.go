package helper

import (
	"affluo/constant"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// encrypt takes a string and returns a base64 encoded encrypted string
func Encrypt(text string) (string, error) {
	block, err := aes.NewCipher(constant.EncryptionKey)
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// decrypt takes a base64 encoded encrypted string and returns the original string
func Decrypt(cryptoText string) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(constant.EncryptionKey)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

func DecodeTrackingCode(encryptedCode string) (publisherID int64, targetType string, targetValue string, expiredAt time.Time, err error) {
	decrypted, err := Decrypt(encryptedCode)
	if err != nil {
		return 0, "", "", time.Time{}, fmt.Errorf("failed to decrypt tracking code: %v", err)
	}

	// Split by underscore to separate publisher, target, and expiration parts
	parts := strings.Split(decrypted, "_")
	if len(parts) != 3 {
		return 0, "", "", time.Time{}, fmt.Errorf("invalid tracking code format")
	}

	// Extract publisher ID from "publisher-{id}"
	pubParts := strings.Split(parts[0], "-")
	if len(pubParts) != 2 || pubParts[0] != "publisher" {
		return 0, "", "", time.Time{}, fmt.Errorf("invalid publisher format")
	}
	publisherID, err = strconv.ParseInt(pubParts[1], 10, 64)
	if err != nil {
		return 0, "", "", time.Time{}, fmt.Errorf("invalid publisher ID: %v", err)
	}

	// Extract target type and value from second part
	targetParts := strings.Split(parts[1], "-")
	if len(targetParts) != 2 {
		return 0, "", "", time.Time{}, fmt.Errorf("invalid target format")
	}
	targetType = targetParts[0]  // "banner" or "gig"
	targetValue = targetParts[1] // could be numeric ID or string

	// Extract expiration time from "expired-{timestamp}"
	expiredParts := strings.Split(parts[2], "-")
	if len(expiredParts) != 2 || expiredParts[0] != "expired" {
		return 0, "", "", time.Time{}, fmt.Errorf("invalid expiration format %s", expiredParts)
	}
	expiredAt, err = time.Parse("20060102150405", expiredParts[1])
	if err != nil {
		return 0, "", "", time.Time{}, fmt.Errorf("invalid expiration time: %v", err)
	}

	return publisherID, targetType, targetValue, expiredAt, nil
}

// Helper function to parse target value as int64 if needed
func ParseTargetID(targetValue string) (int64, error) {
	return strconv.ParseInt(targetValue, 10, 64)
}
