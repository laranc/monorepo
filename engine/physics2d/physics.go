package physics2d

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/laranc/monorepo/engine/global"
)

type OnHit func(self *Body, other *Body, hit Hit)

type OnHitStatic func(self *Body, other *StaticBody, hit Hit)

type AABB struct {
	position mgl32.Vec2
	halfSize mgl32.Vec2
}

type Body struct {
	aabb           AABB
	velocity       mgl32.Vec2
	acceleration   mgl32.Vec2
	collisionLayer uint8
	collisionMask  uint8
	onHit          OnHit
	onHitStatic    OnHitStatic
	isKinematic    bool
	isActive       bool
	self           uint64
}

type StaticBody struct {
	aabb           AABB
	collisionLayer uint8
	self           uint64
}

type Hit struct {
	isHit    bool
	time     float64
	position mgl32.Vec2
	normal   mgl32.Vec2
	other    uint64
}

type PhysicsState struct {
	gravity          float32
	terminalVelocity float32
	bodies           []*Body
	staticBodies     []*StaticBody
}

const (
	iterations = 2
	tickRate   = 1.0 / iterations
)

func MakePhysicsState() PhysicsState {
	return PhysicsState{
		gravity:          -79,
		terminalVelocity: -7000,
		bodies:           make([]*Body, 0),
		staticBodies:     make([]*StaticBody, 0),
	}
}

func (state *PhysicsState) Update() {
	for _, body := range state.bodies {
		if !body.isActive {
			continue
		}
		if !body.isKinematic {
			body.velocity[1] += state.gravity
			if state.terminalVelocity > body.velocity[1] {
				body.velocity[1] = state.terminalVelocity
			}
		}
		body.velocity = body.velocity.Add(body.acceleration)
		scaledVelocity := body.velocity.Mul(float32(global.State.Time.Delta) * tickRate)
		for range iterations {
			state.sweepResponse(body, scaledVelocity)
			state.stationaryResponse(body)
		}
	}
}

// Constructors

func (state *PhysicsState) CreateBody(position, size, velocity mgl32.Vec2, collisionLayer, collisionMask uint8, onHit OnHit, onHitStatic OnHitStatic, isKinematic, isActive bool) uint64 {
	id := state.BodyCount()
	for i, body := range state.bodies {
		if !body.isActive {
			id = uint64(i)
			break
		}
	}
	if id == state.BodyCount() {
		state.bodies = append(state.bodies, new(Body))
	}
	body := state.GetBody(id)
	*body = Body{
		aabb: AABB{
			position: position,
			halfSize: mgl32.Vec2{size[0] / 2, size[1] / 2},
		},
		velocity:       velocity,
		collisionLayer: collisionLayer,
		collisionMask:  collisionMask,
		onHit:          onHit,
		onHitStatic:    onHitStatic,
		isKinematic:    isKinematic,
		isActive:       isActive,
		self:           id,
	}
	return id
}

func (state *PhysicsState) CreateStaticBody(position, size mgl32.Vec2, collisionLayer uint8) uint64 {
	id := state.StaticBodyCount()
	staticBody := &StaticBody{
		aabb: AABB{
			position: position,
			halfSize: mgl32.Vec2{size[0] / 2, size[1] / 2},
		},
		collisionLayer: collisionLayer,
		self:           id,
	}
	state.staticBodies = append(state.staticBodies, staticBody)
	return id
}

func (state *PhysicsState) CreateTrigger(position, size mgl32.Vec2, collisionLayer, collisionMask uint8, onHit OnHit) uint64 {
	return state.CreateBody(position, size, mgl32.Vec2{0, 0}, collisionLayer, collisionMask, onHit, nil, true, true)
}

func (state *PhysicsState) DestroyBody(id uint64) {
	body := state.GetBody(id)
	body.isActive = false
}

// Getters

func (state *PhysicsState) GetBody(id uint64) *Body {
	if id <= uint64(len(state.bodies)) {
		return state.bodies[id]
	}
	return nil
}

