# Gitlias

Swap between `git` aliases so that you can `git commit` as the right author.

Set a number of aliases in a `gitlias.toml` file. For example, you might switch between a `work` and `personal` alias when committing to your own projects on a lunch break.

## Install

The simplest way to install is by using Go

    go install github.com/jdockerty/gitlias@latest

## Usage

You can see the full usage and supported flags by running `gitlias --help`. Below is a short example of how the program should be utilised.

```toml
# ${HOME}/gitlias.toml
[alias]

  [alias.work]
  user = "John S"
  email = "john@example.com"

  [alias.personal]
  user = "John Smith"
  email = "j.smith@example.com"
```

Switching between them so that your commit messages have the corresponding author

```bash
gitlias work
gitlias personal # current active alias
```

Now your commit messages will have the user `John Smith` and email `j.smith@gmail.com` assigned to them.

**Note: this currently writes to the global git config.**


## TODO

- [ ] add tests
- [ ] add CI