package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	_ "github.com/jmoiron/sqlx"

	"github.com/febriliankr/entable/datatype"
	"github.com/stoewer/go-strcase"
)

const TemplateUsecase = "template/usecase.gotemplate"
const TemplateRepo = "template/repo.gotemplate"
const TemplateStruct = "template/struct.gotemplate"
const TemplateInterface = "template/interfaces.gotemplate"

func main() {

	// get flag from running command

	if len(os.Args) < 1 {
		fmt.Println("please provide table name")
		return
	}

	tableFilename := os.Args[1]

	createRequiredFolders()
	err := Generate(tableFilename)

	if err != nil {
		fmt.Println(err)
	}
}

func createRequiredFolders() {
	os.Mkdir("generated", 0755)
	os.Mkdir("generated/entities", 0755)
	os.Mkdir("generated/repo", 0755)
	os.Mkdir("generated/interfaces", 0755)
	os.Mkdir("generated/usecase", 0755)
}

func Generate(filepath string) error {

	tableName := strings.Split(filepath, ".")[0]

	file, err := os.ReadFile(filepath)

	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var entity EntityRequest

	entity.TableNameSnake = tableName
	entity.TableNameCamel = strcase.UpperCamelCase(tableName)

	rows := strings.Split(string(file), "\n")

	for index, row := range rows {
		columnName := strings.Split(row, "\t")[1]
		columnSQLType := strings.Split(row, "\t")[2]
		goType := datatype.GenerateDataType(columnSQLType)
		camelColumnName := strcase.UpperCamelCase(columnName)

		entityColumn := EntityColumn{
			Index:     index + 1,
			CamelName: camelColumnName,
			Type:      goType,
			SnakeName: columnName,
		}

		entity.Columns = append(entity.Columns, entityColumn)

		if containsIDorUUID(columnName) && entity.PrimaryKey.CamelName == "" {
			entity.PrimaryKey = entityColumn
		}

	}

	if entity.PrimaryKey.SnakeName == "" {
		return fmt.Errorf("primary key not found")
	}

	err = generateAndWriteToFile(entity)
	if err != nil {
		return err
	}
	return nil
}

func generateAndWriteToFile(entity EntityRequest) error {
	entityCode, err := generateCode(TemplateStruct, entity)
	if err != nil {
		return err
	}
	err = os.WriteFile("generated/entities/"+entity.TableNameCamel+".go", []byte(entityCode), 0644)
	if err != nil {
		return err
	}

	repoCode, err := generateCode(TemplateRepo, entity)
	if err != nil {
		return err
	}
	err = os.WriteFile("generated/repo/"+entity.TableNameSnake+".go", []byte(repoCode), 0644)
	if err != nil {
		return err
	}

	interfaceCode, err := generateCode(TemplateInterface, entity)
	if err != nil {
		return err
	}
	err = os.WriteFile("generated/interfaces/"+entity.TableNameSnake+".go", []byte(interfaceCode), 0644)
	if err != nil {
		return err
	}

	usecaseCode, err := generateCode(TemplateUsecase, entity)
	if err != nil {
		return err
	}

	err = os.WriteFile("generated/usecase/"+entity.TableNameSnake+".go", []byte(usecaseCode), 0644)
	if err != nil {
		return err
	}

	return nil

}

type EntityRequest struct {
	PrimaryKey     EntityColumn
	Columns        []EntityColumn
	TableNameSnake string
	TableNameCamel string
}

type EntityColumn struct {
	Index     int
	CamelName string
	Type      string
	SnakeName string
}

func generateCode(templatePath string, data interface{}) (string, error) {
	var code string
	htmlTemplate, err := template.ParseFiles(templatePath)
	if err != nil {
		return code, fmt.Errorf("error parsing template: %v", err)
	}
	buf := new(bytes.Buffer)
	err = htmlTemplate.Execute(buf, data)
	if err != nil {
		return code, fmt.Errorf("error executing template: %v", err)
	}
	code = buf.String()
	return code, nil
}

func containsIDorUUID(str string) bool {
	return strings.Contains(str, "id") || strings.Contains(str, "uuid")
}
