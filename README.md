CLIMON
======

A utility tool that executes arbitrary commands along with file changes, inspired by [nodemon](https://github.com/remy/nodemon).

[日本語ドキュメント](./README-ja.md)

Installation
------------

It is written on Go and can be used immediately after downloading regardless of dependency.
See Lily Strag.

Usage and Demonstration
-----------------------

![GIF of the demonstration.It can be seen that the Python script is created and executed, but is automatically executed at the same time as saving.](./img/demo.gif)

```shell
climon -path='./tmp' -- echo Changed
```

Example: When developing CLI in Go language

```shell
climon -path='./' -- go run main.go
```

Example: When developing CLI with `Create-ink-app`

```bash
create-ink-app my-app
climon -path='./' -- my-app
```

By detecting files, the specified command is automatically executed, making it convenient to develop while checking the operation.

| argument              | IsRequired | Description                                                       | default |
| --------------------- | ---------- | ----------------------------------------------------------------- | ------- |
| -path                 | Yes        | watch directory                                                   |         |
| -err                  | No         | True: Output error/False: Do not output error                     | true    |
| -stdo                 | No         | True: Output standard output/False: Do not output standard output | true    |
| -- \<command arg...\> | No         | Process to execute when the change is detected                    | true    |

Motivation
-------

This tool has been performing many times while developing the CLI of another project and wanting the automatic re-running function of `nodemon` and`Create-React-app` during the web development.is.

![Comparison between using Climon and not using it.It claims that using Climon will reduce the number of operations and make it convenient.](./img/Development-process-Before-After.png)
