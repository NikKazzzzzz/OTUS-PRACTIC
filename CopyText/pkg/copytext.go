package copytext

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
)

func Copy(from string, to string, limit int, offset int) error {
	sourceFile, err := os.Open(from)
	if err != nil {
		return err
	}
	defer func(sourceFile *os.File) {
		err := sourceFile.Close()
		if err != nil {

		}
	}(sourceFile)

	sourceInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	if !sourceInfo.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", from)
	}

	if offset > int(sourceInfo.Size()) {
		return fmt.Errorf("%s is too large (%d > %d)", from, sourceInfo.Size(), offset)
	}

	_, err = sourceFile.Seek(int64(offset), io.SeekStart)
	if err != nil {
		return err
	}

	destFile, err := os.Create(to)
	if err != nil {
		return err
	}
	defer func(destFile *os.File) {
		err := destFile.Close()
		if err != nil {

		}
	}(destFile)

	bytesToCopy := int(sourceInfo.Size()) - offset
	if limit > 0 && limit < bytesToCopy {
		bytesToCopy = limit
	}

	bar := pb.StartNew(bytesToCopy)

	buffer := make([]byte, 1024)
	totalCopied := 0

	for {
		bytesRead, readErr := sourceFile.Read(buffer)
		if bytesRead > 0 {
			bytesWritten, writeErr := destFile.Write(buffer[:bytesRead])
			if writeErr != nil {
				return writeErr
			}
			totalCopied += bytesWritten
			bar.Add(bytesWritten)
		}
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			return readErr
		}
		if limit > 0 && totalCopied >= limit {
			break
		}
	}
	bar.Finish()
	return nil
}
