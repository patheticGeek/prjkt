# prjkt

Project scaffolding made easy!

## Motivation
Heavily Inspired by [degit](https://www.npmjs.com/package/degit). I liked the concept of it and what it had for plan. But by looking at repo it seems there's no development going on with it so i decided to make this while learning go.

[See this issue for the progress](https://github.com/patheticGeek/prjkt/issues/1)

## Installation

#### With npm (Preferred)
```
npm i -g prjkt-installer
```

<details>
<summary>By cloning the repo</summary>

1. Clone the repo
    ```
    git clone https://github.com/patheticGeek/prjkt.git
    ```
2. Install
    ```
    go install .
    ```
3. Test if it's working with
    ```
    prjkt help
    ```
</details>

## Usage:

Basic usage:
```
prjkt c -u patheticGeek/pg-nextjs-boilerplate -d my-new-project
```
This will clone and remove git from the project and run any actions if specified

More options for create command:
```
prjkt help create
```

You can also have a `prjkt.yaml` file in your repo that defined what should be done after cloning.

For details about it, [see here](./prjkt.yaml-file).
