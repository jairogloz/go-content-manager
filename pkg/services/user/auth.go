package user

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/jairogloz/go-content-manager/pkg/domain"
)

func (s *Service) Auth(apiKey string) (*domain.User, error) {

	hashedKey := hashText(apiKey)

	user, err := s.Repo.FindByAPIKey(hashedKey)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func hashText(text string) string {
	// Crear un nuevo hasher SHA-256
	hasher := sha256.New()

	// Escribir el texto en el hasher
	hasher.Write([]byte(text))

	// Obtener el hash en formato byte slice
	hashBytes := hasher.Sum(nil)

	// Convertir el hash a una cadena hexadecimal
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}
