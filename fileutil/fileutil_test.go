package fileutil

import (
	"os"
	"testing"
)

func TestCreateFolder(t *testing.T) {

	rootFolder := "storage_test"

	tests := []struct {
		folderPath string
	}{
		{"test_folder_1"},
		{"test_folder_2/subfolder"},
		{"test_folder_3/subfolder/subsubfolder"},
	}

	for i := range tests {
		tests[i].folderPath = rootFolder + "/" + tests[i].folderPath
	}

	for _, test := range tests {
		// Clean up before test
		os.RemoveAll(test.folderPath)

		CreateFolder(test.folderPath)

		// Check if folder was created
		info, err := os.Stat(test.folderPath)
		if os.IsNotExist(err) {
			t.Errorf("expected folder %s to be created, but it does not exist", test.folderPath)
		}

		if !info.IsDir() {
			t.Errorf("expected %s to be a directory, but it is not", test.folderPath)
		}

		// Clean up after test
		os.RemoveAll(rootFolder)
	}
}
func TestDeleteFolder(t *testing.T) {
	rootFolder := "storage_test_delete"

	// Create a nested folder structure to delete
	folderPath := rootFolder + "/to_delete/subfolder"
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		t.Fatalf("failed to set up test folder: %v", err)
	}

	// Ensure the folder exists before deletion
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		t.Fatalf("expected folder %s to exist before deletion", folderPath)
	}

	// Call DeleteFolder on the root of the structure
	DeleteFolder(rootFolder + "/to_delete")

	// Check that the folder and its contents are deleted
	if _, err := os.Stat(rootFolder + "/to_delete"); !os.IsNotExist(err) {
		t.Errorf("expected folder %s to be deleted, but it still exists", rootFolder+"/to_delete")
	}

	// Clean up
	os.RemoveAll(rootFolder)
}
