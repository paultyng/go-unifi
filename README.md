# Unifi Go SDK [![GoDoc](https://godoc.org/github.com/paultyng/go-unifi?status.svg)](https://godoc.org/github.com/paultyng/go-unifi)

This was written primarily for use in my [Terraform provider for Unifi](https://github.com/paultyng/terraform-provider-unifi).

## Versioning

Many of the naming adjustments are breaking changes, but to simplify things, treating naming errors as minor changes for the 1.0.0 version (probably should have just started at 0.1.0).

## Note on Code Generation

The data models and basic REST methods are "generated" from JSON files in the JAR that show all fields and the associated regex/validation information.

To regenerate the code, you can bump the Unifi Controller version number in [unifi/gen.go] and run `go generate` inside the `unifi` directory.

This code generation is kind of gross, I wanted to switch to using the java classes in the jar like scala2go but the jar is obfuscated and I couldn't find a way to extract that information from anywhere else. Maybe it exists somewhere in the web UI, but I was unable to find it in there in a way that was extractable in a practical way.

Still planning to dig through the bits some more later on.
