package unifi

// This will generate the *.generated.go files in this package for the specified
// Unifi controller version.
//go:generate go run ../fields/ -version-base-dir=../fields/ 6.0.43
//go:generate gofmt -w -s ./
