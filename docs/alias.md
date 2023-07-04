# Setup An Alias

To create the alias "tdl" for the command "Command-Line-Task-Manager" in Linux or macOS, you can follow these steps:

1. Open a terminal: Launch the Terminal application on your Linux or macOS system. You can typically find it in the "Utilities" folder or by searching for "Terminal" in the applications.

2. Locate your shell configuration file: Different shells have different configuration files where you can define aliases. The most common ones are:
   - Bash shell: `~/.bashrc` or `~/.bash_profile`
   - Zsh shell: `~/.zshrc`
   - Fish shell: `~/.config/fish/config.fish`

   Note: If you're unsure which shell you're using, you can type `echo $SHELL` in the terminal to display the path of your current shell.

3. Open the shell configuration file: Use a text editor like Nano, Vim, or any other of your preference to open the shell configuration file. For example, if you're using Bash, you can run the following command to open `~/.bashrc`:
   ```shell
   nano ~/.bashrc
   ```

4. Define the alias: Add the following line to the configuration file, replacing "Command-Line-Task-Manager" with the actual command you want to alias:
   ```shell
   alias tdl="Command-Line-Task-Manager"
   ```

5. Save and exit the file: In Nano, press Ctrl + X, then Y, and finally Enter to save the changes.

6. Update the shell configuration: To apply the changes and make the alias available in your current terminal session, run the appropriate command based on your shell:
   - For Bash: Run `source ~/.bashrc` or `source ~/.bash_profile`.
   - For Zsh: Run `source ~/.zshrc`.
   - For Fish: The changes should take effect automatically without reloading.

7. Verify the alias: Type `tdl` in the terminal, and it should execute the command you specified in the alias.

Congratulations! You have successfully created the "tdl" alias for the "Command-Line-Task-Manager" command. Now you can use the shorter alias whenever you want to execute that command.