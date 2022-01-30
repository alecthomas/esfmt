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
		Source []string `arg:"" help:"Source to format." type:"existingfile"`
	}
	// Token types that trigger a prefix indent.
	prefixIndentTypes = map[string]bool{
		"comment": true,
	}
	postfixNewlineTypes = map[string]bool{
		"comment":         true,
		"statement_block": true,
	}
	inhibitSpaceTypes = map[string]bool{
		"property_identifier": true,
		"member_expression":   true,
	}
)

func format(w io.Writer, source []byte, indent string, node *sitter.Node) {
	typ := node.Type()
	// fmt.Println(typ)
	isDeclaration := strings.HasSuffix(typ, "_declaration")
	if prefixIndentTypes[typ] || isDeclaration {
		fmt.Fprint(w, indent)
	}
	if node.ChildCount() == 0 {
		if !inhibitSpaceTypes[typ] {
			fmt.Fprint(w, " ")
		}
		fmt.Fprintf(w, "%s", source[node.StartByte():node.EndByte()])
	}
	if postfixNewlineTypes[typ] || isDeclaration {
		fmt.Fprintln(w)
	}
	childCount := int(node.ChildCount())
	for i := 0; i < childCount; i++ {
		format(w, source, indent, node.Child(i))
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
		// fmt.Println(tree.RootNode())
		format(os.Stdout, content, "", tree.RootNode())
	}
}
