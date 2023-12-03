### loggap

A simple tool to mark where time passed between log lines.

### Installation

```
go install github.com/mkuratczyk/loggap@main
```

### Usage

```sh
$ tail -f foo.log | loggap 1s
2023-12-03 23:28:45.965529+01:00 [debug] <0.2771.0> User 'guest' authenticated successfully by backend rabbit_auth_backend_internal
.......... 5.043858s later
2023-12-03 23:28:51.009387+01:00 [debug] <0.2773.0> User 'guest' authenticated successfully by backend rabbit_auth_backend_internal
```

The line with dots was injected by `loggap`.
