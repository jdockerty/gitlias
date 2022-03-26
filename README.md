# Gitlias

Swap between *aliases* so that you can `git commit` as the right author.

Set a number of aliases in a `gitlias.toml` file, e.g.

```toml
[alias]

  [alias.example]
  user = "hello"
  email = "hello@example.com"

  [alias.personal]
  user = "John Smith"
  email = "j.smith@gmail.com"
```

And simply switch between them so that your commit message have the right author

    gitlias personal

Now your commit messages will have the user `John Smith` and `j.smith@gmail.com` assigned to them.

**Note: this writes to the global git config.**