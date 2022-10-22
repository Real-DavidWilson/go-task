package main

type Task struct {
	TaskHandler func()
	Proprity    int
}

// type TaskHandler func()

func AddTask(task *Task) {
	pool.push(*task)
}

func (task *Task) Run() {
	task.TaskHandler()
}
