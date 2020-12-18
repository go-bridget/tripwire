# Install

go get -u github.com/go-bridget/tripwire

# Configure

Create a `tripwire.json`:

~~~
[
  {
    "command": "./tripwire/connection.php", "arguments": ["..."]
  }
]
~~~

Run `tripwire` from the same folder.

# Writing tripwire checks

Any tripwire check must return an array of results, with the keys `key`
and `value` for each result. The results must contain at least one result.

# Features

- [x] Set config path with `-f' (default to `tripwire.json`)
- [x] Set work dir path with `-w` (default to `.` - current folder)
- [ ] Set a healthcheck timeout for running a check
- [ ] Check exit code from a check (non-zero = error?)
- [ ] Run checks in parallel to be faster
- [ ] Customize check key prefix over ENV
