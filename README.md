# golang-poc

## Table of contents
- [golang-poc](#-cookiecutterapp_name-)
  - [Table of contents](#table-of-contents)
  - [How to contribute](#how-to-contribute)
  - [Requirement](#requirement)
  - [How to configure on your local machine](#how-to-configure-on-your-local-machine)
  - [How to run migration](#how-to-run-migration)
  - [How to add dummy data to database](#how-to-add-dummy-data-to-database)
  - [Unit test](#unit-test)
  - [Enable unit test threshold checking (on-merge policy)](#enable-unit-test-threshold-checking-on-merge-policy)
  - [Need more help](#need-more-help)

## How to contribute

Read [CONTRIBUTING.md](https://github.com/kitabisa/golang-poc/blob/main/CONTRIBUTING.md) to know more about how to contribute to this repo and how to deploy service. If you are new, it is mandatory to read this file first.

## Requirement

1. Go 1.13 or above (download [here](https://golang.org/dl/)).
2. Git (download [here](https://git-scm.com/downloads)).
3. Redis (download [here](https://hub.docker.com/_/redis)).

## How to configure on your local machine

1. Clone this repostory to your local.
   ```bash
   $ git clone git@github.com:kitabisa/golang-poc.git
   ```

2. Change working directory to `golang-poc` folder.
   ```bash
   $ cd golang-poc
   ```

3. Create configuration files.
   ```bash
   $ cp params/.env.example params/.env
   ```

4. Edit configuration values in `params/.env` according to your setting.
5. Run `go get` and `make build`
```
go get && make build
```
6. Run docker-compose
```
make run-compose
```

## How to run migration

This migration can do these actions:

1. Migration up

   This command will migrate the database to the most recent version available. Migration files can be seen in this folder `migrations/sql/`.
   ```bash
   $ go run main.go migrate
   ```

2. Migration down

   This command will undo/rollback database migration.
   ```bash
   $ go run main.go migratedown
   ```

3. Migration new

   This command will generate new migration file based on unix timestamp format
   ```bash
   $ go run main.go migratenew create_table_A
   2019/11/28 13:08:01 New migration file has been created: migrations/sql/1574921281_create_table_A.sql
   ```

To get any help about these command, add `--help` at the end of the command.
```bash
$ go run main.go --help
$ go run main.go migrate --help
$ go run main.go migratedown --help
$ go run main.go migratenew --help
```

## How to add dummy data to database

Just import any `*.sql` seeds data from this folder `migration/seeds/`.

## Unit test

Use the following commands to run unit test and to see code coverage.

```bash
$ make test
$ make coverage
```

## Enable unit test threshold checking (on-merge policy)

You can enable unit test coverage threshold checking on PR created. If your branch unit test is lower than threshold, `build-test` action will fail and merge PR button will be disabled.

Here are the steps to activate this github action:

1. Enable branch protection rule from repository setting for main or master branch. This will disable merge PR if there is fail CI job.

<img width="983" alt="Screen Shot 2021-01-21 at 10 21 45" src="https://user-images.githubusercontent.com/9508513/105275695-79397300-5bd2-11eb-8295-b061b7274f97.png">

2. Add your test coverage threshold via github action on `.github/workflows/build.yaml` file. Setting your value 0 will disable unit test coverage checking.

```
env:
  COVERAGE_THRESHOLD: 75
```

3. Any result regarding your unit test will be reported on the PR's comment.

## Need more help

If you need more help or anything else, please ask Kitabisa backend engineer team. We would be happy to help you.
