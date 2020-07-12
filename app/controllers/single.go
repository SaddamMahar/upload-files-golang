package controllers

import (
	"bytes"
	"github.com/revel/examples/upload/app/connection"
	"github.com/revel/examples/upload/app/models"
	"github.com/revel/revel"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

type Single struct {
	App
}

func (c *Single) Upload() revel.Result {
	files := c.ListFiles()
	return c.Render(files)
}

func (c *Single) HandleUpload(file []byte) revel.Result {
	// Handle errors.
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect((*Single).Upload)
	}

	c.UploadFile(c.Params.Files["file"][0].Filename, len(file), file)
	//return c.Render(files)
	//return c.Redirect(routes.Hotels.Index())
	return c.Redirect((*Single).Upload)
}

func (c *Single) Show(name string) revel.Result {
	con := connection.Connection

	//reading
	re, err := con.Retr(name)
	if err != nil {
		panic("Error in retriving file(" + name + "), error : " + err.Error())
	}
	defer re.Close()

	//create new from reader
	file, err := os.Create(name)
	if err != nil {
		panic("Error in creating temporary file(" + name + "), error : " + err.Error())
	}
	_, err = io.Copy(file, re)
	if err != nil {
		panic("Error in copying content to temporary file(" + name + "), error : " + err.Error())
	}
	return c.RenderFile(file, "attachment")
}

func (c *Single) UploadFile(fileName string, size int, file []byte) {
	con := connection.Connection

	dir := "/golang"
	CheckCurrentDir(dir)
	err := con.Stor(fileName, bytes.NewReader(file))
	if err != nil {
		panic("Error in storing file, error : " + err.Error())
	}
}

func (c *Single) ListFiles() []models.FilesList {
	con := connection.Connection
	dir := "/golang"

	CheckCurrentDir(dir)
	erntries, err := con.List(dir)
	if err != nil {
		panic("Error in getting list of files from dir: " + dir + ", error : " + err.Error())
	}
	filesList := make([]models.FilesList, 0)
	for _, v := range erntries {
		var fileInfo models.FilesList
		fileInfo.Name = v.Name
		fileInfo.Size = v.Size
		fileInfo.Created = v.Time
		filesList = append(filesList, fileInfo)
	}
	return filesList
}

func CheckCurrentDir(dir string) {
	con := connection.Connection
	dirCurrent := ""
	var err error
	if dirCurrent, err = con.CurrentDir(); err != nil {
		panic("Error in getting current path , error : " + err.Error())
	}
	if dirCurrent != dir {
		err = con.ChangeDir(dir)
		if err != nil {
			// if dir not present then create first
			err = con.MakeDir(dir)
			if err != nil {
				panic("Error in making dir: " + dir + ", error : " + err.Error())
			}
			err = con.ChangeDir(dir)
			if err != nil {
				panic("Error in changing dir: " + dir + ", error : " + err.Error())
			}
		}
	}
	if dirCurrent, err = con.CurrentDir(); err != nil {
		panic("Error in getting current path , error : " + err.Error())
	}
}
