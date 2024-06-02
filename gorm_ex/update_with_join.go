package gorm_ex

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

// UpdateSupportJoin hook gorm, 使gorm的update操作兼容join
// 使用方法:
// 1. db.Table(xx).Join(xx).Where(xx).Model(&UpdateSupportJoin{}).Updates(xx)
// 2. type yy struct{UpdateSupportJoin}; db.Table(xx).Join(xx).Where(xx).Model(&yy{}).Updates(xx)
type UpdateSupportJoin struct{}

func (u *UpdateSupportJoin) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Joins == nil || len(tx.Statement.Joins) == 0 {
		return nil
	}

	joins := make([]clause.Join, 0)
	for _, join := range tx.Statement.Joins {
		if tx.Statement.Schema == nil {
			joins = append(joins, clause.Join{
				Expression: clause.NamedExpr{SQL: join.Name, Vars: join.Conds},
			})
			continue
		}

		relation, ok := tx.Statement.Schema.Relationships.Relations[join.Name]
		if !ok {
			joins = append(joins, clause.Join{Expression: clause.NamedExpr{SQL: join.Name, Vars: join.Conds}})
			continue
		}

		tableAliasName := relation.Name
		express := make([]clause.Expression, len(relation.References))
		for idx, ref := range relation.References {
			if ref.OwnPrimaryKey {
				express[idx] = clause.Eq{
					Column: clause.Column{Table: clause.CurrentTable, Name: ref.PrimaryKey.DBName},
					Value:  clause.Column{Table: tableAliasName, Name: ref.ForeignKey.DBName},
				}
			} else {
				if ref.PrimaryValue == "" {
					express[idx] = clause.Eq{
						Column: clause.Column{Table: clause.CurrentTable, Name: ref.ForeignKey.DBName},
						Value:  clause.Column{Table: tableAliasName, Name: ref.PrimaryKey.DBName},
					}
				} else {
					express[idx] = clause.Eq{
						Column: clause.Column{Table: tableAliasName, Name: ref.ForeignKey.DBName},
						Value:  ref.PrimaryValue,
					}
				}
			}
		}

		if join.On != nil {
			onStmt := gorm.Statement{Table: tableAliasName, DB: tx}
			join.On.Build(&onStmt)
			onSQL := onStmt.SQL.String()
			vars := onStmt.Vars
			for idx, v := range onStmt.Vars {
				bindVar := strings.Builder{}
				onStmt.Vars = vars[0 : idx+1]
				tx.Dialector.BindVarTo(&bindVar, &onStmt, v)
				onSQL = strings.Replace(onSQL, bindVar.String(), "?", 1)
			}

			express = append(express, clause.Expr{SQL: onSQL, Vars: vars})
		}

		joins = append(joins, clause.Join{
			Type:  clause.LeftJoin,
			Table: clause.Table{Name: relation.FieldSchema.Table, Alias: tableAliasName},
			ON:    clause.Where{Exprs: express},
		})
	}

	tx.Statement.AddClause(updateClauseSupportJoin{
		Joins: joins,
	})
	return nil
}

type updateClauseSupportJoin struct {
	clause.Update
	Joins []clause.Join
}

// Name update clause name
func (updateEx updateClauseSupportJoin) Name() string {
	return "UPDATE"
}

func (updateEx updateClauseSupportJoin) Build(builder clause.Builder) {
	updateEx.Update.Build(builder)
	if updateEx.Joins != nil && len(updateEx.Joins) > 0 {
		for _, join := range updateEx.Joins {
			_ = builder.WriteByte(' ')
			join.Build(builder)
		}
	}
}

// MergeClause merge update clause
func (updateEx updateClauseSupportJoin) MergeClause(c *clause.Clause) {
	if v, ok := c.Expression.(updateClauseSupportJoin); ok {
		if updateEx.Modifier == "" {
			updateEx.Modifier = v.Modifier
		}
		if updateEx.Table.Name == "" {
			updateEx.Table = v.Table
		}
		if updateEx.Joins == nil {
			updateEx.Joins = v.Joins
		} else if v.Joins != nil && len(v.Joins) > 0 {
			updateEx.Joins = append(updateEx.Joins, v.Joins...)
		}
	}
	c.Expression = updateEx
}
