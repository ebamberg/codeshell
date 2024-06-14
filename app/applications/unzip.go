package applications

import (
	"archive/zip"
	"codeshell/output"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func unzipSource(source, destination string, ignoreRootFolder bool, archive appSourceArchiveInfo) error {

	if archive.extractcommand != "" {
		return unzipWithCommand(source, destination, archive.extractcommand)
	}

	// 1. Open the zip file
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	// 2. Get the absolute destination path
	destination, err = filepath.Abs(destination)
	if err != nil {
		return err
	}

	// 3. Iterate over zip files inside the archive and unzip each of them
	for _, f := range reader.File {
		if ignoreRootFolder {
			pathElements := strings.Split(f.Name, "/") // zip file separator mkght not be the same as os.FileSeparator
			if len(pathElements) > 1 {
				f.Name = strings.Join(pathElements[1:], "/")
			} else {
				f.Name = ""
			}
		}
		if archive.rootfolder != "" {
			if strings.HasPrefix(f.Name, archive.rootfolder) {
				f.Name = strings.Replace(f.Name, archive.rootfolder, "", 1)
			}
		}
		if f.Name != "" {
			err := unzipFile(f, destination)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func unzipFile(f *zip.File, destination string) error {
	// 4. Check if file paths are not vulnerable to Zip Slip
	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	// 5. Create directory tree
	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// 6. Create a destination file for unzipped content
	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// 7. Unzip the content of a file and copy it to the destination file
	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}
	return nil
}

func unzipWithCommand(source string, destination string, extractcommand string) error {
	prompt := strings.ReplaceAll(extractcommand, "${source}", source)
	prompt = strings.ReplaceAll(prompt, "${targetfolder}", destination)
	cmdArgs := strings.Split(prompt, " ")
	exe := cmdArgs[0]
	args := cmdArgs[1:]
	cmd := exec.Command(exe, args...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		output.Println(string(out))
		return err
	} else {
		output.Println(string(out))
	}
	return nil
}
