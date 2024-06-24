package main

import "fmt"

type Position struct {
	x float64
	y float64
}

func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
}

func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}

// embeding or inheriting Position
// can use all the methods available from Position
type Player struct {
	*Position
}

//also inherits Person
type SuperPosition struct {
  Position
}

func (s *SuperPosition) SuperMove(x,y float64) {
  s.x += x*x 
  s.y += y*y 
}


type Enemy struct {
  *SuperPosition
}

// person constuctor
func NewPlayer() *Player{
  return &Player{
    Position: &Position{},
  }
}

// enemy constuctor
func NewEnemy() *Enemy{
  return &Enemy{
    SuperPosition: &SuperPosition{},
  }
}

func main() {
  boss:= NewEnemy()
  player := NewPlayer()

  player.Move(10,5)
  fmt.Println("move player", player.Position)

  boss.Move(10,5)
  fmt.Println("move boss", boss.Position)

  boss.SuperMove(10,5)
  fmt.Println("super move boss", boss.Position)
}
