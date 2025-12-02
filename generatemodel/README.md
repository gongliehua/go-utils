# generate-model

生成模型

configs/generate-model.yml

```
admins:
  - type: Many2Many
    field: Roles
    table: roles
    many2many: admin_roles
    foreignKey: ID
    references: ID
    joinForeignKey: AdminID
    joinReferences: RoleID
roles:
  - type: Many2Many
    field: Permissions
    table: permissions
    many2many: role_permissions
    foreignKey: ID
    references: ID
    joinForeignKey: RoleID
    joinReferences: PermissionID
permissions:
  - type: HasMany
    field: Children
    table: permissions
    foreignKey: ParentID
    references: ID
```

cmd/generate-model/main.go

```
package main

import (
	"flag"
	"gorm.io/gen"
	"os"
)

func main() {
	// 生成 model
	dir, _ := os.Getwd()
	g := gen.NewGenerator(gen.Config{
		OutPath:      dir + "/internal/query",
		ModelPkgPath: dir + "/internal/model",
	})
	g.UseDB(global.DB)
	err = generatemodel.GenerateModel(global.DB, g, "./configs/generate-model.yml")
	if err != nil {
		panic(err)
	}
	// 不生成 query
	for k := range g.Data {
		delete(g.Data, k)
	}
	g.Execute()
}
```
