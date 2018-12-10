package manager

import (
	"context"
	"crypto/sha256"
	"fmt"
)

var salt = "strong#"

func GenPasswordHash(ctx context.Context, password string) string {
	sum := sha256.Sum256([]byte(password + salt))
	return fmt.Sprintf("%x", sum)
}
