package admin

import (
	"fmt"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/constant"
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/lib"
	"github.com/gocarina/gocsv"
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type UploadCsvController struct {
	Svc        domains.UploadCsvService
	SvcInvoice domains.InvoiceService
}

func (co *UploadCsvController) UploadCsvController(c echo.Context) error {
	file, err := c.FormFile("csv_file")

	if err != nil {
		return err
	}

	pathCsvInServer := os.Getenv("BASE_URL") + constant.STATIC_FILE_UPLOAD_CSV + file.Filename

	if !lib.CheckExtensionFile(filepath.Ext(pathCsvInServer)) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid extension file",
			"status":  http.StatusBadRequest,
		})
	}

	src, er := file.Open()
	if er != nil {
		return er
	}

	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			err.Error()
		}
	}(src)

	dir, _ := os.Getwd()
	locationFile := filepath.Join(dir, constant.DIR_FILE_UPLOAD_CSV, file.Filename)

	csvFile, fails := os.OpenFile(locationFile, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if fails != nil {
		panic(fails)
	}
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			err.Error()
		}
	}(csvFile)

	if _, fail := io.Copy(csvFile, src); fail != nil {
		return fail
	}

	csvfile, failed := os.OpenFile(locationFile, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if failed != nil {
		return failed
	}
	defer func(csvfile *os.File) {
		err := csvfile.Close()
		if err != nil {
			err.Error()
		}
	}(csvfile)
	invoices := []*admin.Invoice{}

	f := gocsv.UnmarshalFile(csvfile, &invoices)
	if f != nil {
		fmt.Println(f)
		return f
	}

	for _, invoice := range invoices {
		err := co.SvcInvoice.CreateInvoiceService(*invoice)
		if err != nil {
			return err
		}
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success to upload file",
		"data":    invoices,
	})
}
