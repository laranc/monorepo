package ecs

import (
	"reflect"
	"sync"
	"sync/atomic"
)

// ### TYPES, AND CONSTANTS ###

type Entity uint64

type Component interface {
	Type() reflect.Type
}

type System func(ecs *ECS)

type Resource interface {
	Type() reflect.Type
}

const (
	StageUpdate = iota
	StageStartup
	stageNum
)

type ECS struct {
	nextEntity atomic.Uint64
	components map[Entity][]Component
	resources  map[reflect.Type]Resource
	systems    [stageNum][]System
	mutex      sync.RWMutex
}

// ### STARTUP FUNCTIONS ###

func NewECS() *ECS {
	var systems [stageNum][]System
	for i := range stageNum {
		systems[i] = make([]System, 0)
	}
	return &ECS{
		nextEntity: atomic.Uint64{},
		components: make(map[Entity][]Component),
		resources:  make(map[reflect.Type]Resource),
		systems:    systems,
	}
}

func (ecs *ECS) RegisterResource(resource Resource) {
	ecs.resources[resource.Type()] = resource
}

func (ecs *ECS) RegisterSystem(system System, stage uint) {
	ecs.systems[stage] = append(ecs.systems[stage], system)
}

func (ecs *ECS) RegisterDefaults() {
	// do something
}

// ### RUNTIME FUNCTIONS ###

func (ecs *ECS) CreateEntity() Entity {
	entity := Entity(ecs.nextEntity.Add(1))
	return entity
}

func (ecs *ECS) AddComponent(entity Entity, component Component) {
	ecs.mutex.Lock()
	ecs.components[entity] = append(ecs.components[entity], component)
	ecs.mutex.Unlock()
}

func (ecs *ECS) ComponentQuery(component reflect.Type, with []reflect.Type, without []reflect.Type) ([]Component, bool) {
	result := make([]Component, 0)
	ecs.mutex.RLock()
	for _, components := range ecs.components {
		hasComponent := false
		var queriedComponent Component
		for _, c := range components {
			if c.Type() == component {
				hasComponent = true
				queriedComponent = c
				break
			}
		}
		if !hasComponent {
			continue
		}
		hasWith := true
		for _, w := range with {
			foundWith := false
			for _, c := range components {
				if c.Type() == w {
					foundWith = true
					break
				}
			}
			if !foundWith {
				hasWith = false
				break
			}
		}
		if !hasWith {
			continue
		}
		hasWithout := false
		for _, w := range without {
			for _, c := range components {
				if c.Type() == w {
					hasWithout = true
					break
				}
			}
			if hasWithout {
				break
			}
		}
		if hasWithout {
			continue
		}
		result = append(result, queriedComponent)
	}
	ecs.mutex.RUnlock()
	return result, len(result) > 0
}

func (ecs *ECS) GetComponents(entity Entity) ([]Component, bool) {
	ecs.mutex.RLock()
	components, found := ecs.components[entity]
	ecs.mutex.RUnlock()
	return components, found
}

func (ecs *ECS) GetResource(r reflect.Type) (Resource, bool) {
	ecs.mutex.RLock()
	resource, found := ecs.resources[r]
	ecs.mutex.RUnlock()
	return resource, found
}

func (ecs *ECS) EntityQuery(with []reflect.Type, without []reflect.Type) ([]Entity, bool) {
	result := make([]Entity, 0)
	ecs.mutex.RLock()
	for entity, components := range ecs.components {
		entityComponents := make(map[reflect.Type]bool)
		for _, c := range components {
			entityComponents[c.Type()] = true
		}
		hasWith := true
		for _, w := range with {
			if !entityComponents[w] {
				hasWith = false
				break
			}
		}
		if !hasWith {
			continue
		}
		hasWithout := false
		for _, w := range without {
			if !entityComponents[w] {
				hasWithout = true
				break
			}
		}
		if hasWithout {
			continue
		}
		result = append(result, entity)
	}
	ecs.mutex.RUnlock()
	return result, len(result) > 0
}

// ### SYSTEM FUNCTIONS ###

func (ecs *ECS) Start() {
	for _, system := range ecs.systems[StageStartup] {
		go system(ecs)
	}
}

func (ecs *ECS) ExecuteSystems(threads int) {
	if threads == 0 {
		threads = len(ecs.systems[StageUpdate])
	}
	systems := make(chan System, len(ecs.systems[StageUpdate]))
	var wg sync.WaitGroup
	for range threads {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for system := range systems {
				system(ecs)
			}
		}()
	}
	for _, s := range ecs.systems[StageUpdate] {
		systems <- s
	}
	close(systems)
	wg.Wait()
}
