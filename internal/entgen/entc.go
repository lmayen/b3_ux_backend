//go:build ignore

package entgen

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

type assetTemplate struct {
	Source string
	Name   string
	Target string
}

func main() {
	config := &gen.Config{
		Target:  "./internal/entities",
		Package: "b3_ux_backend/internal/entities",

		Hooks: []gen.Hook{
			generateAssets(
				assetTemplate{
					Source: "./internal/entgen/models/data.tmpl",
					Name:   "data",
					Target: "./models/data.ts",
				},
				assetTemplate{
					Source: "./internal/entgen/doc/schema.tmpl",
					Name:   "schema",
					Target: "./doc/schema.md",
				},
			),
		},
	}

	err := entc.Generate(
		"./internal/schema",
		config,

		// Normal custom Go templates.
		entc.TemplateDir("./internal/entgen/templates"),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func generateAssets(assets ...assetTemplate) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(graph *gen.Graph) error {
			// Generate Ent/Go files first.
			if err := next.Generate(graph); err != nil {
				return err
			}

			for _, asset := range assets {
				if err := generateAsset(graph, asset); err != nil {
					return err
				}
			}

			return nil
		})
	}
}

func generateAsset(graph *gen.Graph, asset assetTemplate) error {
	tmpl, err := gen.NewTemplate(asset.Name).ParseFiles(asset.Source)
	if err != nil {
		return fmt.Errorf("parse template %s: %w", asset.Source, err)
	}

	var buffer bytes.Buffer

	if err := tmpl.ExecuteTemplate(&buffer, asset.Name, graph); err != nil {
		return fmt.Errorf("execute template %s: %w", asset.Source, err)
	}

	if err := os.MkdirAll(filepath.Dir(asset.Target), 0755); err != nil {
		return err
	}

	if err := os.WriteFile(asset.Target, buffer.Bytes(), 0644); err != nil {
		return fmt.Errorf("write %s: %w", asset.Target, err)
	}

	return nil
}
