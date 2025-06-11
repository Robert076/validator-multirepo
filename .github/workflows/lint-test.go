name: CI

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

permissions:
  contents: read

jobs:
  golangci:
    name: lint
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: 1.24
      - name: Run golangci-lint
        uses: golangci/golangci-lint-action@v7
        with:
          version: v2.0

  test:
    name: test
    runs-on: ubuntu-latest
    needs: golangci
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: 1.24
      - name: Run tests
        run: go test ./...