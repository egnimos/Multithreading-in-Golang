package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func (b *Boid) moveOne() {
	// update the velocity after calculating the acceleration
	accel := b.calcAcceleration()
	// lock the code for shared memory
	lock.Lock()
	b.velocity = b.velocity.Add(accel)
	b.velocity = b.velocity.Limit(-1, 1)

	// update the boidIdMatrix position after changing the position and resetting the position
	boidIdMatrix[int(b.position.x)][int(b.position.y)] = -1
	//change the position
	b.position = b.position.Add(b.velocity)
	boidIdMatrix[int(b.position.x)][int(b.position.y)] = b.id
	// check for new position
	// next := b.position.Add(b.velocity)
	// if next.x >= screenWidth || next.x < 0 {
	// 	b.velocity = Vector2D{-b.velocity.x, b.velocity.y}
	// }
	// if next.y >= screenHeight || next.y < 0 {
	// 	b.velocity = Vector2D{b.velocity.x, -b.velocity.y}
	// }
	lock.Unlock()
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxBorderPos-viewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

//  ALIGNMENT
// COHESIVENESS
// SEPERATION

func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.position.AddV(viewRadius), b.position.AddV(-viewRadius)
	avgVelocity := Vector2D{0, 0} // for alignment
	avgPosition := Vector2D{0, 0} // for cohesive
	separation := Vector2D{0, 0}  // for seperation
	count := 0.0

	// as you can see in the code that there are lot of memory shared by different thread
	// like screenwidth, screenheight and boids but these data stored in memory are not volatile/changing
	// but boidIdMatrix changes and write operations is also perporfed by the thread on this portion of memory
	// so we will lock this section till the  single thread uses this memory and then we will unlock it for other thread
	// so that other threads can use this after the completion
	lock.RLock()
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			//check for the boid it is there or not and self boid should bot be taken into consideration
			if otherBoidId := boidIdMatrix[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				//get the distance of the boids and check if the dist is less than viewRadius
				if dist := boids[otherBoidId].position.Distance(b.position); dist < viewRadius {
					count++
					// add all the velocity
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
					// add all the positions of the all the boids which are within the boid radius
					avgPosition = avgPosition.Add(boids[otherBoidId].position)
					// distance between 2 boids
					dist_of_2_boids := b.position.Subtract(boids[otherBoidId].position)
					// seperation will be within the viewradius
					separation = separation.Add(dist_of_2_boids).DivideV(dist)
				}
			}
		}
	}
	lock.RUnlock()

	acceleration := Vector2D{b.borderBounce(b.position.x, screenWidth), b.borderBounce(b.position.y, screenHeight)}
	if count > 0 {
		// calculate the average velocity
		avgVelocity = avgVelocity.DivideV(count)
		avgPosition = avgPosition.DivideV(count)
		accelerationAlignment := avgVelocity.Subtract(b.velocity).MultiplyV(adjRate)
		accelerationCohesive := avgPosition.Subtract(b.position).MultiplyV(adjRate)
		accelerationSeparation := separation.MultiplyV(adjRate)
		acceleration = acceleration.Add(accelerationAlignment).Add(accelerationCohesive).Add(accelerationSeparation)
	}
	return acceleration
}

func createBoid(i int) {
	postion := Vector2D{
		x: rand.Float64() * screenWidth,
		y: rand.Float64() * screenHeight,
	}
	velocity := Vector2D{
		x: (rand.Float64() * 2) - 1.0,
		y: (rand.Float64() * 2) - 1.0,
	}

	b := &Boid{
		position: postion,
		velocity: velocity,
		id:       i,
	}

	//save each boid in memory
	boids[i] = b
	// save the boid position
	boidIdMatrix[int(b.position.x)][int(b.position.y)] = b.id

	// start a thread
	go b.start()
}
