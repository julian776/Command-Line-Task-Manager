# Command-Line Task Manager

The Command-Line Task Manager is a tool designed to help you manage your tasks efficiently using simple commands from the command line interface.

### Installation

There are two ways to install the Command-Line Task Manager:

1. **Download and Environment Variable Setup**

   - Download the repository from [GitHub](https://github.com/julian776/Task-Manager.git).
   - Extract the contents of the repository to your desired location.
   - Set up an environment variable pointing to the extracted location.

2. **Installation with Go**

   - Ensure you have [Go](https://golang.org/) installed on your system.
   - Open the command line interface and execute the following command:
     ```
     $ go install github.com/julian776/Command-Line-Task-Manager
     ```

### Alias (Optional)
Creating an alias can significantly simplify the usage of the "Command-Line-Task-Manager" command by providing a shorter and more intuitive alternative. In this case, the chosen alias "tdl" is derived from the term "to-do list," which aligns with the purpose of the task manager.

By using the alias "tdl" you can conveniently execute the "Command-Line-Task-Manager" command with just three keystrokes, saving time and effort compared to typing the entire command each time.

Check how to setup an alias [here](./docs/alias.md)

### Available Commands

The Command-Line Task Manager offers the following commands to help you manage your tasks effectively:

1. **ls** - List Tasks

   Use the `ls` command to list all your tasks. This command does not receive any options.

   Example:
   ```
   $ Command-Line-Task-Manager ls
   ```

2. **add** - Add a Task

   The `add` command allows you to add a new task to your task list. It requires a title and a description for the task.

   Syntax:
   ```
   $ Command-Line-Task-Manager add [title] [description]
   ```

   Example:
   ```
   $ Command-Line-Task-Manager add "Complete-Project" "Finish the final report and submit it by Friday."
   ```

3. **show** - View Specific Task

   The `show` command displays the details of a specific task based on its title.

   Syntax:
   ```
   $ Command-Line-Task-Manager show [title]
   ```

   Example:
   ```
   $ Command-Line-Task-Manager show "Complete-Project"
   ```

4. **done** - Mark Task as Done

   The `done` command allows you to mark a task as completed. Specify the title of the task you want to mark as done.

   Syntax:
   ```
   $ Command-Line-Task-Manager done [title]
   ```

   Example:
   ```
   $ Command-Line-Task-Manager done "Complete-Project"
   ```

5. **[more commands]** - Coming Soon