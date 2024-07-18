package entity

import (
	"github.com/laranc/monorepo/engine/physics2d"
	"github.com/veandco/go-sdl2/sdl"
)

type Entity struct {
	self       uint64
	isActive   bool
	spriteRect sdl.Rect
	flags      uint16
}

type EntityHandler struct {
	entities        []*Entity
	entityBodyTable map[uint64]uint64
	bodyEntityTable map[uint64]uint64
	entityAnimTable map[uint64]uint64
}

func MakeEntityHandler() EntityHandler {
	return EntityHandler{
		entities:        make([]*Entity, 0),
		entityBodyTable: make(map[uint64]uint64),
		bodyEntityTable: make(map[uint64]uint64),
		entityAnimTable: make(map[uint64]uint64),
	}
}

func (h *EntityHandler) CreateEntity(bodyID, animID uint64, spriteRect sdl.Rect, flags uint16) uint64 {
	id := h.EntityCount()
	for i, entity := range h.entities {
		if !entity.isActive {
			id = uint64(i)
			break
		}
	}
	if id == h.EntityCount() {
		h.entities = append(h.entities, new(Entity))
	}
	entity := &Entity{
		self:       id,
		isActive:   true,
		spriteRect: spriteRect,
		flags:      flags,
	}
	h.entities[id] = entity
	h.entityBodyTable[id] = bodyID
	h.bodyEntityTable[bodyID] = id
	h.entityAnimTable[id] = animID
	return id
}

func (h *EntityHandler) EntityCount() uint64 {
	return uint64(len(h.entities))
}

func (h *EntityHandler) GetEntity(id uint64) *Entity {
	if id < h.EntityCount() {
		return h.entities[id]
	}
	return nil
}

func (h *EntityHandler) GetEntityFromBody(bodyID uint64) *Entity {
	entity, found := h.bodyEntityTable[bodyID]
	if found {
		return h.entities[entity]
	}
	return nil
}

func (h *EntityHandler) GetBodyID(id uint64) (uint64, bool) {
	body, found := h.entityBodyTable[id]
	return body, found
}

func (h *EntityHandler) GetAnimID(id uint64) (uint64, bool) {
	anim, found := h.entityAnimTable[id]
	return anim, found
}

func (h *EntityHandler) DestroyEntity(physicsState physics2d.PhysicsState, id uint64) {
	entity := h.GetEntity(id)
	entity.isActive = false
	bodyID, found := h.entityBodyTable[id]
	if found {
		physicsState.DestroyBody(bodyID)
	}
}
