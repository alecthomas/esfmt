# esfmt - an opinionated, zero-configuration formatter for ES/TS/ESX/TSX

Status: total hack

In the spirit of [esbuild](https://esbuild.github.io/), esfmt strives to bring
the speed and convenience of a statically compiled single binary
distribution to formatting TSX and its relatives.

Leveraging the excellent [Treesitter](https://tree-sitter.github.io/tree-sitter/)
will hopefully allow us to do this with relatively little effort.

## Developing

Copy JS/TSX/ES/ESX into the [Tree-Sitter playground](https://tree-sitter.github.io/tree-sitter/playground).
This formats the parse tree in a nice way and allows interactive selection.