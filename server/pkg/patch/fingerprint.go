package patch

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
)

func Fingerprint(raw THRL6P) (string, error) {
	raw.Data.Tone.Global = Global{} // don't include global data in hash
	bytes, err := json.Marshal(raw.Data.Tone)
	if err != nil {
		return "", fmt.Errorf("failed to marshal tone data: %w", err)
	}
	return fmt.Sprintf("%x", sha1.Sum(bytes)), nil
}
