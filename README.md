To code in this project using neovim you need to install the following packages:
```bash
go install github.com/nametake/golangci-lint-langserver@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest 
go install github.com/a-h/templ/cmd/templ@latest
```
Also remember to configure the PATH to run executables globally, this is needed for the configuration of neovim-lspconfig check the [docs](https://github.com/neovim/nvim-lspconfig/blob/master/doc/configs.md#golangci_lint_ls) for more info on this, also you can check the following [question](https://stackoverflow.com/questions/36650052/golang-equivalent-of-npm-install-g) on stackoverflow.
```bash
export PATH="$HOME/go/bin:$PATH"
```

For more domentation check the following resources:
### Documentation
- [Go standard library documentation](https://pkg.go.dev/std)
- [GoDoc - place to find opensource packages](https://pkg.go.dev/)
### Courses
- [Highly recomended video series](https://www.youtube.com/playlist?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)
- [Reddit helpful courses](https://www.reddit.com/r/golang/comments/1cb0e7l/what_courses_were_extremely_helpful/)
### Design patterns
- [Golang patterns](https://github.com/tmrts/go-patterns?tab=readme-ov-file) Really good one!!
- [Reddit Golang desing patterns](https://www.reddit.com/r/golang/comments/1887y1b/favorite_golang_design_patterns/) Really good one!!

## Go mod
Se necesita usar `go mod init` para que el proyecto funcione de manera correcta.
```bash
go mod init github.com/ZOLUXERO/go-test
```
