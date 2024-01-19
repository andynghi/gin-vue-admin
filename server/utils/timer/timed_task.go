package timer

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type Timer interface {
	// Find all Cron
	FindCronList() map[string]*taskManager
	//Add Task method form in seconds
	AddTaskByFuncWithSecond(cronName string, spec string, fun func(), taskName string, option ...cron.Option) (cron.EntryID, error) // Add Task Func in seconds
	//Add Task interface form in seconds
	AddTaskByJobWithSeconds(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error)
	//Add tasks through function methods
	AddTaskByFunc(cronName string, spec string, task func(), taskName string, option ...cron.Option) (cron.EntryID, error)
	//Add tasks through interface methods To implement an interface trigger with Run method
	AddTaskByJob(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error)
	// Get the cron corresponding to taskName, which may be empty
	FindCron(cronName string) (*taskManager, bool)
	//Specify cron to start execution
	StartCron(cronName string)
	//Specify cron to stop execution
	StopCron(cronName string)
	// Find the specified task under the specified cron
	FindTask(cronName string, taskName string) (*task, bool)
	//Delete the specified task under the specified cron based on the id
	RemoveTask(cronName string, id int)
	//Delete the specified task under the specified cron based on taskName
	RemoveTaskByName(cronName string, taskName string)
	// Clean up the specified cronName
	Clear(cronName string)
	// Stop all cron
	Close()
}

type task struct {
	EntryID  cron.EntryID
	Spec     string
	TaskName string
}

type taskManager struct {
	corn  *cron.Cron
	tasks map[cron.EntryID]*task
}

// timer scheduled task management
type timer struct {
	cronList map[string]*taskManager
	sync.Mutex
}

// AddTaskByFunc adds a task through the function method
func (t *timer) AddTaskByFunc(cronName string, spec string, fun func(), taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].corn.AddFunc(spec, fun)
	t.cronList[cronName].corn.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// AddTaskByFuncWithSeconds uses WithSeconds to add tasks through the function method
func (t *timer) AddTaskByFuncWithSecond(cronName string, spec string, fun func(), taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	option = append(option, cron.WithSeconds())
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].corn.AddFunc(spec, fun)
	t.cronList[cronName].corn.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// AddTaskByJob adds a task through the interface method
func (t *timer) AddTaskByJob(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].corn.AddJob(spec, job)
	t.cronList[cronName].corn.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// AddTaskByJobWithSeconds adds tasks through the interface method
func (t *timer) AddTaskByJobWithSeconds(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	option = append(option, cron.WithSeconds())
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].corn.AddJob(spec, job)
	t.cronList[cronName].corn.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// FindTask obtains the cron corresponding to cronName and may be empty.
func (t *timer) FindCron(cronName string) (*taskManager, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.cronList[cronName]
	return v, ok
}

// FindTask obtains the cron corresponding to cronName and may be empty.
func (t *timer) FindTask(cronName string, taskName string) (*task, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.cronList[cronName]
	if !ok {
		return nil, ok
	}
	for _, t2 := range v.tasks {
		if t2.TaskName == taskName {
			return t2, true
		}
	}
	return nil, false
}

// FindCronList gets all task lists
func (t *timer) FindCronList() map[string]*taskManager {
	t.Lock()
	defer t.Unlock()
	return t.cronList
}

// StartCron starts the task
func (t *timer) StartCron(cromName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cromName]; ok {
		v.corn.Start()
	}
}

// StopCron stops the task
func (t *timer) StopCron(cromName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cromName]; ok {
		v.corn.Stop()
	}
}

// Remove Remove the specified task from cronName
func (t *timer) RemoveTask(cronName string, id int) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cronName]; ok {
		v.corn.Remove(cron.EntryID(id))
		delete(v.tasks, cron.EntryID(id))
	}
}

// RemoveTaskByName removes the specified task from cronName using taskName
func (t *timer) RemoveTaskByName(cronName string, taskName string) {
	fTask, ok := t.FindTask(cronName, taskName)
	if !ok {
		return
	}
	t.RemoveTask(cronName, int(fTask.EntryID))
}

// Clear clear task
func (t *timer) Clear(cronName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cronName]; ok {
		v.corn.Stop()
		delete(t.cronList, cronName)
	}
}

// Close releases resources
func (t *timer) Close() {
	t.Lock()
	defer t.Unlock()
	for _, v := range t.cronList {
		v.corn.Stop()
	}
}

func NewTimerTask() Timer {
	return &timer{cronList: make(map[string]*taskManager)}
}
