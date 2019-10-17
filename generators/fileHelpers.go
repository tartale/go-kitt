package generators

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"

	"github.com/MarcGrol/golangAnnotations/model"
	"github.com/pkg/errors"

	"github.com/tartale/go-kitt/config"
	"github.com/tartale/go-kitt/lib/errorz"
)

const GeneratedPathNewExtension = "gen.new.go"
const GeneratedPathExtension = "gen.go"

func NewGeneratedPathForInterface(intf model.Interface, suffix string) string {
	interfaceDir := path.Dir(intf.Filename)
	interfaceBasename := path.Base(intf.Filename)
	interfaceBasenameNoExt := strings.TrimSuffix(interfaceBasename, path.Ext(interfaceBasename))

	return path.Join(interfaceDir, fmt.Sprintf("%s%s.%s", interfaceBasenameNoExt, suffix, GeneratedPathNewExtension))
}

type GeneratedPaths []string

func (o GeneratedPaths) FinalizeAll() error {

	var errs errorz.Errors
	var finalPaths []string
	for _, path := range o {
		newPath := strings.TrimSuffix(path, GeneratedPathNewExtension) + GeneratedPathExtension
		err := os.Rename(path, newPath)
		if err != nil {
			errs = append(errs, err)
		} else {
			finalPaths = append(finalPaths, newPath)
		}
	}
	if config.Config.FormatTool != "" {
		fmtArgs := []string{"-w"}
		fmtArgs = append(fmtArgs, finalPaths...)
		output, err := exec.Command(config.Config.FormatTool, fmtArgs...).CombinedOutput()
		if err != nil {
			errs = append(errs, errors.Wrap(err, string(output)))
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (o GeneratedPaths) RemoveAll() error {

	var errs errorz.Errors
	for _, gp := range o {
		err := os.Remove(string(gp))
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errs
	}

	return nil
}

func GenerateFile(filePath string, tmpl *template.Template, data interface{}) error {
	generatedWriter, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0660)
	if err != nil {
		return err
	}
	defer generatedWriter.Close()
	return tmpl.Execute(generatedWriter, data)
}
