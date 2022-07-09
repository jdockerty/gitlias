# Gitlias

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jdockerty/gitlias?style=plastic)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/jdockerty/gitlias?style=plastic)

Swap between `git` aliases so that you can `git commit` as the right author.

Set a number of aliases in a `gitlias.toml` file. For example, you might switch between a `work` and `personal` alias when committing to your own projects on a lunch break. This saves you the hassle of using

  `git config --set [--global]`

or other such lines and provides a more familiar name to refer to them by.

## Install

The simplest way to install is by using Go

    go install github.com/jdockerty/gitlias@v0.2.2

Alternatively, you can use the provided [releases](https://github.com/jdockerty/gitlias/releases) to download a tarball or zip file.

## Usage

You can see the full usage and supported flags by running `gitlias --help`. Below is a short example of how the program should be utilised.

Using `gitlias init` you can generate a skeleton configuration file which is written to `${HOME}/gitlias.toml`.

You can add various aliases using `gitlias add --alias <alias> --user <user> --email <email_address>`, this will populate your configuration file.


Once you have added some aliases, your file may look like this.

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
gitlias switch work
gitlias switch personal # current active alias
```

Now your commit messages will have the user `John Smith` and email `j.smith@gmail.com` assigned to them.

**Note: this currently writes to the global git config.**


You can view all configured aliases using `gitlias list` and the current one with `gitlias list --current`.

If you no longer wish to use an alias, you can remove is using `gitlias rm <alias_name>`.

