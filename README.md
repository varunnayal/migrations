# migration

## Idea

1. Add files in `src/entity` directory
2. Update `loader/main.go` and add the entity.

3. Run `atlas migrate diff --env gorm`. This will result in a new file getting generated in `migrations/atlas-dir` directory(specified in [atlas.hcl](./atlas.hcl)).

```sh
tree migrations/atlas-dir
migrations/atlas-dir
├── 20230921084606.sql # generated by adding changes in loader/main.go file
└── atlas.sum # checksum file
```

Add a new entity

```sh
cat >> src/entity/user.go << EOL
package entity

type User struct {
  BaseModel
  Name  string `gorm:"type:varchar(100);not null"`
  Email string `gorm:"type:varchar(100);not null"`
}

EOL
```

Follow step 2 to 4 from above.
