# Grafana2Hashcat

A conversion tool to transform Grafana-style password hashes into Hashcat-compatible format for `Hash Cracking`.

[![Go Report Card](https://goreportcard.com/badge/github.com/incommatose/grafana2hashcat)](https://goreportcard.com/report/github.com/incommatose/grafana2hashcat)

## Features

- Converts Grafana's SHA256 hashes to Hashcat format (`sha256:10000:<salt>:<hash>`)
- Handles hexadecimal hash decoding and Base64 encoding
- Batch processing of multiple hashes
- Lightweight and fast (no external dependencies)

## Installation

### Prerequisites
- Go 1.20+ (for building from source)
- Hashed password
- Salt

### Build from source
You can build the tool yourself by performing the following steps

```bash
# Clone the repository
git clone https://github.com/incommatose/grafana2hashcat.git
cd grafana2hashcat

# Build binary
go build -o grafana2hashcat -ldflags="-s -w"
upx grafana2hashcat

# Make executable available system-wide (optional)
sudo mv grafana2hashcat /usr/local/bin/
```


## Usage
You can see a help panel if you run the script

~~~ bash
./grafana2hashcat                         
[+] Usage: ./grafana2hashcat <hash_file>

[+] Input file format:
  Each line must contain: <hex_hash>:<salt>
  Example: 2ac9cb7dc02b3c0083eb70898e549b63:12345

[+] Output:
  Creates 'out_hashes.txt' with converted hashes in Hashcat format
~~~

### Creating Hashed Passwords File
Save hashed passwords in a file as follows

~~~ text
2ac9cb7dc02b3c0083eb70898e549b63:12345
441a715bd788e928170be7954b17cb1a:67890
~~~

This is an example of use with a list of hashes with the suggested format in `example_hashes.txt`

```bash
grafana2hashcat ../../content/hashes.txt 

[+] Conversion complete:
    Total lines processed: 1
    Successfully converted: 1
    Failed conversions: 0

You can use the #10900 mode in hashcat!

Example: hashcat -m 10900 out_hashes.txt -a 0 /usr/share/wordlists/rockyou.txt -O
[+] Results saved to: out_hashes.txt
```


