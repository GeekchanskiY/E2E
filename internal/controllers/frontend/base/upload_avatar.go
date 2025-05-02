package base

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"

	"finworker/internal/models"
)

func (c *controller) UploadAvatar(ctx context.Context, userID int64, fileHeader *multipart.FileHeader, file []byte) error {
	var (
		dst                        *os.File
		newFileName                string
		newFilePath                string
		fileInfo                   os.FileInfo
		previousFileName           string
		fileNameGenerationAttempts uint

		err error
	)

	if _, err = os.Stat(mediaPath); os.IsNotExist(err) {
		if err = os.Mkdir(mediaPath, 0755); err != nil {
			return err
		}
	}

	if _, err = os.Stat(filepath.Join(mediaPath, avatarPath)); os.IsNotExist(err) {
		if err = os.Mkdir(filepath.Join(mediaPath, avatarPath), 0755); err != nil {
			return err
		}
	}

	fileNameGenerationAttempts = 0
	for {

		newFileName = generateUTF8FileName(fileHeader)
		newFilePath = filepath.Join(mediaPath, avatarPath, newFileName)

		// Avoid overwriting an existing file
		if fileInfo, err = os.Stat(newFilePath); os.IsNotExist(err) {
			break
		}
		fmt.Println(fileInfo)
		fileNameGenerationAttempts += 1
		if fileNameGenerationAttempts > 100 {
			c.logger.Error("media.upload.controller", zap.String("event", "failed to generate a new file name"), zap.Uint("attempts", fileNameGenerationAttempts))

			_, err = c.registryRepo.Push(ctx, &models.Event{
				Name:    "error",
				Event:   models.RegistryLog,
				Content: "failed to generate avatar filename",
			})

			return errors.New("failed to generate avatar filename, too many attempts")
		}

		c.logger.Debug("media.upload.controller", zap.String("event", "file already exists, generating new name"), zap.Uint("attempts", fileNameGenerationAttempts))
	}

	// Create a new file
	if dst, err = os.Create(newFilePath); err != nil {
		return err
	}

	defer func() {
		if err = dst.Close(); err != nil {
			c.logger.Fatal("media.upload.controller failed to close dst", zap.Error(err))
		}
	}()

	if _, err = dst.Write(file); err != nil {
		return err
	}

	// Save avatar and delete previous
	if previousFileName, err = c.userRepo.UpdateAvatar(ctx, userID, newFilePath); err != nil {
		return err
	}

	if previousFileName != "" {
		if err = os.Remove(previousFileName); err != nil {
			c.logger.Error("media.upload.controller failed to remove previous avatar", zap.Error(err))
			_, err = c.registryRepo.Push(ctx, &models.Event{
				Name:    "error",
				Event:   models.RegistryLog,
				Content: fmt.Sprintf("failed to delete [%s]", previousFileName),
			})
		}
	}

	return nil
}

func generateUTF8FileName(fileHeader *multipart.FileHeader) string {
	var (
		runePool      []rune
		randomRunes   []rune
		fileExtension string
		idx           *big.Int
		err           error
	)
	runePool = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	if lastDot := strings.LastIndex(fileHeader.Filename, "."); lastDot != -1 {
		fileExtension = fileHeader.Filename[lastDot:]
	} else {
		return "error"
	}

	randomRunes = make([]rune, 255-len(fileExtension)-len(mediaPath)-len(avatarPath)-2)
	for i := range randomRunes {
		if idx, err = rand.Int(rand.Reader, big.NewInt(int64(len(runePool)))); err != nil {
			idx = big.NewInt(0)
		}

		randomRunes[i] = runePool[idx.Int64()]
	}

	newFileName := string(randomRunes) + fileExtension

	return newFileName
}