func (state *PhysicsState) GetStaticBody(id uint64) *StaticBody {
	if id <= uint64(len(state.staticBodies)) {
		return state.staticBodies[id]
	}
	return nil
}

func (state *PhysicsState) BodyCount() uint64 {
	return uint64(len(state.bodies))
}

func (state *PhysicsState) StaticBodyCount() uint64 {
	return uint64(len(state.staticBodies))
}

// Math

func PointIntersectAABB(point mgl32.Vec2, aabb AABB) bool {
	min, max := AABBMinMax(aabb)
	return point[0] >= min[0] && point[0] <= max[0] && point[1] >= min[1] && point[1] <= max[1]
}

func AABBIntersectAABB(a, b AABB) bool {
	min, max := AABBMinMax(AABBMinkowskiDifference(a, b))
	return min[0] <= 0 && max[0] >= 0 && min[1] <= 0 && max[1] >= 0

}

func AABBMinkowskiDifference(a, b AABB) AABB {
	var result AABB
	result.position = a.position.Sub(b.position)
	result.halfSize = a.halfSize.Add(b.halfSize)
	return result
}

func AABBPenetrationVector(aabb AABB) mgl32.Vec2 {
	var r mgl32.Vec2
	min, max := AABBMinMax(aabb)
	minDist := mgl32.Abs(min[0])
	r[0] = min[0]
	r[1] = 0

	if mgl32.Abs(max[0]) < minDist {
		minDist = mgl32.Abs(max[0])
		r[0] = max[0]
	}
	if mgl32.Abs(min[1]) < minDist {
		minDist = mgl32.Abs(min[1])
		r[0] = 0
		r[1] = min[1]
	}
	if mgl32.Abs(max[1]) < minDist {
		minDist = mgl32.Abs(max[1])
		r[0] = 0
		r[1] = min[1]
	}
	return r
}

func AABBMinMax(aabb AABB) (min, max mgl32.Vec2) {
	min = aabb.position.Sub(aabb.halfSize)
	max = aabb.position.Add(aabb.halfSize)
	return min, max
}

func RayIntersectAABB(position, magnitude mgl32.Vec2, aabb AABB) Hit {
	hit := Hit{}
	min, max := AABBMinMax(aabb)
	lastEntry := float32(math.Inf(-1))
	firstExit := float32(math.Inf(1))
	for i := range 2 {
		if magnitude[i] != 0 {
			t1 := (min[i] - position[i]) / magnitude[i]
			t2 := (max[i] - position[i]) / magnitude[i]
			lastEntry = float32(math.Max(float64(lastEntry), math.Min(float64(t1), float64(t2))))
			firstExit = float32(math.Min(float64(firstExit), math.Max(float64(t1), float64(t2))))
		} else if position[i] <= min[i] || position[i] >= max[i] {
			return hit
		}
	}
	if firstExit > lastEntry && firstExit > 0 && lastEntry < 1 {
		hit.position[0] = position[0] + magnitude[0]*lastEntry
		hit.position[1] = position[1] + magnitude[1]*lastEntry
		hit.isHit = true
		hit.time = float64(lastEntry)
		dx := hit.position[0] - aabb.position[0]
		dy := hit.position[1] - aabb.position[1]
		px := aabb.halfSize[0] - mgl32.Abs(dx)
		py := aabb.halfSize[1] - mgl32.Abs(dx)
		if px < py {
			if dx > 0 {
				hit.normal[0] = 1
			} else if dx < 0 {
				hit.normal[0] = -1
			} else {
				hit.normal[0] = 0
			}
		} else {
			if dy > 0 {
				hit.normal[1] = 1
			} else if dy < 0 {
				hit.normal[1] = -1
			} else {
				hit.normal[1] = 0
			}

		}
	}

	return hit
}

// Internal

