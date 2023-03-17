package obfuscator

type Obfuscator interface {
	Obfuscate(code string) (string, error)
}
