package utils

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// What is the file slice sent from the front end and the current slice?
// After getting it, the backend compares whether the sub-slices are uploaded or whether they are incomplete slices.
//How big is each piece sent by the front end?
// The front end informs whether it is the last piece and whether it is completed.

const (
	breakpointDir = "./breakpointDir/"
	finishDir     = "./fileDir/"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: BreakPointContinue
//@description: resume upload after breakpoint
//@param: content []byte, fileName string, contentNumber int, contentTotal int, fileMd5 string
//@return: error, string

func BreakPointContinue(content []byte, fileName string, contentNumber int, contentTotal int, fileMd5 string) (string, error) {
	path := breakpointDir + fileMd5 + "/"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return path, err
	}
	pathC, err := makeFileContent(content, fileName, path, contentNumber)
	return pathC, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CheckMd5
//@description: Check Md5
//@param: content []byte, chunkMd5 string
//@return: CanUpload bool

func CheckMd5(content []byte, chunkMd5 string) (CanUpload bool) {
	fileMd5 := MD5V(content)
	if fileMd5 == chunkMd5 {
		return true // You can continue uploading
	} else {
		return false // The slice is incomplete and discarded
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: makeFileContent
//@description: Create slice content
//@param: content []byte, fileName string, FileDir string, contentNumber int
//@return: string, error

func makeFileContent(content []byte, fileName string, FileDir string, contentNumber int) (string, error) {
	if strings.Index(fileName, "..") > -1 || strings.Index(FileDir, "..") > -1 {
		return "", errors.New("The file name or path is illegal")
	}
	path := FileDir + fileName + "_" + strconv.Itoa(contentNumber)
	f, err := os.Create(path)
	if err != nil {
		return path, err
	} else {
		_, err = f.Write(content)
		if err != nil {
			return path, err
		}
	}
	defer f.Close()
	return path, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: makeFileContent
//@description: Create slice file
//@param: fileName string, FileMd5 string
//@return: error, string

func MakeFile(fileName string, FileMd5 string) (string, error) {
	rd, err := os.ReadDir(breakpointDir + FileMd5)
	if err != nil {
		return finishDir + fileName, err
	}
	_ = os.MkdirAll(finishDir, os.ModePerm)
	fd, err := os.OpenFile(finishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return finishDir + fileName, err
	}
	defer fd.Close()
	for k := range rd {
		content, _ := os.ReadFile(breakpointDir + FileMd5 + "/" + fileName + "_" + strconv.Itoa(k))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(finishDir + fileName)
			return finishDir + fileName, err
		}
	}
	return finishDir + fileName, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: RemoveChunk
//@description: Remove slice
//@param: FileMd5 string
//@return: error

func RemoveChunk(FileMd5 string) error {
	err := os.RemoveAll(breakpointDir + FileMd5)
	return err
}