func (state *PhysicsState) updateSweepResult(result *Hit, body *Body, otherID uint64, velocity mgl32.Vec2) {
	other := state.GetBody(otherID)
	if (body.collisionMask & other.collisionLayer) == 0 {
		return
	}
	sum := other.aabb
	sum.halfSize = sum.halfSize.Add(body.aabb.halfSize)
	hit := RayIntersectAABB(body.aabb.position, velocity, sum)
	if hit.isHit {
		if body.onHit != nil && (body.collisionMask&other.collisionLayer) == 0 {
			body.onHit(body, other, hit)
		}
		if hit.time < result.time {
			*result = hit
		} else if hit.time == result.time {
			if mgl32.Abs(velocity[0]) > mgl32.Abs(velocity[1]) && hit.normal[0] != 0 || mgl32.Abs(velocity[1]) > mgl32.Abs(velocity[0]) && hit.normal[1] != 0 {
				*result = hit
			}
		}
		result.other = otherID
	}
}

func (state *PhysicsState) updateSweeResultStatic(result *Hit, body *Body, otherID uint64, velocity mgl32.Vec2) {
	other := state.GetStaticBody(otherID)
	if (body.collisionMask & other.collisionLayer) == 0 {
		return
	}
	sum := other.aabb
	sum.halfSize = sum.halfSize.Add(body.aabb.halfSize)
	hit := RayIntersectAABB(body.aabb.position, velocity, sum)
	if hit.isHit {
		if hit.time < result.time {
			*result = hit
		} else if hit.time == result.time {
			if mgl32.Abs(velocity[0]) > mgl32.Abs(velocity[1]) && hit.normal[0] != 0 || mgl32.Abs(velocity[1]) > mgl32.Abs(velocity[0]) && hit.normal[1] != 0 {
				*result = hit
			}
		}
		result.other = otherID
	}
}

func (state *PhysicsState) sweepBodies(body *Body, velocity mgl32.Vec2) Hit {
	result := Hit{time: math.Inf(1)}
	for i, other := range state.bodies {
		if body == other {
			continue
		}
		state.updateSweepResult(&result, body, uint64(i), velocity)
	}
	return result
}

func (state *PhysicsState) sweepStaticBodies(body *Body, velocity mgl32.Vec2) Hit {
	result := Hit{time: math.Inf(1)}
	for i := range state.bodies {
		state.updateSweeResultStatic(&result, body, uint64(i), velocity)
	}
	return result
}

func (state *PhysicsState) sweepResponse(body *Body, velocity mgl32.Vec2) {
	hit := state.sweepStaticBodies(body, velocity)
	hitMoving := state.sweepBodies(body, velocity)
	if hitMoving.isHit && body.onHit != nil {
		body.onHit(body, state.GetBody(hitMoving.other), hitMoving)
	}

	if hit.isHit {
		body.aabb.position = hit.position
		if hit.normal[0] != 0 {
			body.aabb.position[1] += velocity[1]
			body.velocity[0] = 0
		} else if hit.normal[1] != 0 {
			body.aabb.position[0] += velocity[0]
			body.velocity[1] = 0
		}
		if body.onHitStatic != nil {
			body.onHitStatic(body, state.GetStaticBody(hit.other), hit)
		}
	} else {
		body.aabb.position.Add(velocity)
	}
}

func (state *PhysicsState) stationaryResponse(body *Body) {
	for _, staticBody := range state.staticBodies {
		aabb := AABBMinkowskiDifference(staticBody.aabb, body.aabb)
		min, max := AABBMinMax(aabb)
		if min[0] <= 0 && max[0] >= 0 && min[1] <= 0 && max[1] >= 0 {
			penetrationVector := AABBPenetrationVector(aabb)
			body.aabb.position = body.aabb.position.Add(penetrationVector)
		}
	}
	for id, other := range state.bodies {
		if body.onHit == nil {
			continue
		}
		if (body.collisionMask & other.collisionLayer) == 0 {
			continue
		}

		aabb := AABBMinkowskiDifference(other.aabb, body.aabb)
		min, max := AABBMinMax(aabb)
		if min[0] <= 0 && max[0] >= 0 && min[1] <= 0 && max[1] >= 0 {
			body.onHit(body, other, Hit{isHit: true, other: uint64(id)})
		}
	}
}
