/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"fmt"
	"go/token"
	"strings"

	ast "github.com/dave/dst"
)

// TextLiteral creates the AST node for a literal text value
// No additional text is included
func TextLiteral(content string) *ast.BasicLit {
	return &ast.BasicLit{
		Value: content,
		Kind:  token.STRING,
	}
}

// TextLiteralf creates the AST node for literal text based on a format string
func TextLiteralf(format string, a ...interface{}) *ast.BasicLit {
	return TextLiteral(fmt.Sprintf(format, a...))
}

// StringLiteral creates the AST node for a literal string value
// Leading and trailing quotes are added as required and any existing quotes are escaped
func StringLiteral(content string) *ast.BasicLit {
	// Pay attention to the string escaping here!
	escaped := content
	escaped = strings.ReplaceAll(escaped, "\\", "\\\\")
	escaped = strings.ReplaceAll(escaped, "\"", "\\\"")

	c := "\"" + escaped + "\""
	return TextLiteral(c)
}

// StringLiteralf creates the AST node for a literal string value based on a format string
// Leading and trailing quotes are added as required and any existing quotes are escaped
func StringLiteralf(format string, a ...interface{}) *ast.BasicLit {
	return StringLiteral(fmt.Sprintf(format, a...))
}

// IntLiteral() create the AST node for a literal integer value
func IntLiteral(value int) *ast.BasicLit {
	return &ast.BasicLit{
		Value: fmt.Sprint(value),
		Kind:  token.INT,
	}
}
