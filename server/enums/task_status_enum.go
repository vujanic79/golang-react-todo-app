package enums

type TaskStatus int

const (
	ACTIVE TaskStatus = iota
	PAUSED
	COMPLETED
	DELETED
)

var taskStatusNames = map[TaskStatus]string{
	ACTIVE:    "ACTIVE",
	PAUSED:    "PAUSED",
	COMPLETED: "COMPLETED",
	DELETED:   "DELETED",
}

func (ts TaskStatus) String() string {
	return taskStatusNames[ts]
}
