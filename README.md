# CLI Task Tracker

A simple command-line task tracker written in Go. Manage tasks, track their status, and save them to a JSON file.

## Features

- Add, update, and delete tasks
- Mark tasks as `todo`, `in-progress`, or `done`
- List all tasks or filter by status
- Save and load tasks from a JSON file

## Installation

Make sure [Go](https://golang.org/) is installed (version 1.20+ recommended). Clone the repository and build the application:

```bash
git clone <your-repo-url>
cd cli_task_tracker
go build -o task-cli main.go
````

## Usage

Run the program from the command line:

```bash
./task-cli <command> [arguments]
```

### Commands

* Add a new task: `./task-cli add "Buy groceries"`
* Update a task description: `./task-cli update 1 "Buy groceries and cook dinner"`
* Delete a task: `./task-cli delete 1`
* Mark a task as in progress: `./task-cli mark-in-progress 1`
* Mark a task as done: `./task-cli mark-done 1`
* List all tasks: `./task-cli list`
* List tasks by status: `./task-cli list todo`, `./task-cli list in-progress`, `./task-cli list done`
* Help: `./task-cli help`

## Task Properties

Each task has the following fields:

* `id`: Unique task identifier
* `description`: Task description
* `status`: Task status (`todo`, `in-progress`, `done`)
* `createdAt`: Creation timestamp
* `updatedAt`: Last update timestamp

## JSON Storage

Tasks are stored in a `tasks.json` file in the project directory. The file is automatically loaded at startup and saved on exit.

## Error Handling

* Invalid commands and missing arguments are checked
* Duplicate task descriptions are not allowed
* Status updates only accept `todo`, `in-progress`, or `done`

## Example

```bash
./task-cli add "Buy groceries"
Task added successfully (ID: 1)
./task-cli add "Cook dinner"
Task added successfully (ID: 2)
./task-cli list
1: Buy groceries [todo]
2: Cook dinner [todo]
./task-cli mark-in-progress 1
./task-cli list in-progress
1: Buy groceries [in-progress]
./task-cli mark-done 1
./task-cli list done
1: Buy groceries [done]
```

