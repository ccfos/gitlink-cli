package auth

import (
	"os"
	"path/filepath"

	"github.com/zalando/go-keyring"
)

const (
	keyringService = "gitlink-cli"
	keyringUser    = "default"
)

func StoreToken(token string) error {
	err := keyring.Set(keyringService, keyringUser, token)
	if err != nil {
		// Fallback to file-based storage
		return storeTokenFile(token)
	}
	return nil
}

func LoadToken() (string, error) {
	token, err := keyring.Get(keyringService, keyringUser)
	if err != nil {
		// Fallback to file-based storage
		return loadTokenFile()
	}
	return token, nil
}

func DeleteToken() error {
	err := keyring.Delete(keyringService, keyringUser)
	if err != nil {
		return deleteTokenFile()
	}
	return nil
}

// File-based fallback

func credentialPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "gitlink-cli", "credentials")
}

func storeTokenFile(token string) error {
	p := credentialPath()
	if err := os.MkdirAll(filepath.Dir(p), 0700); err != nil {
		return err
	}
	return os.WriteFile(p, []byte(token), 0600)
}

func loadTokenFile() (string, error) {
	data, err := os.ReadFile(credentialPath())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func deleteTokenFile() error {
	return os.Remove(credentialPath())
}
