# Contributing

Contributions are welcome, and they are greatly appreciated. Every little bit helps, and credit will always be given.

## Table of contents
* [Add new feature](#add-new-feature)
* [How to deploy](#how-to-deploy)
* [How to change configuration](#how-to-change-configuration)

## Add new feature

1. Clone this repo to your local  (if you haven't done this).
   ```bash
   $ git clone git@github.com:kitabisa/golang-poc.git
   ```

2. Create feature branch.

3. Make your own code. Make sure that you add unit test & pass all scenarios. If pre-commit hook detects that unit test is fail, your code wouldn't be committed.

4. Once you have finished your code, push your feature branch to `origin` remote.
   ```bash
   $ git push [remote-name] [branch-name]
   ```

   Example:
   ```bash
   $ git push origin new-great-feature
   ```

5. Make a pull request from your branch to `main` origin branch.

## How to deploy

golang-poc use our new deployment workflow and utilize Github Action as a CI/CD.

1. Deploy to dev&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;--> deploy manually using github action `Build-Push-Deploy-Dev`.
2. Deploy to staging&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;--> merge PR to `main` branch and automatically deployed to staging.
3. Deploy to uat&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;--> create a release tag and **SELECT IT** as a `pre-release` in the checkbox.
4. Deploy to production&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;--> uncheck `pre-release` from the checkbox and publish it.

## How to change configuration

Config file now is updated though `config.yaml` and `secret.yaml` in the `.infra/helm/{environment}/` directory. All changes are now should be done through git. Configurations will be stored as environment variables.

1. Change the configuration in the `.infra/helm/{environment}/config.yaml` file.
2. Any config related to credential, you should use `.infra/helm/{environment}/secret.yaml`.
3. If you want to change credential, you should **decrypt** secret file first. It will generate decrypted secret file.
4. Change your credential in decrypted secret file.
5. Then encrypt your secret file again.
6. After all your configurations have been updated, commit them.
7. And then create PR and deploy it to your desired environment.

Read [this documentation](https://www.notion.so/How-to-Setup-helm-secrets-010c234f5de848669321c700eb7f7f9c) to know more about encrypting and decrypting secret file.
