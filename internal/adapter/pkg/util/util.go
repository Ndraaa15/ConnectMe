package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/rs/zerolog/log"
)

const numberCharset = "0123456789"

func GenerateCode(length int) string {
	code := make([]byte, length)
	for i := range code {
		n, err := rand.Int(rand.Reader, big.NewInt(9))
		if err != nil {
			log.Error().Err(err).Msg("Failed to generate code")
		}

		code[i] = numberCharset[n.Int64()]
	}
	return string(code)
}

func GetTimeSinceCreation(createdAt time.Time) string {
	duration := time.Since(createdAt)

	seconds := int(duration.Seconds())
	minutes := int(duration.Minutes())
	hours := int(duration.Hours())
	days := int(duration.Hours() / 24)
	months := int(duration.Hours() / (24 * 30))
	years := int(duration.Hours() / (24 * 365))

	switch {
	case seconds < 60:
		return fmt.Sprintf("%d detik lalu", seconds)
	case minutes < 60:
		return fmt.Sprintf("%d menit lalu", minutes)
	case hours < 24:
		return fmt.Sprintf("%d jam lalu", hours)
	case days < 30:
		return fmt.Sprintf("%d hari lalu", days)
	case months < 12:
		return fmt.Sprintf("%d bulan lalu", months)
	default:
		return fmt.Sprintf("%d tahun lalu", years)
	}
}
