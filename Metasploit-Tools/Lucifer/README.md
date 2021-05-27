# Lucifer

[![ForTheBadge made-with-python](http://ForTheBadge.com/images/badges/made-with-python.svg)](https://www.python.org/)
[![ForTheBadge built-with-love](http://ForTheBadge.com/images/badges/built-with-love.svg)](https://GitHub.com/Skiller9090/)
[![GitHub repo size](https://img.shields.io/github/repo-size/Skiller9090/Lucifer?color=CC0000&style=for-the-badge)]()

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/066ed76111cd49d98ba5dad52240b2bd)](https://app.codacy.com/manual/alexspam9090/Lucifer?utm_source=github.com&utm_medium=referral&utm_content=Skiller9090/Lucifer&utm_campaign=Badge_Grade_Dashboard)
[![Maintainability](https://api.codeclimate.com/v1/badges/4cf141dbb9b75910c663/maintainability)](https://codeclimate.com/github/Skiller9090/Lucifer/maintainability)
[![CodeFactor](https://www.codefactor.io/repository/github/skiller9090/lucifer/badge)](https://www.codefactor.io/repository/github/skiller9090/lucifer)
[![GitHub release](https://img.shields.io/github/release/Skiller9090/Lucifer.svg)](https://GitHub.com/Skiller9090/Lucifer/releases/)
[![GitHub commits](https://img.shields.io/github/commits-since/Skiller9090/Lucifer/latest)](https://GitHub.com/Skiller9090/Lucifer/commit/)
[![GitHub contributors](https://img.shields.io/github/contributors/Skiller9090/Lucifer)](https://GitHub.com/Skiller9090/Lucifer/graphs/contributors/)
[![GitHub issues](https://img.shields.io/github/issues/Skiller9090/Lucifer)](https://GitHub.com/Skiller9090/Lucifer/issues/)
[![Open Source? Yes!](https://badgen.net/badge/Open%20Source%20%3F/Yes%21/blue?icon=github)](https://github.com/Skiller9090/badges/)

[![CodeScene Code Health](https://codescene.io/projects/9260/status-badges/code-health)](https://codescene.io/projects/9260)
[![CodeScene System Mastery](https://codescene.io/projects/9260/status-badges/system-mastery)](https://codescene.io/projects/9260)

[![Language grade: Python](https://img.shields.io/lgtm/grade/python/g/Skiller9090/Lucifer.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/Skiller9090/Lucifer/context:python)

`Initial Release`: 09/09/2020

## Setup

```shell
git clone https://github.com/Skiller9090/Lucifer.git
cd Lucifer
pip install -r requirements.txt
python main.py --help
```

If you want the cutting edge changes add `-b dev` to the end of `git clone https://github.com/Skiller9090/Lucifer.git`

## Commands

| Command            | Description                                                             |
| ------------------ | ----------------------------------------------------------------------- |
|help                | Displays This Menu                                                      |
|name                | Shows name of current shell                                             |
|id                  | Displays current shell's id                                             |
|show                | Shows options or modules based on input, EX: show <options/modules>     | 
|options             | Shows a list of variable/options already set                            |
|set                 | Sets a variable or option, EX: set <var> <data>                         |
|set_vars            | Auto sets need variables for loaded module                              |
|description         | Displays description of the module loaded                               |
|auto_vars           | Displays is auto_vars is True or False for current shell                |  
|change_auto_vars    | Changes the auto_var options for one shell, all shells or future shells | 
|reindex             | Re-indexes all modules, allows for dynamic additions of modules         |
|use                 | Move into a module, EX: use <module>                                    |
|run                 | Runs the current module, can also use exploit to do the same            | 
|spawn_shell         | Spawns a alternative shell                                              |
|open_shell          | Open a shell by id EX: open_shell <id>                                  |
|show_shells         | Show all shell ids and attached name                                    |
|set_name            | Sets current shells name EX: set_name <name>                            |
|set_name_id         | Set a shells name by id EX: set_name_id <id> <name>                     | 
|clear               | Clear screen                                                            |
|close               | Kills current input into opened shell                                   |
|reset               | Resets Everything                                                       |
|exit                | Exits the program, can also use quit to do the same                     |  

### Command Use

- #### No-Arg Commands
    - `help` - to display help menu

    - `name` - shows name of current shell

    - `id` - shows current shell id

    - `options` - shows a table of set options/vars

    - `set_vars` - automatically sets vars needed for the loaded module (default defined in a module)

    - `description` - show description of current loaded module

    - `auto_vars` - displays current setting of auto_vars (auto_vars if true will automatically run set_vars on module
      load)

    - `run` - runs the module with the current options, `exploit` works the same

    - `spawn_shell` - spawns a new Shell instance

    - `show_shells` - shows all open shells ids and names

    - `clear` - clears the terminal/console screen

    - `close` - kills the input to current shell

    - `reset` - resets everything (not implemented)

    - `exit` - quits the program


- #### Arg Commands
    - `show <options/modules>` - displays a list of set options or modules depending on argument.

    - `set <var_name> <value>` - sets a variable/option

    - `change_auto_vars <to_set> <args>`:
        - `<to_set>` - can be true or false (t or f) (-t or -f)

        - `<args>`:
            - `-g` = global - sets for all shells spawned

            - `-n` = new - sets this option for future shell spawns

            - `-i` = inclusive - no matter what, set current shell to <to_set>

    - `use <module> <args>`:
        - `<module>` - path to module

        - `<args>`:
            - `-R` - Override cache (reload dynamically)

    - `open_shell <id>` - opens a shell by its id

    - `set_name <name>` - set the name of the current shell

    - `set_name_id <id> <name>` - set the name of the shell specified by <id>

# Using Java

Lucifer allows for Python and Java code to work side by side through the use of LMI.Java extension. For this to work you
will need to install jpype1, to do this run the following command in your python environment:  
`pip install jpype1`  
From here you are free to interact with LMI.Java.compiler and LMI.Java.luciferJVM which allows you to call java
functions and instantiate java classes through python, more documentation of this will be created later on, on the
lucifer wiki.

## Examples

### Settings Variables

![example two](./assets/examples/two.png)

### Running Module

![example one](./assets/examples/one.png)

### Settings

![example three](./assets/examples/three.png)

## Versioning

The standard of versioning on this project is:

### MAJOR.MINOR.PATCH.STAGE.BUILD

#### Major:

- incremented when either there has been a significant amount of new features since the start of the major or if there
  is a change which is so big that is can cause compatibility issues (Major of 0 if very unstable
- Could cause incompatibility issues

#### Minor:

- incremented when a new feature or feature-set is added to the project
- should not cause incompatibility errors due to only additions made

#### Patch:

- incremented on bugfixes or if feature is so small that it is worth incrementing minor
- very low risk of incompatibility error

#### Stage:

- The stage of current MAJOR.MINOR.PATCH BUILD, either alpha, beta, release candidate or release
- Indicates how far through development the new MAJOR.MINOR.PATCH is
- Stage number to name translation:
    - 0 => beta (b)
    - 1 => alpha (a)
    - 2 => release candidate (rc)
    - 3 => release (r)

#### Build:

- this should be incremented on every change made to the code, even on a one character change

This version structure can be stored and displayed in a few ways:

- The best way to store the data within code is via a tuple such as:
    - (Major, Minor, Patch, Stage, Build)
        - Example is: (1, 4, 1, 2, 331)
- The long display would be:
    - `{stage} {major}.{minor}.{patch} Build {build}`
        - Example is: Alpha 1.4.1 Build 331
- The short display would be:
    - `{major}.{minor}.{patch}{stage}{build}`
        - Example is: 1.4.1a331
  
