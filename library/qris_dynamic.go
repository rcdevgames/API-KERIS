package library

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/skip2/go-qrcode"
)

func ConvertQris(qris_data string, amount int64) (result string) {
	qris := qris_data[:len(qris_data)-4]
	step1 := strings.Replace(qris, "010211", "010212", 1)
	step2 := strings.Split(step1, "5802ID")
	uang := fmt.Sprintf("54%02d%d5802ID", len(fmt.Sprintf("%d", amount)), amount)

	fix := strings.TrimSpace(step2[0]) + uang + strings.TrimSpace(step2[1])
	fix += _convertCRC16(fix)

	fmt.Printf("\n[+] Result: %s\n", fix)

	// Generate QR code
	qrCode, err := qrcode.Encode(fix, qrcode.High, 256)
	if err != nil {
		fmt.Println("Failed to generate QR code:", err)
		return
	}

	// Convert QR code to Base64
	result = base64.StdEncoding.EncodeToString(qrCode)
	return
}

func _convertCRC16(str string) string {
	charCodeAt := func(str string, i int) byte {
		return str[i]
	}
	crc := 0xFFFF
	strlen := len(str)
	for c := 0; c < strlen; c++ {
		crc ^= int(charCodeAt(str, c)) << 8
		for i := 0; i < 8; i++ {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc = crc << 1
			}
		}
	}
	hex := crc & 0xFFFF
	hexStr := fmt.Sprintf("%X", hex)
	if len(hexStr) == 3 {
		hexStr = "0" + hexStr
	}
	return hexStr
}
