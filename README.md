![Travis CI](https://travis-ci.org/foomo/htpasswd.svg?branch=master)

# This is a simple utility library to manipulate htpasswd files

If you want to authenticate against a htpasswd file use something like https://github.com/abbot/go-http-auth .

## Supported hashing algorithms:

- sha-512 (requires mkpasswd to be installed)
- bcrypt

## This is what you can

Set user credentials in a htpasswd file:

```Go
file := "/tmp/demo.htpasswd"
name := "joe"
password := "secret"
err := htpasswd.SetPassword(file, name, password, htpasswd.HashBCrypt)
```

Remove a user:

```Go
err := htpasswd.RemoveUser(file, name)
```

Read user hash table:

```Go
passwords, err := htpasswd.ParseHtpasswdFile(file)
```

Have fun.
