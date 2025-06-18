package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"math/big"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// PickPDF opens a file dialog to select a PDF and returns the selected path
func (a *App) PickPDF() (string, error) {
	options := runtime.OpenDialogOptions{
		Title:   "Choose a PDF file to encrypt",
		Filters: []runtime.FileFilter{{DisplayName: "PDF Files", Pattern: "*.pdf"}},
	}
	path, err := runtime.OpenFileDialog(a.ctx, options)
	if err != nil {
		return "", err
	}
	return path, nil
}

// Generate randowm password for Owner Password
func generateRandomPassword() string {
	length, _ := rand.Int(rand.Reader, big.NewInt(16))
	n := 16 + int(length.Int64())

	b := make([]byte, n)
	_, _ = rand.Read(b) // ignore error on purpose

	return base64.RawURLEncoding.EncodeToString(b)
}

// EncryptPDF encrypts a PDF file with the given password and returns the output file path or an error
func (a *App) EncryptPDF(inputPath string, password string, permInt int) (string, error) {
	outputPath := inputPath[:len(inputPath)-4] + "_encrypted.pdf"

	conf := api.LoadConfiguration()
	conf.UserPW = password
	conf.OwnerPW = generateRandomPassword()
	conf.EncryptUsingAES = true
	conf.EncryptKeyLength = 256
	conf.Permissions = model.PermissionFlags(permInt)

	err := api.EncryptFile(inputPath, outputPath, conf)
	if err != nil {
		return "", err
	}

	return outputPath, nil
}

func (a *App) GetPDFPermissions(filePath string, password string) (int, error) {
	conf := api.LoadConfiguration()
	conf.OwnerPW = password
	conf.UserPW = password

	perm, err := api.GetPermissionsFile(filePath, conf)
	if err != nil {
		return 0, err
	}
	return int(*perm), nil
}

// DecryptPDF decrypts a PDF file using the given password
func (a *App) DecryptPDF(inputPath string, password string) (string, error) {
	outputPath := inputPath[:len(inputPath)-4] + "_decrypted.pdf"

	conf := api.LoadConfiguration()
	conf.OwnerPW = password
	conf.UserPW = password

	err := api.DecryptFile(inputPath, outputPath, conf)
	if err != nil {
		return "", err
	}

	return outputPath, nil
}
