package generatemodel

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type block struct {
	Type           string `yaml:"type"`
	Field          string `yaml:"field"`
	Table          string `yaml:"table"`
	Many2Many      string `yaml:"many2many"`
	ForeignKey     string `yaml:"foreignKey"`
	JoinForeignKey string `yaml:"joinForeignKey"`
	References     string `yaml:"references"`
	JoinReferences string `yaml:"joinReferences"`
}

func GenerateModel(db *gorm.DB, g *gen.Generator, path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var blockMap map[string][]block
	err = yaml.Unmarshal(file, &blockMap)
	if err != nil {
		return err
	}

	tableList, err := db.Migrator().GetTables()
	if err != nil {
		return err
	}

	l := fmt.Sprintf("find %d table from db: %s", len(tableList), tableList)
	db.Logger.Info(context.Background(), l)
	log.Println(l)

	modelStructNameMap := make(map[string]string)
	for _, v := range tableList {
		model := g.GenerateModel(v)
		modelStructNameMap[v] = model.ModelStructName
		g.ApplyBasic(model)
	}

	for k, v := range blockMap {
		opts := make([]gen.ModelOpt, 0)
		for _, b := range v {
			relationship, err := getRelationshipType(b)
			if err != nil {
				continue
			}
			tag := getGormTag(b)
			opts = append(opts, gen.FieldRelate(
				relationship,
				b.Field,
				g.Data[modelStructNameMap[b.Table]].QueryStructMeta,
				&field.RelateConfig{
					GORMTag: tag,
				},
			))
		}
		model := g.GenerateModel(k, opts...)
		modelStructNameMap[k] = model.ModelStructName
		g.ApplyBasic(model)
	}

	return nil
}

func getRelationshipType(b block) (field.RelationshipType, error) {
	switch b.Type {
	case "HasOne":
		return field.HasOne, nil
	case "HasMany":
		return field.HasMany, nil
	case "BelongsTo":
		return field.BelongsTo, nil
	case "Many2Many":
		return field.Many2Many, nil
	default:
		return "", errors.New("type error: " + b.Type)
	}
}

func getGormTag(b block) field.GormTag {
	tag := make(map[string][]string)
	tag["foreignKey"] = []string{b.ForeignKey}
	tag["references"] = []string{b.References}
	if b.Type == "Many2Many" {
		tag["many2many"] = []string{b.Many2Many}
		tag["joinForeignKey"] = []string{b.JoinForeignKey}
		tag["joinReferences"] = []string{b.JoinReferences}
	}
	return tag
}
