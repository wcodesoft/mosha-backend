# GitHub Actions

This project was configured in a way that you're able to run GitHub actions locally using https://github.com/nektos/act.

The suggested way of installing it is using Brew:

```bash
brew install act
```

## Running actions

To run the libraries test workflows execute:

```bash
act workflow_dispatch -j test_author_service
```

In the root of the project.

## Environment variables

To simulate the environment variables that are set in the GitHub actions, 
you can create a `.secrets` file in the root of the project with the following content:

```bash
CODACY_PROJECT_TOKEN=<CODACY_PROJECT_TOKEN>
```

And then run the actions with the `--secret-file` flag:

```bash
act workflow_dispatch -j test_author_service --secret-file .secrets
```