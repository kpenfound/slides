# Intro to Github Actions with Dagger

A brief overview of Github Actions and how you can run Dagger in Actions

---

## The Github Actions Basics

- Runs workflows based on events in your repository
- Configured with YAML
- GitHub offers a free tier but there are many options for more advanced needs

---

## Github Actions YAML

`./.github/workflows/ci.yaml`

```yaml
name: ci
on:                                # when this workflow should run
  push:
    branches: ["**", "!main"]
jobs:
  ci:
    name: ci
    runs-on: ubuntu-latest         # the execution environment
    steps:
      - uses: actions/checkout@v4  # checkout the code
      - uses: actions/setup-go@v5  # A "setup" action
        with:
          go-version: '>=1.17.0'
      - name: test                 # run a shell command
        run: go test ./...
        env:                       # environment variables
          CGO_ENABLED: 1
          MY_VAR_FOO: ${{ secrets.MY_VAR_FOO }}
```

---

### Events that trigger workflows

Run a workflow...

when a commit is pushed:
```yaml
on:
  push:
    branches: ["**", "!main"]
```

when a pull request is opened:
```yaml
on:
  pull_request:
    types: [opened, reopened]
```

when an issue or pull request is commented on:
```yaml
on:
  issue_comment:
    types: [created]
```

on a schedule:
```yaml
on:
  schedule:
    - cron:  '30 5,17 * * *'
```

---

### Steps

A step can be...

an action:
```yaml
- uses: actions/setup-go@v5
  with:
    go-version: '>=1.21.0'
```

a shell command:
```yaml
- name: test
  run: go test ./...
```

---

## Dagger Functions

```bash
dagger -m github.com/kpenfound/greetings-api functions
```

Learn how to Daggerize your project at [https://docs.dagger.io/quickstart/ci](https://docs.dagger.io/quickstart/ci)

---

## The dagger-for-github Action

[https://github.com/dagger/dagger-for-github](https://github.com/dagger/dagger-for-github)

`./.github/workflows/ci.yaml`

```yaml
name: ci
on:
  push
jobs:
  ci:
    name: ci
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: check
        uses: dagger/dagger-for-github@8.0.0
        with:
          call: check
          cloud-token: ${{ secrets.DAGGER_CLOUD_TOKEN }}
```

---

## The dagger-for-github Action

[https://github.com/dagger/dagger-for-github](https://github.com/dagger/dagger-for-github)

`./.github/workflows/ci.yaml`

```yaml
name: ci
on:
  push
jobs:
  ci:
    name: ci
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: dagger/dagger-for-github@8.0.0
      - name: check
        run: dagger call check
        env:
          DAGGER_CLOUD_TOKEN: ${{ secrets.DAGGER_CLOUD_TOKEN }}
```

---

## Next steps

- Self host your Github Actions runners [https://docs.dagger.io/ci/integrations/kubernetes/](https://docs.dagger.io/ci/integrations/kubernetes/)

- Use [Depot](https://depot.dev) for faster runners with caching
