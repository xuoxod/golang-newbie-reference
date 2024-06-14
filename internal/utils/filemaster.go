package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

// createDirIfNotExist creates a directory if it does not exist
func CreateDirIfNotExist(path string) error {
	if PathIsDir(path) {
		const mode = 0755
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.Mkdir(path, mode)
			if err != nil {
				log.Println(err)
				return err
			}
		}
		return nil
	}
	// return errors.New(fmt.Sprintf("Path '%s' is not a directory\n", path))
	return fmt.Errorf("Path '%v' is not a direcory", path)
}

func PathExists(filepath string) (map[string]interface{}, bool) {
	fi, err := os.Stat(filepath)

	if err != nil {
		log.Printf("os.stat error:\t%s\n", err.Error())
		return nil, false
	}

	fileInfo := make(map[string]interface{})
	fileInfo["name"] = fi.Name()
	fileInfo["modified"] = fi.ModTime()
	fileInfo["mode"] = fi.Mode()
	fileInfo["size"] = fi.Size()
	fileInfo["source"] = fi.Sys()
	fileInfo["isdir"] = fi.IsDir()
	fileInfo["extension"] = path.Ext(filepath)
	fileInfo["base"] = path.Base(filepath)
	fileInfo["parent"] = path.Dir(filepath)
	fileInfo["isabs"] = path.IsAbs(filepath)

	return fileInfo, true
}

func PathIsFile(filepath string) bool {
	fi, err := os.Stat(filepath)

	if err != nil {
		log.Printf("os.stat error:\t%s\n", err.Error())
		return false
	}

	return !fi.IsDir()
}

func PathIsDir(filepath string) bool {
	fi, err := os.Stat(filepath)

	if err != nil {
		log.Printf("os.stat error:\t%s\n", err.Error())
		return false
	}

	return fi.IsDir()
}

func ReadFileToConsole(filepath string) error {
	if PathIsFile(filepath) {
		var mode os.FileMode = 0700
		// file, err := os.Open(filepath)
		file, err := os.OpenFile(filepath, os.O_RDONLY, mode)

		if err != nil {
			fmt.Println("Error opening file:\t", err.Error())
			return err
		}

		defer file.Close()

		const maxSz = 4096

		// make a buffer
		buf := make([]byte, maxSz)

		for {
			readTotal, err := file.Read(buf)

			if err != nil {
				if err != io.EOF {
					fmt.Println(err.Error())
					return err
				}
				break
			}
			fmt.Println(string(buf[:readTotal]))
		}
		return nil
	}
	// return errors.New(fmt.Sprintf("Path '%v' is not a file\n", filepath))
	return fmt.Errorf("Path '%v' is not a file", filepath)
}

func ReadFileToList(filepath string) []string {
	var mode os.FileMode = 0700
	// file, err := os.Open(filepath)
	file, err := os.OpenFile(filepath, os.O_RDONLY, mode)

	if err != nil {
		fmt.Println("Error opening file:\t", err.Error())
	}

	defer file.Close()

	content := []string{}

	const maxSz = 4096

	// make a buffer
	buf := make([]byte, maxSz)

	for {
		readTotal, err := file.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
			}
			break
		}
		content = append(content, string(buf[:readTotal]))
	}
	return content
}
