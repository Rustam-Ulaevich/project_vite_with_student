package todo

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) {
	l.task[task.Title] = task
}

func (l *List) ListTasks() map[string]Task {
	return l.tasks
}