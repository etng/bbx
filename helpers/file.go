package helpers

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func PickInTgz(source, filename, target string) error {
	reader, err := os.Open(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()
	tarReader := tar.NewReader(archive)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		info := header.FileInfo()
		fmt.Println("checking ", header.Name, info)
		if info.IsDir() {

		} else {
			if header.Name == filename {
				fmt.Println("may be we can save it to ", target)
				file, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
				if err != nil {
					return err
				}
				n, _ := io.Copy(file, tarReader)
				fmt.Println("extracted to ", target, n)
				file.Close()
			}
		}
	}
	return nil
}
func UnTgz(source, target string) error {
	reader, err := os.Open(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()
	fmt.Println("archive name:", archive.Name)
	if archive.Name != "" {
		target = filepath.Join(target, archive.Name)
		writer, err := os.Create(target)
		if err != nil {
			return err
		}
		defer writer.Close()

		_, err = io.Copy(writer, archive)
		return err
	}
	tarReader := tar.NewReader(archive)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if err != nil {
			return err
		}
	}
	return nil
}
