## Command-Line Task Manager Documentation

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
     $ go install https://github.com/julian776/Task-Manager.git
     ```

### Available Commands

The Command-Line Task Manager offers the following commands to help you manage your tasks effectively:

1. **ls** - List Tasks

   Use the `ls` command to list all your tasks. This command does not receive any options.

   Example:
   ```
   $ task-manager ls
   ```

2. **add** - Add a Task

   The `add` command allows you to add a new task to your task list. It requires a title and a description for the task.

   Syntax:
   ```
   $ task-manager add [title] [description]
   ```

   Example:
   ```
   $ task-manager add "Complete-Project" "Finish the final report and submit it by Friday."
   ```

3. **show** - View Specific Task

   The `show` command displays the details of a specific task based on its title.

   Syntax:
   ```
   $ task-manager show [title]
   ```

   Example:
   ```
   $ task-manager show "Complete-Project"
   ```

4. **[more commands]** - Coming Soon

   The Command-Line Task Manager team is actively working on introducing more commands to enhance your task management experience. Stay tuned for updates and additional functionalities.

That's it! You are now ready to manage your tasks efficiently using the Command-Line Task Manager. If you have any further questions or need assistance, please refer to the project documentation or contact our support team.