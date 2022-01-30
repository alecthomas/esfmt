package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/typescript/tsx"
)

var (
	cli struct {
		DebugAST  bool     `help:"Debug dump AST."`
		DebugType bool     `help:"Debug AST type information."`
		Source    []string `arg:"" help:"Source to format." type:"existingfile"`
	}
	// Token types that trigger a prefix indent.
	prefixIndentTypes = map[string]bool{
		"comment": true,
	}
	postfixIndentTypes = map[string]bool{
		";": true,
		"}": true,
	}
	postfixNewlineTypes = map[string]bool{
		"comment":         true,
		"statement_block": true,
		"{":               true,
		"}":               true,
		";":               true,
	}
	postfixSpaceTypes = map[string]bool{
		"interface": true,
		"class":     true,
		"const":     true,
		":":         true,
	}
	prefixSpaceTypes = map[string]bool{
		"{": true,
	}
)

func format(w io.Writer, source []byte, indent string, node *sitter.Node) {
	typ := node.Type()
	isDeclaration := strings.HasSuffix(typ, "_declaration")
	if prefixIndentTypes[typ] || isDeclaration {
		fmt.Fprint(w, indent)
	}
	if prefixSpaceTypes[typ] {
		fmt.Fprint(w, " ")
	}
	if node.ChildCount() == 0 {
		fmt.Fprintf(w, "%s", source[node.StartByte():node.EndByte()])
		if cli.DebugType {
			fmt.Fprintf(w, "\033[32m(%s)\033[0m", typ)
		}
	}
	if postfixSpaceTypes[typ] {
		fmt.Fprint(w, " ")
	}
	childCount := int(node.ChildCount())
	for i := 0; i < childCount; i++ {
		format(w, source, indent, node.Child(i))
	}
	if isDeclaration || postfixNewlineTypes[typ] {
		fmt.Fprintln(w)
	}
}

func main() {
	kctx := kong.Parse(&cli)
	parser := sitter.NewParser()
	lang := tsx.GetLanguage()
	parser.SetLanguage(lang)
	for _, source := range cli.Source {
		content, err := ioutil.ReadFile(source)
		kctx.FatalIfErrorf(err)
		tree := parser.Parse(nil, content)
		if cli.DebugAST {
			fmt.Fprintln(os.Stderr, tree.RootNode())
		}
		// fmt.Println(tree.RootNode())
		format(os.Stdout, content, "", tree.RootNode())
	}
}
