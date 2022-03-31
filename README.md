# myapp

A small highly opinionated Go template. **I do not recommend this as a good for everything (even not best practice) template. Meant to be used for with reductionist approach by deleting stuff before use.**

- Lots of hooks in pre-commit (most non Go related)
- `run.sh` - alternative to `Makefile` (call with `bash run.sh <function>`).
- Smallest and secure `Dockerfile` for Go app based on scratch.
- Ready to use skeleton for multiple Go services.
- Automatically reload multiple apps on change using [reflex](https://github.com/cespare/reflex) (`reflex.conf`).
- Small bits:
  - Structured logging (using `zap`, package `logging` - colored plaintext locally, json in production)
  - Example of env vars handling (`pkg/myapp/config`)

## Usage

To start using:

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
