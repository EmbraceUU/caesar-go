package conf

var engine = &Engine{
	updated: make(chan struct{}),
}

type Engine struct {
	observers []Observer
	updated   chan struct{}
	config    *Config
}

func InitEngine() {
	engine = new(Engine)
	engine.updated = make(chan struct{})
	engine.config = initConfig(engine.updated)

	go func() {
		for {
			<-engine.updated
			for _, observer := range engine.observers {
				observer.Update(engine.config)
			}
		}
	}()
}

func GetConfig() *Config {
	return engine.config
}

func PushObserver(observer Observer) {
	engine.observers = append(engine.observers, observer)
}
