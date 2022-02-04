# prjkt

Project scaffolding made easy!

![Demo](/assets/demo.gif "Demo")

## Table of contents

- [Motivation](#motivation)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
    - [CLI Usage](#cli-usage)
      - [Create](#create)
      - [Create flags](#create-flags)
      - [List Default Actions](#list-default-actions)
    - [prjkt.yaml file](#prjktyaml-file)
      - [Example file](#example-file)
      - [Messages](#messages)
      - [Actions](#actions)

## Motivation
Heavily Inspired by [degit](https://www.npmjs.com/package/degit). I liked the concept of it and what it had for plan. But by looking at repo it seems there's no development going on with it so i decided to make this while learning go.

## Features

- Clones the repo and removes git
- Magically detects type of project and suggests automatic actions to easily setup new projects
- Repos can have a config file to customize actions done after cloning of repo

For more progress check the [Plan #1](https://github.com/patheticGeek/prjkt/issues/1)

## Installation

##### With npm (Preferred)

```
npm i -g prjkt
prjkt help
```

##### With npx
You can directly run it

```
npx prjkt help
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

### CLI usage:

#### Create

This will clone in the `my-project` folder and remove git from the project and run any actions if specified in repo or will detect the type of project and suggest some default actions.

```
prjkt create patheticGeek/pg-nextjs-boilerplate my-project
# or
prjkt c patheticGeek/pg-nextjs-boilerplate my-project
# or
prjkt c patheticGeek/pg-nextjs-boilerplate
```
The last one will create a folder called my-prjkt.

The repo can have a [`prjkt.yaml` file](#prjktyaml-file) that defines what [actions](#actions) should be done after cloning.

#### Create Flags

Create commands has some flags you can specify


- `preserve-git`: Doesn't remove git from the project
  ```
  prjkt create --preserve-git patheticGeek/pg-nextjs-boilerplate
  ```
- `no-actions`: Don't run any actions specified in prjkt.yaml
  ```
  prjkt create --no-actions patheticGeek/pg-nextjs-boilerplate
  ```
- `action`: Specify a default action, this overrides the prjkt.yaml if preset
  ```
  prjkt create --action=node patheticGeek/pg-nextjs-boilerplate
  ```
  To check available default action see [List default actions](#list-default-actions)

NOTE: These flags come before the arguments (the url and destination)

#### List default actions

Lists all the available default actions

```
prjkt list-actions
```

#### More

For more details run:
```
prjkt help
```

For more details on create run:
```
prjkt help create
```

### `prjkt.yaml` file

It's a simple file you can store in the root folder and does stuff after repo has been cloned.

### Example file

Here's what a basic config would look like for a js project

```yaml
welcome_message: "Getting things ready for ya mate"

actions:
  - 
    name: Delete readme, prjkt.yaml, lock file(s)
    type: delete
    files: README.md, prjkt.yaml, *lock*
  -
    name: Install dependencies
    type: exec-option
    prompt: Pick your poison
    options: pnpm, yarn, npm, none
    option-pnpm: pnpm i
    option-yarn: yarn
    option-npm: npm i
  -
    name: Set project name
    type: replace
    prompt: So what are we makin today?
    to_replace: <<project_name>>
    files: package.json, **/*.ts

no_actions_message: "Such a smart A$$"

error_message: "LOL this shit broke already!"

success_message: "This went smoother than it did while showing off to my friends! ✨ Enjoy"
```

For more templates/examples [look here](https://github.com/patheticGeek/prjkt-templates/tree/main/defaults)

### Messages

You can have 3 types of messages:
1. `welcome_message`: This shows up when the repo is cloned and before the actions have run
2. `error_message`: This shows if there was an error in one of the actions
3. `success_message`: This shows up if the actions run successfully
4. `no_actions_message`: If user sets the no-actions flag this will be shown instead of success/error messages

### Actions

There can be an actions array to define what to do after repo has been cloned
The following actions are present currently:

- `replace`: Replace a string in files with user input
    ```yaml
    actions:
      -
        name: Set the project name
        type: replace
        # What user sees as the question
        prompt: Name of the project
        # The string to replace
        to_replace: <<project_name>>
        # The files to set the name in
        # Can be a glob pattern or file path
        files: package.json, **/*.ts
    ```
- `delete`: Delete some files
    ```yaml
    actions:
      - 
        name: Deleting readme, prjkt.yaml, lock file(s)
        type: delete
        # Files or globe patterns to delete
        files: README.md, prjkt.yaml, *lock*
    ```
- `exec`: Execute a command
    ```yaml
    actions:
      -
        name: Run a command
        type: exec
        # Command to run, will run in the project dir
        run: echo Hello world!
        # keep going even if this command fails (default: false)
        continue: true
    ```
- `exec-option`: Execute command based on user's selection
    ```yaml
    actions:
      -
        name: Install dependencies
        type: exec-option
        # The message user sees
        prompt: Pick your poison
        # options the user can choose from
        options: pnpm, yarn, npm, none
        # option-[user choice] will be ran
        # if key is not found just continue
        option-pnpm: pnpm i
        option-yarn: yarn
        option-npm: npm i
    ```
