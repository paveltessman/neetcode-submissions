import (
	"encoding/binary"
)


type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	buf := make([]byte, 0)

	for _, word := range strs {
		buf = binary.BigEndian.AppendUint32(buf, uint32(len(word)))
		buf = append(buf, []byte(word)...)
	}

	return string(buf)
}

func (s *Solution) Decode(encoded string) []string {
	encodedBytes := []byte(encoded)
	result := make([]string, 0)

	var i uint32
	for i < uint32(len(encodedBytes)) {

		length := binary.BigEndian.Uint32(encodedBytes[i : i+4])
		i += 4

		word := encodedBytes[i : i+length]
		result = append(result, string(word))
		i = i + length

	}
	return result

}
