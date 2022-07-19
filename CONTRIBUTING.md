# Contributing to `cocainate`
We love your input! We want to make contributing to this project as easy and transparent as possible, whether it's:
* Reporting a bug
* Discussing the current state of the code
* Submitting a fix
* Proposing new features
* Becoming a maintainer

## [Questions/Problems](https://github.com/AppleGamer22/cocainate/discussions)
We use GitHub Discussions to answer your questions that might not require an issue ad a pull request.

Available discussion categories:
* [General](https://github.com/AppleGamer22/cocainate/discussions/categories/general)
* [Ideas](https://github.com/AppleGamer22/cocainate/discussions/categories/ideas)
  * Discussion categories changes should be suggested here.
* [Q&A](https://github.com/AppleGamer22/cocainate/discussions/categories/q-a)
* [Show & Tell](https://github.com/AppleGamer22/cocainate/discussions/categories/show-tell)

## [Submissions](https://guides.github.com/introduction/flow/index.html)
Pull requests are the best way to propose changes to the codebase (we use Github Flow). We actively welcome your pull requests:

> The following workflow uses the [GitHub CLI](https://cli.github.com/) for illustration purposes, so feel free to use other Git compatible tools that you are familiar with.

>![Gitflow Workflow](https://wac-cdn.atlassian.com/dam/jcr:61ccc620-5249-4338-be66-94d563f2843c/05%20(2).svg?cdnVersion=1393)
> Atlassian. (2021). Gitflow Workflow. Atlassian; Atlassian. https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow

‌

1. Fork the repo and create your branch from `master`/`main`.
```bash
# Clone your fork of the repo into the current directory
git clone https://github.com/AppleGamer22/cocainate
# Navigate to the newly cloned directory
cd cocainate
# Assign the original repo to a remote called "GitHub"
git remote add GitHub https://github.com/AppleGamer22/cocainate
# Create a new branch from HEAD
git branch <your branch>
```
2. If you've added code that should be tested, add tests (`.spec.ts` files).
3. If you've changed APIs, update the [documentation](https://tsdoc.org/).
4. Ensure the test suite passes.
```bash
# Run library tests
go test ./session ./ps
```
5. Commit & push your changes to your branch
```bash
# Commit all changed files with a message
git commit -am "<commit message>"
# Push changes to your remote GitHub branch
git push GitHub <your branch>
```
6. [Create a pull request](https://cli.github.com/manual/gh_pr_create)
```bash
gh pr create [flags]
```
7. [Mark a pull request as ready for review](https://cli.github.com/manual/gh_pr_ready) if appropriate
```bash
gh pr ready [<number> | <url> | <branch>] [flags]
```
8. Wait for others to review the pull request
9. [Close the pull request](https://cli.github.com/manual/gh_pr_close) if appropriate
```bash
gh pr close [<number> | <url> | <branch>] [flags]
```

## [Bug Reports/Feature Requests](https://github.com/AppleGamer22/cocainate/issues)
We use GitHub Issues to track public bugs. Report a bug by opening a new issue.

Your bug reports should have:
* A quick summary and/or background
* Steps to reproduce with:
  * Shell commands
  * Sample code
  * Configuration files
* Expected behaviour
* Actual behaviour
* Notes about:
  * including why you think this might be happening,
  * or stuff you tried that didn't work

## [Coding Style](https://editorconfig.org/)
Our `.editorconfig` file:
```toml
root = true

[*]
charset = utf-8
indent_style = tab
indent_size = 4
insert_final_newline = false
trim_trailing_whitespace = true

[*.md]
max_line_length = off
trim_trailing_whitespace = false

[*.{yml, yaml}]
indent_style = space
indent_size = 2
```
## [License](https://github.com/AppleGamer22/cocainate/blob/master/LICENSE.md)
When you submit code changes, your submissions are understood to be under the same [GNU GPL-3 License](https://choosealicense.com/licenses/gpl-3.0/) that covers the project. Feel free to contact the maintainers if that's a concern.

By contributing, you agree that your contributions will be licensed under its GNU GPL-3 License.