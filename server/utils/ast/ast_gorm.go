package ast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

// Automatically register an automatic migration for gorm.go
func AddRegisterTablesAst(path, funcName, pk, varName, dbName, model string) {
	modelPk := fmt.Sprintf("github.com/flipped-aurora/gin-vue-admin/server/model/%s", pk)
	src, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "", src, 0)
	if err != nil {
		fmt.Println(err)
	}
	AddImport(astFile, modelPk)
	FuncNode := FindFunction(astFile, funcName)
	if FuncNode != nil {
		ast.Print(fileSet, FuncNode)
	}
	addDBVar(FuncNode.Body, varName, dbName)
	addAutoMigrate(FuncNode.Body, varName, pk, model)
	var out []byte
	bf := bytes.NewBuffer(out)
	printer.Fprint(bf, fileSet, astFile)

	os.WriteFile(path, bf.Bytes(), 0666)
}

// Add a db library variable
func addDBVar(astBody *ast.BlockStmt, varName, dbName string) {
	if dbName == "" {
		return
	}
	dbStr := fmt.Sprintf("\"%s\"", dbName)

	for i := range astBody.List {
		if assignStmt, ok := astBody.List[i].(*ast.AssignStmt); ok {
			if ident, ok := assignStmt.Lhs[0].(*ast.Ident); ok {
				if ident.Name == varName {
					return
				}
			}
		}
	}
	assignNode := &ast.AssignStmt{
		Lhs: []ast.Expr{
			&ast.Ident{
				Name: varName,
			},
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "global",
					},
					Sel: &ast.Ident{
						Name: "GetGlobalDBByDBName",
					},
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: dbStr,
					},
				},
			},
		},
	}
	astBody.List = append([]ast.Stmt{assignNode}, astBody.List...)
}

// Add AutoMigrate method to db library variables
func addAutoMigrate(astBody *ast.BlockStmt, dbname string, pk string, model string) {
	if dbname == "" {
		dbname = "db"
	}
	flag := true
	ast.Inspect(astBody, func(node ast.Node) bool {
		// First determine whether the method call statement that needs to be added exists. If it does not exist, go directly to the logic below.
		switch n := node.(type) {
		case *ast.CallExpr:
			// Determine whether the AutoMigrate statement is found
			if s, ok := n.Fun.(*ast.SelectorExpr); ok {
				if x, ok := s.X.(*ast.Ident); ok {
					if s.Sel.Name == "AutoMigrate" && x.Name == dbname {
						flag = false
						if !NeedAppendModel(n, pk, model) {
							return false
						}
						// Determine that the AutoMigrate statement has been found
						n.Args = append(n.Args, &ast.CompositeLit{
							Type: &ast.SelectorExpr{
								X: &ast.Ident{
									Name: pk,
								},
								Sel: &ast.Ident{
									Name: model,
								},
							},
						})
						return false
					}
				}
			}
		}
		return true
		//Then determine whether pk.model exists. If it exists, jump out directly. If it does not exist, push a message to the node of the method call statement that has been found.
	})

	if flag {
		exprStmt := &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: dbname,
					},
					Sel: &ast.Ident{
						Name: "AutoMigrate",
					},
				},
				Args: []ast.Expr{
					&ast.CompositeLit{
						Type: &ast.SelectorExpr{
							X: &ast.Ident{
								Name: pk,
							},
							Sel: &ast.Ident{
								Name: model,
							},
						},
					},
				},
			}}
		astBody.List = append(astBody.List, exprStmt)
	}
}

// Add actual parameters to automigrate
func NeedAppendModel(callNode ast.Node, pk string, model string) bool {
	flag := true
	ast.Inspect(callNode, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.SelectorExpr:
			if x, ok := n.X.(*ast.Ident); ok {
				if n.Sel.Name == model && x.Name == pk {
					flag = false
					return false
				}
			}
		}
		return true
	})
	return flag
}
