# myapp

## Features

- Obnoxious amount of `pre-commit` hooks. (`.pre-commit-config.yaml`)
- `run.sh` - convenient alternative to `Makefile` (call with `bash run.sh <function>`). Can be used locally and in CI systems.
- Smallest and secure `Dockerfile` for Go app based on scratch. (`dockerfiles/myapp.dockerfile`)
- Ready to use skeleton for multiple Go apps. (Shared packages under `pkg/`, app-specific packages under `pkg/myapp/`).
- Automatically reload multiple apps on change using [reflex](https://github.com/cespare/reflex) (`reflex.conf`).
- Small bits ready to be modified:
  - Structured logging (using `zap`, package `logging` - colored plaintext locally, json in production)
  - Example of e2e test (`pkg/myapp/e2e_test.go`)
  - Example of env vars handling (`pkg/myapp/config`)
  - Example of signal handling (`pkg/system/signals.go`)

> `.pre-conmmit-config.yaml` includes some Python-specific checks that are not run when no `python` files are present.

## Usage

To start using this template perform these steps:

- `rm go.sum go.mod`
- `go mod init <your_module_name>`
- delete all imports of gotemplate package
- replace `myapp` in filenames (folders in `cmd` and `pkg` and dockerfile in
`dockerfiles`) and source to the name of your app
- `go mod tidy`

### pre-commit

To use `pre-commit` locally run:

```bash
pre-commit install
pre-commit install --hook-type commit-msg
```

### GitHub

If you want `pre-commit` ci support - register action [here](https://pre-commit.ci/).

### GitLab

Use this Dockerfile:

```dockerfile
FROM python:3.10-slim
RUN apt update && apt install -y --no-install-recommends git && \
rm -f /var/cache/apt/archives/*.deb /var/cache/apt/archives/partial/*.deb /var/cache/apt/*.bin || true
RUN pip install --no-cache-dir pre-commit
```

With a job like this one:

```yaml
pre-commit:
  stage: .pre
  rules:
    - if: '$PRE_COMMIT_SKIP_BRANCH_PIPELINE && $CI_COMMIT_BRANCH'
      when: never
    - if: '$CI_PIPELINE_SOURCE == "merge_request_event"'
      exists:
        - .pre-commit-config.yaml
      when: on_success
    - if: '$CI_COMMIT_BRANCH'
      exists:
        - .pre-commit-config.yaml
      when: on_success
    - when: never
  image: registry.gitlab.com/.../docker-pre-commit
  script: >
    PRE_COMMIT_HOME=$CI_PROJECT_DIR/.cache/pre-commit
    SKIP=docker-compose-check,openapi-linter,dotenv-linter pre-commit run --all-files
  cache:
    paths:
      - $CI_PROJECT_DIR/.cache
```

### VSCode

#### golangci-lint integration

Paste this in your `~/.bashrc`:

```bash
function yes_or_no {
        QUESTION=$1
        DEFAULT=$2
        if [ "$DEFAULT" = true ]; then
                OPTIONS="[Y/n]"
                DEFAULT="y"
            else
                OPTIONS="[y/N]"
                DEFAULT="n"
        fi
        read -p "$QUESTION $OPTIONS " -n 1 -s -r INPUT
        INPUT=${INPUT:-${DEFAULT}}
        echo ${INPUT}
        if [[ "$INPUT" =~ ^[yY]$ ]]; then
            return 0
        else
            return 1
        fi
}
```

And use this task for VSCode. By using the function above you can clear all problems
by rerunning the task and answering `n` in interactive prompt.

```json
{
    "label": "golangci-lint",
    "type": "shell",
    "command": "bash -ic 'yes_or_no \"Run golangci-lint?\" true && golangci-lint run --enable-all || true'",
    "problemMatcher": {
        "owner": "golangci-lint",
        "fileLocation": [
            "relative",
            "${workspaceFolder}"
        ],
        "pattern": {
            "regexp": "^(.*):(\\d+):(\\d+):\\s+(.*)$",
            "file": 1,
            "line": 2,
            "column": 3,
            "message": 4
        }
    },
    "presentation": {
        "echo": false,
        "reveal": "always",
        "focus": true,
        "panel": "shared",
        "showReuseMessage": false,
        "clear": true
    }
}
```

#### golines integration

You can use [golines](https://github.com/segmentio/golines) formatter to fix long lines:

```bash
go install github.com/segmentio/golines@latest
```

Use this task:

```json
{
    "label": "golines",
    "type": "shell",
    "command":"bash -ic \"golines -w -m 100 ${file}\"",
    "presentation": {
        "echo": false,
        "reveal": "never",
        "focus": false,
        "panel": "shared",
        "showReuseMessage": false,
        "clear": false
    }
}
```
