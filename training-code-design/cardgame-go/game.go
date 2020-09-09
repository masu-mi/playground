package cardgame

type Rounds struct {
	next    int
	Rounds  []Game
	Players map[int]*Player
}

func NewRounds(game Game, names []string) *Rounds {
	players := NewPlayers(names)
	game.RegisterPlayers(players...)
	r := &Rounds{
		Rounds:  []Game{game},
		Players: map[int]*Player{},
	}
	for _, p := range players {
		r.Players[p.ID] = p
	}
	return r
}

func NewPlayers(names []string) []*Player {
	var ps []*Player
	for id, name := range names {
		ps = append(ps, &Player{
			ID:   id,
			Name: name,
		})
	}
	return ps
}

func (r *Rounds) PlayAllRound() {
	for r.next < len(r.Rounds) {
		r.PlayRound(r.next)
		r.next++
	}
}

func (r *Rounds) PlayRound(round int) {
	result := r.Rounds[round].Play()
	r.Update(result)
}

func (r *Rounds) Update(result Result) {
	for id := range result.Scores {
		r.Players[id].Status.Score += result.Scores[id]
	}
}

type Game interface {
	RegisterPlayers(playerNames ...*Player)
	Play() Result
}

type Result struct {
	Scores map[int]int
}

type Player struct {
	ID   int
	Name string

	Status
}

type Status struct {
	Score int
}
