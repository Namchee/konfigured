# Konfigured

[![Go Version Badge](https://img.shields.io/github/go-mod/go-version/namchee/konfigured)](https://github.com/Namchee/konfigured) [![Go Report Card](https://goreportcard.com/badge/github.com/Namchee/konfigured)](https://goreportcard.com/report/github.com/Namchee/konfigured)

Prevent malformed configuration files from being merged to your project. No more breaking builds caused by bad configuration files.

## Usage 

You can integrate Konfigured to your existing GitHub action workflow by using `Namchee/konfigured@<version>` in one of your jobs using the YAML syntax.

Below is the example of integrating Konfigured to your workflow in your GitHub action workflow.

```yml
on:
  pull_request:

jobs:
  cpr:
    runs-on: ubuntu-latest
    steps:
      - name: Validate configuration file
        uses: Namchee/konfigured@v(version)
        with:
          token: <YOUR_GITHUB_ACCESS_TOKEN_HERE>
```

Please refer to [GitHub workflow syntax](https://docs.github.com/en/free-pro-team@latest/actions/reference/workflow-syntax-for-github-actions#about-yaml-syntax-for-workflows) for more advanced usage.

> Access token is **required**. Please generate one or use `${{ secrets.GITHUB_TOKEN }}` as your access token and the `github-actions` bot will run the job for you.

## Supported File Type

Below are the currently supported configuration files that will be validated by Konfigured:

- `.ini`
- `.json`
- `.yaml`, `.yml`
- `.toml`

## Inputs

You can customize this actions with these following options (fill it on `with` section):

| **Name**              | **Required?** | **Default Value**                       | **Description**                                                                                                                                                                                                                                                                                                            |
| --------------------- | ------------- | --------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `token`        | `true`        |                                         | [GitHub access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token) to interact with the GitHub API. It is recommended to store this token with [GitHub Secrets](https://docs.github.com/en/free-pro-team@latest/actions/reference/encrypted-secrets). **To support automatic close, labeling, and comment report, please grant a write access to the token** |
| `newline` | `false` | `false` | Require all valid configuration file to end with newline.
| `include` | `false` | `**/*.{json,ini,yaml,yml,toml}` | Files to be validated in [glob pattern](https://en.wikipedia.org/wiki/Glob_(programming))

## License

This project is licensed under the [MIT License](./LICENSE)
