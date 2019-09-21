package robot

// Converts Dir, which is an int under the hood, to a "direction name" string
func (d Dir) String() string {
	return "NESW"[d : d+1]
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y += 1
		break
	case E:
		Step1Robot.X += 1
		break
	case S:
		Step1Robot.Y -= 1
		break
	case W:
		Step1Robot.X -= 1
		break
	}
}

func Left() {
	Step1Robot.Dir = (Step1Robot.Dir + 3) % 4
}

func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

func (robot *Step2Robot) Advance(extent Rect) {
	switch robot.Dir {
	case N:
		if robot.Pos.Northing < extent.Max.Northing {
			robot.Pos.Northing += 1
		}
		break
	case E:
		if robot.Pos.Easting < extent.Max.Easting {
			robot.Pos.Easting += 1
		}
		break
	case S:
		if robot.Pos.Northing > extent.Min.Northing {
			robot.Pos.Northing -= 1
		}
		break
	case W:
		if robot.Pos.Easting > extent.Min.Easting {
			robot.Pos.Easting -= 1
		}
		break
	}
}

func (robot *Step2Robot) Left() {
	robot.Dir = (robot.Dir + 3) % 4
}

func (robot *Step2Robot) Right() {
	robot.Dir = (robot.Dir + 1) % 4
}

func StartRobot(cmd chan Command, act chan Action) {
	for command := range cmd {
		act <- Action(command)
	}
	close(act)
}

func Room(extent Rect, robot Step2Robot, act chan Action, report chan Step2Robot) {
	for action := range act {
		switch action {
		case 'A':
			robot.Advance(extent)
			break
		case 'L':
			robot.Left()
			break
		case 'R':
			robot.Right()
			break
		}
	}
	report <- robot
	close(report)
}

type Action3 struct {
	name   string
	action byte
}

const beep = 7

func (robot *Step3Robot) Advance(extent Rect, log chan string, posMap map[Pos]bool) {
	var newPosition = robot.Pos

	switch robot.Dir {
	case N:
		newPosition.Northing += 1
		if newPosition.Northing > extent.Max.Northing {
			log <- "Robot hit wall"
			return
		}
	case E:
		newPosition.Easting += 1
		if newPosition.Easting > extent.Max.Easting {
			log <- "Robot hit wall"
			return
		}
	case S:
		newPosition.Northing -= 1
		if newPosition.Northing < extent.Min.Northing {
			log <- "Robot hit wall"
			return
		}
	case W:
		newPosition.Easting -= 1
		if newPosition.Easting < extent.Min.Easting {
			log <- "Robot hit wall"
			return
		}
	}
	if posMap[newPosition] == true {
		log <- "robot bumped robot"
		return
	}
	// Position change is valid.  Remove the robot's starting position from our posMap
	// and add the new position.
	delete(posMap, robot.Pos)
	robot.Pos = newPosition
	posMap[newPosition] = true
}

func (robot *Step3Robot) Left() {
	robot.Dir = (robot.Dir + 3) % 4
}

func (robot *Step3Robot) Right() {
	robot.Dir = (robot.Dir + 1) % 4
}

func in(room Rect, pos Pos) bool {
	return room.Min.Easting <= pos.Easting && room.Max.Easting >= pos.Easting && room.Min.Northing <= pos.Northing && room.Max.Northing >= pos.Northing
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {
	defer func() { report <- robots }()
	posMap := getRobotPositions(robots)

	if !validateRobotPositions(posMap, robots, log, extent) {
		return
	}

	robotMap := buildNameMap(robots, log)
	if len(robotMap) != len(robots) {
		return
	}

	done := 0
	for robotAction := range action {
		command := robotAction.action
		name := robotAction.name
		if robot, exists := robotMap[name]; exists {
			switch command {
			case 'A':
				robot.Advance(extent, log, posMap)
				break
			case 'L':
				robot.Left()
				break
			case 'R':
				robot.Right()
				break
			case beep:
				if done++; done == len(robots) {
					for i, robot := range robots {
						if robot.Name == robotAction.name {
							robots[i] = *(robotMap[robot.Name])
							break
						}
					}
					return
				}
			default:
				log <- "undefined command in script"
				return
			}
		} else {
			log <- "unknown robot name in script"
			return
		}
	}
}

func buildNameMap(robots []Step3Robot, log chan string) map[string]*Step3Robot {
	robotMap := map[string]*Step3Robot{}
	for _, robot := range robots {
		if robot.Name == "" {
			log <- "robot missing name"
		}
		if _, exists := robotMap[robot.Name]; exists {
			log <- "robot name duplicated"
		} else {
			robotMap[robot.Name] = &robot
		}
	}
	return robotMap
}

func validateRobotPositions(posMap map[Pos]bool, robots []Step3Robot, log chan string, extent Rect) bool {
	if len(posMap) < len(robots) {
		log <- "Two or more robots have the same starting location"
		return false
	}
	for pos, _ := range posMap {
		if !in(extent, pos) {
			log <- "robot placement outside room boundaries"
			return false
		}
	}
	return true
}

func getRobotPositions(robots []Step3Robot) map[Pos]bool {
	posMap := map[Pos]bool{}
	for _, robot := range robots {
		posMap[robot.Step2Robot.Pos] = true
	}
	return posMap
}

func StartRobot3(name, script string, act chan Action3, log chan string) {
	for i := 0; i < len(script); i++ {
		act <- Action3{name, script[i]}
	}
	act <- Action3{name, beep}
}
