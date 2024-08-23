# Task Tracker

This proyect was maded as roadmap.sh [Proyect](<https://roadmap.sh/projects/task-tracker>)
## Info

This cli-app will manage your task and save it in a JSON file located in ./Data/task.json


## Build

~~~
go build -o dist/task-cli.exe main.go
~~~

I chosse dist, because i like the way Angular manage his builds.

## Commands
~~~
# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
~~~
