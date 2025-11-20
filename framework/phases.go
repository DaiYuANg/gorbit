package framework

type Phase string

const (
	PhaseRegister Phase = "register"
	PhaseInit     Phase = "init"
	PhaseStart    Phase = "start"
	PhaseStop     Phase = "stop" // 以后扩展用
)

func (p Phase) String() string {
	return string(p)
}

type Action string

const (
	ActionRegistering  Action = "Registering"
	ActionRegistered   Action = "Registered"
	ActionInitializing Action = "Initializing"
	ActionInitialized  Action = "Initialized"
	ActionStarting     Action = "Starting"
	ActionStarted      Action = "Started"
	ActionStopping     Action = "Stopping"
	ActionStopped      Action = "Stopped"
)

func (a Action) String() string {
	return string(a)
}

type PhaseConfig struct {
	Phase      Phase
	EventStart string
	EventDone  string
	Action     [2]Action // [开始动作, 完成动作]
	Handler    func(f *Framework, m Module) error
}

var defaultPhases = []PhaseConfig{
	{
		Phase:      PhaseRegister,
		EventStart: EventFrameworkRegisterStart,
		EventDone:  EventFrameworkRegisterDone,
		Action:     [2]Action{ActionRegistering, ActionRegistered},
		Handler: func(f *Framework, m Module) error {
			return m.Register(f.injector)
		},
	},
	{
		Phase:      PhaseInit,
		EventStart: EventFrameworkInitStart,
		EventDone:  EventFrameworkInitDone,
		Action:     [2]Action{ActionInitializing, ActionInitialized},
		Handler: func(f *Framework, m Module) error {
			return m.Init(f.ctx)
		},
	},
	{
		Phase:      PhaseStart,
		EventStart: EventFrameworkStart,
		EventDone:  EventFrameworkStartDone,
		Action:     [2]Action{ActionStarting, ActionStarted},
		Handler: func(f *Framework, m Module) error {
			return m.Start(f.ctx)
		},
	},
}
