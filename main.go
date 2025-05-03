/*
 * grafana2hashcat - Convert Grafana-style hashes to Hashcat format
 *
 * This tool reads hashes and salts from a file (formatted as "hash:salt")
 * and converts them to Hashcat's sha256 format with base64-encoded values:
 * "sha256:10000:<base64_salt>:<base64_hash>"
 *
 * Usage: ./grafana2hashcat <input_file>
 * Input format: <hex_hash>:<salt> (one per line)
 * Output: Creates 'out_hashes.txt' with converted hashes
 */

package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check for input file argument
	if len(os.Args) < 2 {
		fmt.Println("[+] Usage: ./grafana2hashcat <hash_file>")
		fmt.Println("\n[+] Input file format:")
		fmt.Println("  Each line must contain: <hex_hash>:<salt>")
		fmt.Println("  Example: 2ac9cb7dc02b3c0083eb70898e549b63:12345")
		fmt.Println("\n[+] Output:")
		fmt.Println("  Creates 'out_hashes.txt' with converted hashes in Hashcat format")
		return
	}

	inputFile := os.Args[1]
	outputFile := "out_hashes.txt"

	// Open input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("[-] Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Creating output file
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("[-] Error creating output file: %v\n", err)
		return
	}
	defer output.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	successCount := 0

	// Processing each line of input file
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineCount++

		// Skip empty lines
		if line == "" {
			continue
		}

		// Split hash and salt
		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			fmt.Printf("[-] Invalid format line %d: %s\n", lineCount, line)
			fmt.Printf("[-] Required format: <hex_hash>:<salt> (separated by colon)\n")
			fmt.Printf("[-] Example: 2ac9cb7dc02b3c0083eb70898e549b63:12345\n")
			continue
		}

		hashHex := parts[0]
		salt := parts[1]

		// Decoding hex hash to bytes
		decodedHash, err := hex.DecodeString(hashHex)
		if err != nil {
			fmt.Printf("[-] Error decoding hash (line %d): %v\n", lineCount, err)
			fmt.Printf("[-] Hash must be valid hexadecimal\n")
			continue
		}

		// Encoding to Base64
		hashB64 := base64.StdEncoding.EncodeToString(decodedHash)
		saltB64 := base64.StdEncoding.EncodeToString([]byte(salt))

		// Writing on Hashcat format
		_, err = output.WriteString(fmt.Sprintf("sha256:10000:%s:%s\n", saltB64, hashB64))
		if err != nil {
			fmt.Printf("[-] Error writing output (line %d): %v\n", lineCount, err)
			continue
		}

		successCount++
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("[-] File reading error: %v\n", err)
		return
	}

	// Print summary
	fmt.Printf("\n[+] Conversion complete:\n")
	fmt.Printf("    Total lines processed: %d\n", lineCount)
	fmt.Printf("    Successfully converted: %d\n", successCount)
	fmt.Printf("    Failed conversions: %d\n", lineCount-successCount)
	fmt.Printf("\nYou can use the #10900 mode in hashcat!\n")
  	fmt.Printf("\nExample: hashcat -m 10900 out_hashes.txt -a 0 /usr/share/wordlists/rockyou.txt -O\n")
	fmt.Printf("[+] Results saved to: %s\n", outputFile)
}
