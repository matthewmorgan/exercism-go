package robot

import "fmt"

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

//func StartRobot3(name, script string, action chan Action3, log chan string) {
//	if name == "" {
//		log <- "Robot is missing name"
//		return
//	}
//	for _, command := range []byte(script){
//		action <- Action3{name, command }
//	}
//}

// ======= Step 3

type Action3 struct {
	name   string
	action byte
}

const beep = 7 // robots beep to communicate that they are done


//func (robot *Step3Robot) Advance(extent Rect, log chan string) {
//	// TODO first compute position, then check for presence of other robot
//	//  or wall.  If either, don't advance and log event
//
//	switch robot.Dir {
//	case N:
//		if robot.Pos.Northing < extent.Max.Northing {
//			robot.Pos.Northing += 1
//		}
//		break
//	case E:
//		if robot.Pos.Easting < extent.Max.Easting {
//			robot.Pos.Easting += 1
//		}
//		break
//	case S:
//		if robot.Pos.Northing > extent.Min.Northing {
//			robot.Pos.Northing -= 1
//		}
//		break
//	case W:
//		if robot.Pos.Easting > extent.Min.Easting {
//			robot.Pos.Easting -= 1
//		}
//		break
//	}
//}
//
//func (robot *Step3Robot) Left() {
//	robot.Dir = (robot.Dir + 3) % 4
//}
//
//func (robot *Step3Robot) Right() {
//	robot.Dir = (robot.Dir + 1) % 4
//}

func in(room Rect, pos Pos) bool {
	return room.Min.Easting <= pos.Easting && room.Max.Easting >= pos.Easting && room.Min.Northing <= pos.Northing && room.Max.Northing >= pos.Northing
}

//func Room3(extent Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {
//	defer func() { report <- robots }()
//	posMap := getRobotPositions(robots)
//	println("position map computed", len(posMap))
//
//	validateRobotPositions(posMap, robots, log, extent)
//	println("robot positions validated")
//
//	robotMap := buildNameMap(robots, log)
//	println("name map completed", len(robotMap))
//
//	println("actions received", len(action))
//
//	for robotAction := range action {
//		command := robotAction.action
//		println("command", command)
//		name := robotAction.name
//		if _, exists := robotMap[name]; exists {
//			switch command {
//			case 'A':
//				robot.Advance(extent, log)
//				break
//			case 'L':
//				robot.Left()
//				break
//			case 'R':
//				robot.Right()
//				break
//			default:
//				log <- "undefined command in script"
//			}
//		} else {
//			log <- "unknown robot name in script"
//		}
//
//	}
//
//
//	println("room function completed")
//	for a := range action {
//		println("action received [", a.action, "]")
//	}
//}
//
//func buildNameMap(robots []Step3Robot, log chan string) map[string]*Step3Robot {
//	robotMap := map[string]*Step3Robot{}
//	for _, robot := range robots {
//		if robot.Name == "" {
//			log <- "robot missing name"
//		}
//		if _, exists := robotMap[robot.Name]; exists {
//			log <- "robot name duplicated"
//		} else {
//			robotMap[robot.Name] = &robot
//		}
//	}
//	return robotMap
//}
//
//func validateRobotPositions(posMap map[Pos]bool, robots []Step3Robot, log chan string, extent Rect) {
//	if len(posMap) < len(robots) {
//		log <- "Two or more robots have the same starting location"
//	}
//	for pos, _ := range posMap {
//		if !in(extent, pos) {
//			log <- "robot placement outside room boundaries"
//		}
//	}
//}
//
//func getRobotPositions(robots []Step3Robot) map[Pos]bool {
//	posMap := map[Pos]bool{}
//	for _, robot := range robots {
//		posMap[robot.Step2Robot.Pos] = true
//	}
//	return posMap
//}

func StartRobot3(name, script string, act chan Action3, log chan string) {
	for i := 0; i < len(script); i++ {
		act <- Action3{name, script[i]}
	}
	act <- Action3{name, beep}
}

