package lvl1

type player struct {
	health  float32
	damage  float32
	v_speed float32
	h_speed float32
	left    bool
	right   bool
	jump    bool
	idle    bool
}

var User *player

func Initialise_player() {
	User = &player{
		health:  20,
		damage:  3,
		v_speed: 1,
		h_speed: 5,
	}
}
