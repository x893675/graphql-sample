package main

import (
	"fmt"
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"os"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}
	// Attaching the mutation function onto modelgen plugin
	p := modelgen.Plugin{
		MutateHook: addGormTags,
	}

	err = api.Generate(cfg,
		api.NoPlugins(),
		api.AddPlugin(&p),
	)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}

//Add the gorm tags to the model definition
func addGormTags(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {
			if field.Name == "id" {
				field.Tag += ` gorm:"` + "primaryKey;type:uuid;column:id;default:uuid_generate_v4();index;" + `"`
			}
			if model.Name == "Question" && field.Name == "title" {
				field.Tag += ` gorm:"unique" db:"` + field.Name + `"`
			} else {
				field.Tag += ` db:"` + field.Name + `"`
			}
		}
	}
	return b
}
