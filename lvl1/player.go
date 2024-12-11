package lvl1

type player struct {
	health  float32
	damage  float32
	v_speed float32
	h_speed float32
}

var User *player

func Initialise_player() {
	User = &player{
		health:  20,
		damage:  3,
		v_speed: 3,
		h_speed: 3.5,
	}
}
