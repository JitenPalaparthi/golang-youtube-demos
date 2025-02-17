Task: Implement a Plugin System Using Golang Interfaces

Objective

Create a plugin system where multiple types can implement a common interface, allowing for flexible extension without modifying the core logic.

Instructions

1. Define an Interface

Create an interface called Plugin with a method Execute() string.

2. Implement Multiple Plugins

Implement at least two types (LoggerPlugin and AuthPlugin) that satisfy the Plugin interface.

LoggerPlugin should return a log message.

AuthPlugin should return an authentication status.

3. Create a Plugin Manager

Define a PluginManager struct that can register multiple plugins.

Use a slice of Plugin interface to store registered plugins.

Implement a method RunAll() that executes all registered plugins and prints their outputs.

4. Write a main function

Instantiate and register both plugins in the PluginManager.

Call RunAll() to execute all plugins.