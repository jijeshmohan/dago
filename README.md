# Daco

Daco is a code generation utility which uses templates

## Features 

- Generate files and folder structure for project
- Support Variables
- Interactive commandline tool to take input variables and generate templates
- Highly configurable

## Install

TODO:

## Getting started 

1. Create a config directory and specify templates directory

    `dago config new` - This command create a new config file in user's config directory `~/.config/dago/dago.yaml`

1. Specify the templates folder in the above config file ( default is `~/.config/dago/templates`)


### How to create a new template 

1. Create a new template `daco new template <name>`

    this will create a new template folder in the configured project templates directory

2. Fill the config file for the templates with variables and tasks 

    Template consist of a `daco-template.yml` config file and then files and folders needs to be generated by that template 

    Variables can be configured which will be asked while generating the template 

    Sample variable configuration
    ```yaml
     - name: "project_name"
       type: "text"
       message: "Enter project name"
       help: "project name will be used to create the app"
       validators:
         - "required"
       transformer: "lower"
    ```
    Variable support different types which are described detailed in the variable section

3. Create the folder structure and files, make use of the golang template syntax. 
    All the variables defined in the template will be available for render

### Generate from template  

1. `dago generate <existing-template-name> <path-where-to-generate>`

2. dago will prompt users to enter values for all the variables configures in the template 

3. dago render all the files and folders from the template folder 

4. dago run tasks specified in the template if any 

### Template Config file  (daco-template.yaml)

The basic structure of the document is like this 

```yaml
name: template_name
variables:
    - name: variable_name
      type: text
tasks:
    - command: "unix_command_to_run"
      arguments:
        - arg1
        - arg2
```


## Commands 

- config
    - new
    - show
- template
    - new
    - validate
- gen
