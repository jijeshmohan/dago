# Dago

Dago is a powerful code generation tool that streamlines the development process. With Dago, you can easily generate a comprehensive file and folder structure for your projects, all while utilizing customizable templates and variables. The interactive command-line interface makes it easy to input variables and generate templates. And with its high level of configurability, Dago is perfect for developers looking to save time and increase productivity.

## Features 

- Generate codes, files and folders from templates
- Highly configurable 
- Support Variables and commands to customize templates according to the needs.
- Interactive commandline to take input variables and generate templates

## Install

### Using go 

`go install github.com/jijeshmohan/dago`

### Download 

You can download binary from the [Releases](https://github.com/jijeshmohan/dago/releases)

## Getting started

Once you have installed dago in the system, verify the installation `dago help `

### Create config file 

Dago uses configuration file to specify the templates path and other configurations, by default the path of this configuration is `~/.config/dago/dago.yaml`

We can generate this config file by using command `dago config new`

Config file has the templates directory where all the templates are located.

`dago config show` - display  current configuration in use and also the content of that config file

### Create a new template 

Lets create a hello world template which ask `author` when generating template and create file content with that name. 

`cargo template new <name-of-template>` to generate a new template 

Use `cargo template new hello-world` to get started. 

Follow the interactive answers :
```
> dago template new hello-world
? Do you want to add a variable? Yes
? Select the variable type text
? Enter the variable name author
? Enter the variable message Enter author name: 
? Enter the variable help Author name will used in the file generated
? Do you want to add a variable? No
Template created successfully
```


The template is created, we can use `dago template list` to see all the available templates. We should be able to see the template name `hello-world` in the list.

We can add any number of variables which different types which will be asked to input while generating template. In the above template we created only one variable with name author. 

These variable names can be used in the template files and folder.  

Template is a folder contain a `dago-template.yaml` file and all the files and folders which needs to be generated when using the template. The `dago-template.yaml` contains all the variables . We can directly modify this file to add more variables later. 

Lets create a file called `hello.txt` in the template folder we created now.  Default path will be `~/.config/dago/templates/hello-world`

```sh
cd ~/.config/dago/templates/hello-world
touch hello.txt
echo "Hello {{.author}}" > hello.txt 
```

We can use [golang template](https://pkg.go.dev/text/template) format for all the variables. It support even for file/folder name. e.g if we need to create a folder with author name, we can add `mkdir "{{.author}}"` in the above folder 

### Generate using template

To generate files and folder with existing template, we can use command `dago generate <template_name>`

In the above example, create a hello world 

```shell 
> dago generate hello-world /tmp
? Enter author name:  Neo
INFO  : generating template hello-world in .
INFO  : template generated successfully
```

This creates a file in the `/tmp` directory called `hello.txt`, also if we have created a directory in the template dir called `{{.author}}`, we will find `/tmp/Neo` directory as well.

```
>cat /tmp/hello.txt
Hello Neo
```


### Running commands as part of template generation

It is common to execute some commands like `git init` or package fetching as part of the template generation process. Dago supports that as well. 

The tasks can be specified in the `dago-template.yaml` file in the template. We can add `git init` as part of the above example . 

Open config file from the `hello-world` template we created.

`open ~/.config/dago/templates/hello-world/dago-template.yaml` 

Right now the content of the file will be something like this 

```yaml
name: hello-world
variables:
    - name: author
      message: 'Enter author name: '
      help: Author name will used in the file generated
      type: text
```

To add tasks, we can add a new section 

```yaml
name: hello-world
variables:
    - name: author
      message: 'Enter author name: '
      help: Author name will used in the file generated
      type: text
tasks:
    - command: "git"
      arguments:
        - "init"
```

Lets run the template again to check if it initialise git 

```sh 
> mkdir /tmp/test
> dago generate hello-world /tmp/test
? Enter author name:  Neo
INFO  : generating template hello-world in /tmp/test
INFO  : Initialized empty Git repository in /tmp/test/.git/

INFO  : template generated successfully
```

Task command line name or argument can use templates as well . e.g If we need to use `go mod init <module name>`, we can create a variable name called `module` which user input while generating template and that name can be an argument like below 

```yaml
tasks:
	- command: "go"
	  arguments: 
		  - "mod"
		  - "init"
		  - "{{.module}}"
```
