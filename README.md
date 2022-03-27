# myapp

Steps:

- `rm go.sum go.mod`
- `go mod init <your_module_name>`
- delete all imports of gotemplate package
- replace `myapp` in filenames (folders in `cmd` and `pkg` and dockerfile in `dockerfiles`) and source to the name of your app
- `go mod tidy`

If you want `pre-commit` ci support - register action [here](https://pre-commit.ci/). To use locally run `pre-commit install`.
