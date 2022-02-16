# myapp

Steps:

- `rm go.sum go.mod`
- `go mod init <your_module_name>`
- delete all imports of gotemplate package
- `go mod tidy`
- replace `myapp` in filenames (folder in `cmd` and dockerfile in `dockerfiles`) and source to the name of your app
