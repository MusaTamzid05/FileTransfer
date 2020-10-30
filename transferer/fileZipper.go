package transferer

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

type FileZipper struct {
}

func NewFileZipper() *FileZipper {
	zipper := FileZipper{}
	return &zipper
}

func (z *FileZipper) Zip(path, dst string) error {

	paths, err := ListFiles(path)

	if err != nil {
		return err
	}

	log.Printf("Total files : %d\n", len(paths))
	err = z.ZipFiles(dst, paths)
	return err

}

func (z *FileZipper) ZipFiles(outputPath string, srcPaths []string) error {

	newZipFile, err := os.Create(outputPath)

	if err != nil {
		return err
	}

	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, path := range srcPaths {
		err = z.AddFileToZip(zipWriter, path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (z *FileZipper) AddFileToZip(zipWriter *zip.Writer, path string) error {

	fileToZip, err := os.Open(path)

	if err != nil {
		return err
	}

	defer fileToZip.Close()

	info, err := fileToZip.Stat()

	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)

	if err != nil {
		return err
	}

	header.Name = path
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)

	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)
	return err
}