func Room3(extent Rect, robots []Step3Robot, actionChannel chan Action3, rep chan []Step3Robot, log chan string) {
	// The function has multiple returns.  No matter what, rep <- is how we
	// communicate to the test program that the room is terminating.
	defer func() { rep <- robots }()
	nameMap := map[string]*Step3Robot{}
	positionMap := map[Pos]*Step3Robot{}
	for _, robot := range robots {
		println(robot.Name)
		if robot.Name == "" {
			log <- "Unnamed robot"
			return
		}
		if _, ok := nameMap[robot.Name]; ok {
			log <- "Duplicate name"
			return
		}
		nameMap[robot.Name] = &robot

		if !in(extent, robot.Step2Robot.Pos) {
			log <- "Robot placed outside room"
			return
		}
		if _, ok := positionMap[robot.Step2Robot.Pos]; ok {
			log <- "Position occupied"
			return
		}
		positionMap[robot.Step2Robot.Pos] = &robot
	}
	done := 0
	for a := range actionChannel {
		println("processing action", a.action)
		var robotPointer *Step3Robot
		robotPointer, ok := nameMap[a.name]
		if !ok {
			log <- "Action by unknown robot"
			return
		}
		var actingRobot *Step2Robot
		actingRobot = &robotPointer.Step2Robot
		switch a.action {
		case 'R':
			actingRobot.Dir = (actingRobot.Dir + 1) % 4
		case 'L':
			actingRobot.Dir = (actingRobot.Dir + 3) % 4
		case 'A':
			newPosition := actingRobot.Pos
			if actingRobot.Dir&1 == 1 {
				newPosition.Easting += 1 - RU(actingRobot.Dir&2)
			} else {
				newPosition.Northing += 1 - RU(actingRobot.Dir&2)
			}
			if !in(extent, newPosition) {
				log <- a.name + " bumps wall!"
				continue
			}
			if r, occupied := positionMap[newPosition]; occupied {
				log <- fmt.Sprintf("%s bumps %s!", a.name, r.Name)
				continue
			}
			delete(positionMap, actingRobot.Pos)
			positionMap[newPosition] = nameMap[a.name]
			actingRobot.Pos = newPosition
		case beep:
			if done++; done == len(robots) {
				return
			}
		default:
			log <- "Undefined command"
			return
		}
	}
}

//func Room3(extent Rect, robots []Step3Robot, act chan Action3, rep chan []Step3Robot, log chan string) {
//	// The function has multiple returns.  No matter what, rep <- is how we
//	// communicate to the test program that the room is terminating.
//	defer func() { rep <- robots }()
//	nameIndex := map[string]int{}  // name index back into robots slice
//	positionIndex := map[Pos]int{} // position index back into robots slice
//	for x, robot := range robots {
//		if robot.Name == "" {
//			log <- "Unnamed robot"
//			return
//		}
//		if _, ok := nameIndex[robot.Name]; ok {
//			log <- "Duplicate name"
//			return
//		}
//		nameIndex[robot.Name] = x
//
//		if !in(extent, robot.Step2Robot.Pos) {
//			log <- "Robot placed outside room"
//			return
//		}
//		if _, ok := positionIndex[robot.Step2Robot.Pos]; ok {
//			log <- "Position occupied"
//			return
//		}
//		positionIndex[robot.Step2Robot.Pos] = x
//	}
//	done := 0
//	for action := range act {
//		indexInRobotSlice, ok := nameIndex[action.name]
//		if !ok {
//			log <- "Action by unknown robot"
//			return
//		}
//		actingRobot := &robots[indexInRobotSlice].Step2Robot
//		switch action.action {
//		case 'R':
//			actingRobot.Dir = (actingRobot.Dir + 1) % 4
//		case 'L':
//			actingRobot.Dir = (actingRobot.Dir + 3) % 4
//		case 'A':
//			newPosition := actingRobot.Pos
//			if actingRobot.Dir&1 == 1 {
//				newPosition.Easting += 1 - RU(actingRobot.Dir&2)
//			} else {
//				newPosition.Northing += 1 - RU(actingRobot.Dir&2)
//			}
//			if !in(extent, newPosition) {
//				log <- action.name + " bumps wall!"
//				continue
//			}
//			if y, occupied := positionIndex[newPosition]; occupied {
//				log <- fmt.Sprintf("%s bumps %s!", action.name, robots[y].Name)
//				continue
//			}
//			delete(positionIndex, actingRobot.Pos)
//			positionIndex[newPosition] = indexInRobotSlice
//			actingRobot.Pos = newPosition
//		case beep:
//			if done++; done == len(robots) {
//				return
//			}
//		default:
//			log <- "Undefined command"
//			return
//		}
//	}
//}