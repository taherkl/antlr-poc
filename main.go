package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/taherkl/antlr-poc/parser"
)

func main() {
	// Your CQL query
	input := "SELECT * FROM table_name WHERE condition;"

	// Create an input stream of the query
	lexer := parser.NewCQLLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser with the token stream
	p := parser.NewCQLParser(stream)

	// Parse the input and get the parse tree
	tree := p.Start()

	// Do something with the parse tree (e.g., traverse or interpret)
	// For example, you might traverse the tree to perform actions based on CQL syntax.

	fmt.Println(tree.ToStringTree(nil, p))
	// Above line prints the parse tree representation (for demonstration purposes)
}
