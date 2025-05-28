package fileutil

import (
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func CreateFolder(filePath string) {
	err := os.MkdirAll(filePath, 0755)
	if err != nil {
		log.Fatalf("Failed to create folder: %v", err)
	}
	log.Debugf("Created folder at path: %s", filePath)
}

func DeleteFolder(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Debugf("Folder does not exist at path: %s, skipping delete", filePath)
		return
	}
	err := os.RemoveAll(filePath)
	if err != nil {
		log.Fatalf("Failed to delete folder: %v", err)
	}
	log.Debugf("Deleted folder at path: %s", filePath)
}

func SaveUrlToFile(url string, filePath string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to download URL: %v", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("Failed to save URL to file: %v", err)
	}

	log.Debugf("Successfully saved URL %s to file %s", url, filePath)
}

func ConvertWebm(srcFile string, dstFile string) {
	// FFmpeg command to extract the first frame and save it as PNG
	cmd := exec.Command("ffmpeg", "-i", srcFile, "-vf", "select=eq(n\\,0)", "-vframes", "1", dstFile)

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Fatalf("FFmpeg command failed: %v", err)
	}

	log.Debugf("Successfully converted %s to %s", srcFile, dstFile)
}
