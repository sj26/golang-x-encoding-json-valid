package json

func IsValid(data []byte) bool {
	err := checkValid(data, &scanner{})
	return err == nil
}
