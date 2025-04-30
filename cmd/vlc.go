package cmd

import (
	"archivator/lib/vlc"
	"errors"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("please specify a file to pack")

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using VLC",
	Run:   pack,
}

func pack(_ *cobra.Command, args []string) {
	if len(args) < 1 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}

	defer func(r *os.File) {
		err := r.Close()
		if err != nil {
			handleErr(err)
		}
	}(r)

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := vlc.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}

}

func packedFileName(path string) string {
	fileName := filepath.Base(path)
	ext := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, ext)
	return baseName + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
