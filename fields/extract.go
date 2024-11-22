package main

import (
	"archive/tar"
	"archive/zip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/ulikunitz/xz"
	"github.com/xor-gate/ar"
)

func downloadJar(url *url.URL, outputDir string) (string, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url.String(), nil)
	if err != nil {
		return "", fmt.Errorf("unable to download deb: %w", err)
	}

	debResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("unable to download deb: %w", err)
	}
	defer debResp.Body.Close()

	var uncompressedReader io.Reader

	arReader := ar.NewReader(debResp.Body)
	for {
		header, err := arReader.Next()
		if errors.Is(err, io.EOF) || header == nil {
			break
		}
		if err != nil {
			return "", fmt.Errorf("in ar next: %w", err)
		}

		// read the data file
		if header.Name == "data.tar.xz" {
			uncompressedReader, err = xz.NewReader(arReader)
			if err != nil {
				return "", fmt.Errorf("in xz reader: %w", err)
			}
			break
		}
	}
	if uncompressedReader == nil {
		return "", errors.New("unable to find .deb data file")
	}

	tarReader := tar.NewReader(uncompressedReader)

	var aceJar *os.File

	for {
		header, err := tarReader.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return "", fmt.Errorf("in next: %w", err)
		}

		if header.Typeflag != tar.TypeReg || header.Name != "./usr/lib/unifi/lib/ace.jar" {
			// skipping
			continue
		}

		aceJar, err = os.Create(filepath.Join(outputDir, "ace.jar"))
		if err != nil {
			return "", fmt.Errorf("unable to create temp file: %w", err)
		}
		_, err = io.Copy(aceJar, tarReader)
		if err != nil {
			return "", fmt.Errorf("unable to write ace.jar temp file: %w", err)
		}
	}

	if aceJar == nil {
		return "", errors.New("unable to find ace.jar")
	}

	defer aceJar.Close()

	return aceJar.Name(), nil
}

func extractJSON(jarFile, fieldsDir string) error {
	jarZip, err := zip.OpenReader(jarFile)
	if err != nil {
		return fmt.Errorf("unable to open jar: %w", err)
	}
	defer jarZip.Close()

	for _, f := range jarZip.File {
		if !strings.HasPrefix(f.Name, "api/fields/") || path.Ext(f.Name) != ".json" {
			// skip file
			continue
		}

		err = func() error {
			src, err := f.Open()
			if err != nil {
				return err
			}

			dst, err := os.Create(filepath.Join(fieldsDir, filepath.Base(f.Name)))
			if err != nil {
				return err
			}
			defer dst.Close()

			_, err = io.Copy(dst, src)
			if err != nil {
				return err
			}

			return nil
		}()
		if err != nil {
			return fmt.Errorf("unable to write JSON file: %w", err)
		}
	}

	settingsData, err := os.ReadFile(filepath.Join(fieldsDir, "Setting.json"))
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("unable to open settings file: %w", err)
	}

	var settings map[string]interface{}
	err = json.Unmarshal(settingsData, &settings)
	if err != nil {
		return fmt.Errorf("unable to unmarshal settings: %w", err)
	}

	for k, v := range settings {
		fileName := fmt.Sprintf("Setting%s.json", strcase.ToCamel(k))

		data, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return fmt.Errorf("unable to marshal setting %q: %w", k, err)
		}

		err = os.WriteFile(filepath.Join(fieldsDir, fileName), data, 0o755)
		if err != nil {
			return fmt.Errorf("unable to write new settings file: %w", err)
		}
	}

	// TODO: cleanup JSON
	return nil
}
