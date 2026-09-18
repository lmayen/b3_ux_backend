## Create a new model
```shell
go run -mod=mod entgo.io/ent/cmd/ent new Asset --target internal/schema
```
## Generate Ent code
```
go run -mod=mod entgo.io/ent/cmd/ent generate --target ./internal/entities ./internal/schema
```

It needs in generate.go
```go 
// internal/entities/generate.go

package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./internal/entities

```

---

## Templates:

````md
# Ent Template Cheat Sheet

## Header

```gotemplate
{{ template "header" $ }}
````

Generates the standard Ent generated-file header and package declaration.

---

## IDE autocompletion

```gotemplate
{{/* gotype: entgo.io/ent/entc/gen.Graph */}}
```

IDE hint only. No effect on generation.

---

## Nodes

```gotemplate
{{ range $node := $.Nodes }}
    {{ $node.Name }}
    {{ $node.ID }}
    {{ $node.Fields }}
    {{ $node.Edges }}
    {{ $node.Indexes }}
    {{ $node.Package }}
    {{ $node.QueryName }}
{{ end }}
```

---

## Fields

```gotemplate
{{ range $field := $node.Fields }}
    {{ $field.Name }}
    {{ $field.StructField }}
    {{ $field.Optional }}
    {{ $field.Nillable }}
    {{ $field.Type }}

    {{ if $field.IsString }}
    {{ end }}
    
    {{ if $field.IsBool }}
    {{ end }}
    
    {{ if $field.IsInt }}
    {{ end }}
    
    {{ if $field.IsInt64 }}
    {{ end }}
    
    {{ if $field.IsTime }}
    {{ end }}
    
    {{ if $field.IsUUID }}
    {{ end }}
    
    {{ if $field.IsEnum }}
    {{ end }}
    
    {{ if $field.IsJSON }}
    {{ end }}
{{ end }}

```

---

## Edges

```gotemplate
{{ range $edge := $node.Edges }}
    {{ $edge.Name }}
    {{ $edge.Type.Name }}
    {{ $edge.StructField }}

    {{ $edge.Type.Name }}
    {{ $edge.Type.Package }}
    {{ $edge.Type.Fields }}
    {{ $edge.Type.Edges }}
    
    {{ $edge.O2O }}
    {{ $edge.O2M }}
    {{ $edge.M2O }}
    {{ $edge.M2M }}
{{ end }}
```

---

## Indexes

```gotemplate
{{ range $index := $node.Indexes }}
    {{ $index.Name }}
    {{ $index.Unique }}
    {{ $index.Columns }}
{{ end }}
```

---

## Logic

```gotemplate
{{ if ... }}
{{ else }}
{{ end }}

{{ if eq $field.Name "name" }}
{{ end }}

{{ range ... }}
{{ end }}

{{ $value := ... }}
```

---

````md
# Ent `entc.go` Cheat Sheet

## Location

```text
/internal
    /schema
    /entities
    /entgen
        entc.go
        /templates
````

---

## Generation command

```bash
go run entgo.io/ent/cmd/ent generate --target ./internal/entities --template ./internal/entgen/templates ./internal/schema
```

## Or
`entc.go`:
```go
//go:build ignore

package main

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
		Package: "YOUR_MODULE/internal/entities",

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
```
