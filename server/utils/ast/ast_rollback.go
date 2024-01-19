package ast

import (
	"bytes"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
)

func RollBackAst(pk, model string) {
	RollGormBack(pk, model)
	RollRouterBack(pk, model)
}

func RollGormBack(pk, model string) {

	// First analyze how many node blocks exist with ttt as the caller
	// If there are multiple, just delete the corresponding block.
	// If it is single, then the import needs to be eliminated.
	path := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize", "gorm.go")
	src, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "", src, 0)
	if err != nil {
		fmt.Println(err)
	}
	var n *ast.CallExpr
	var k int = -1
	var pkNum = 0
	ast.Inspect(astFile, func(node ast.Node) bool {
		if node, ok := node.(*ast.CallExpr); ok {
			for i := range node.Args {
				pkOK := false
				modelOK := false
				ast.Inspect(node.Args[i], func(item ast.Node) bool {
					if ii, ok := item.(*ast.Ident); ok {
						if ii.Name == pk {
							pkOK = true
							pkNum++
						}
						if ii.Name == model {
							modelOK = true
						}
					}
					if pkOK && modelOK {
						n = node
						k = i
					}
					return true
				})
			}
		}
		return true
	})
	if k > 0 {
		n.Args = append(append([]ast.Expr{}, n.Args[:k]...), n.Args[k+1:]...)
	}
	if pkNum == 1 {
		var imI int = -1
		var gp *ast.GenDecl
		ast.Inspect(astFile, func(node ast.Node) bool {
			if gen, ok := node.(*ast.GenDecl); ok {
				for i := range gen.Specs {
					if imspec, ok := gen.Specs[i].(*ast.ImportSpec); ok {
						if imspec.Path.Value == "\"github.com/flipped-aurora/gin-vue-admin/server/model/"+pk+"\"" {
							gp = gen
							imI = i
							return false
						}
					}
				}
			}
			return true
		})

		if imI > -1 {
			gp.Specs = append(append([]ast.Spec{}, gp.Specs[:imI]...), gp.Specs[imI+1:]...)
		}
	}

	var out []byte
	bf := bytes.NewBuffer(out)
	printer.Fprint(bf, fileSet, astFile)
	os.Remove(path)
	os.WriteFile(path, bf.Bytes(), 0666)

}

func RollRouterBack(pk, model string) {

	// First capture all code block structures {}
	// Analyze whether there is a variable called pk+Router in the structure
	// Then get the code block pointer and eliminate the internal code that needs to be rolled back.
	path := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize", "router.go")
	src, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "", src, 0)
	if err != nil {
		fmt.Println(err)
	}

	var block *ast.BlockStmt
	ast.Inspect(astFile, func(node ast.Node) bool {
		if n, ok := node.(*ast.BlockStmt); ok {
			ast.Inspect(n, func(bNode ast.Node) bool {
				if in, ok := bNode.(*ast.Ident); ok {
					if in.Name == pk+"Router" {
						block = n
						return false
					}
				}
				return true
			})
			return true
		}
		return true
	})
	var k int
	for i := range block.List {
		if stmtNode, ok := block.List[i].(*ast.ExprStmt); ok {
			ast.Inspect(stmtNode, func(node ast.Node) bool {
				if n, ok := node.(*ast.Ident); ok {
					if n.Name == "Init"+model+"Router" {
						k = i
						return false
					}
				}
				return true
			})
		}
	}

	block.List = append(append([]ast.Stmt{}, block.List[:k]...), block.List[k+1:]...)

	if len(block.List) == 1 {
		// This means that this block has no meaning.
		block.List = nil
		// TODO delete empty {}
	}

	var out []byte
	bf := bytes.NewBuffer(out)
	printer.Fprint(bf, fileSet, astFile)
	os.Remove(path)
	os.WriteFile(path, bf.Bytes(), 0666)
}
