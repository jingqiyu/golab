package train

type car interface {
	F()
	B()
	L()
	R()
	Init(int, int, Direction)
	Discover(int, int)
}

type orderList []car

type carImpl struct {
}

func originCarRight() {
	return
}

type position struct {
	x int
	y int
}

type margin struct {
	X int
	Y int
}

func Init(pos position, direction Direction) CarStatus {
	return CarStatus{pos.x, pos.y, direction}
}

type Direction string

type CarStatus struct {
	X         int
	Y         int
	Direction Direction
}

//todo
func TurnRight(nowDirection Direction) CarStatus {

	/*if nowDirection == "N" {
		return "E"
	}

	if nowDirection == "E" {
		return "S"
	}

	if nowDirection == "S" {
		return "W"
	}

	if nowDirection == "W" {
		return "N"
	}
	return ""*/
	return CarStatus{}
}

func TurnLeft(nowDirection Direction) Direction {
	return ""
}
