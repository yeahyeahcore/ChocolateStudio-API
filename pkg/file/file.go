package file

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
)

type File struct {
	GinContext     *gin.Context
	PhotographyDir string
	PhotoBookDir   string
}

func NewFile(directoryOfPhotography, directoryOfPhotobooks string) *File {
	return &File{
		PhotographyDir: directoryOfPhotography,
		PhotoBookDir:   directoryOfPhotobooks,
	}
}


func (f File) SaveFilesToFolder(directory string, files []*multipart.FileHeader) {
	//Create a folder/directory at a full qualified path
	err := os.MkdirAll(directory, 0755)
	if err != nil {
		fmt.Println(err)
	}

	//перебирает все файлы и сохраняет каждый в директорию
	for _, file := range files {
		err := f.GinContext.SaveUploadedFile(file, fmt.Sprintf("%s/%s", directory, file.Filename))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (f File) RemoveFilesFromFolder(directory string) {
	//удаление директории с файлами
	err := os.RemoveAll(directory)
	if err != nil {
		fmt.Println(err)
	}
}

func (f File) ReplaceFiles(directory string, oldFileNames []string, files []*multipart.FileHeader) {
	//перебор всех имён файлов, которые нужно удалить
	for _, name := range oldFileNames {
		err := os.Remove(fmt.Sprintf("%s/%s", directory, name))
		if err != nil {
			fmt.Println(err)
		}
	}

	//перебор всех файлов, которые необходимо сохранить
	for _, file := range files {
		err := f.GinContext.SaveUploadedFile(file, directory)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (f File) RenameFolder(path, oldFolderName, newFolderName string) {
	var oldDirectory = fmt.Sprintf("%s/%s", path, oldFolderName)
	var newDirectory = fmt.Sprintf("%s/%s", path, newFolderName)
	err := os.Rename(oldDirectory, newDirectory)
	if err != nil {
		fmt.Println(err)
	}
}

