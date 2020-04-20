# How to reproduce

- Run `cd modules/example/gql`
- Run `go run github.com/99designs/gqlgen --verbose`
- Because go.mod saved in Windows environment have special carriage return character, encounter the error: unable to find type example_gqlgen_windows_issue

# Current workaround
- To fix the error, use go.mod that is created on Mac by rename `go.mod` to something else and rename `go.mac.mod` to `go.mod` to replace.
- Run `go run github.com/99designs/gqlgen --verbose` again
- No error occurs