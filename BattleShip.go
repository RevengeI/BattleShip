package main

import (
	"errors"
	"fmt"
	"math"
)

var (
	Hello             = "Добро пожалвать в игру 'Морской бой'!\nВ данной игре Вам предстоит заполнить поле кораблями, а затем Вашему сопернику необходимо попасть и уничтожить все корабли\n"
	StartChoice       = "Выберите дейтвие:\n1: Начать игру\n2: Запуск теста\n3: Закончить игру\n"
	ErrInput          = "Ошибка ввода! Выберите цифру от 1 до %d\n"
	SelectQuantity    = "Выбирите количество"
	TnxFGame          = "Спасиб за игру! Возвращайтесь снова."
	ShipPlaced        = "Корабли размещены!"
	StartTheGame      = "Выберите действие:\n1: Посмотреть карту\n2: Передать управление игроку.\n3: Закончить игру\n"
	ShipCountText     = "На столе расположено %d катер%s, %d малы%s корабл%s, %d крейсер%s и %d авианос%s. Каждый из них занимает 1, 2, 3 и 4 ячейки соответственно.\n"
	Task              = "Вам предстоит одолеть всю вражескую фотилию! Для того, чтобы это сделать выбирайте нужную ячейку. Но будте бдительны, у Вас всего %d выстрелов!\n"
	GamerTutorial     = "Это - поле боя. Для нанесения удара Вам требуется указать ячейку для удара через пробел (пример: A 1).\n"
	StartTxt          = "Начали!\n"
	GamerChoice       = "1: Вывести текущую карту\n2: Сделать выстрел\n3: Количсество оставшихся выстрелов\n4: Закончить игру\n"
	SentenceChoice    = "Выберите один из вариантов: "
	Arrangement       = "\nТеперь обратите внимание на поле, Вам следует расставить корабли так, чтобы они не соприкасались друг с другом: \n"
	RemainderShots    = "У Вас осталось %d выстрелов!\n"
	FireTxt           = "Выберите поле для выстрела: "
	sizeMap           = 121
	NameLineMap       = string(" ABCDEFGHIJ")
	Space             = "  "
	boatCount         = 0
	minShipCount      = 0
	cruiserCount      = 0
	AirCarCount       = 0
	ErrCoordinatesLen = errors.New("Некорекнтые координаты! Символы должны вводиться через пробел.")
	ErrCoordinatesNum = errors.New("Некорекнтые координаты! Цифра должна быть от 1 до 10.")
	ErrCoordinatesLet = errors.New("Некорекнтые координаты! Бувы должны использоваться только в верхнем регистре, английского алфавита от A до J.")
	ErrLocation       = errors.New("Корабли расположены слишком близко!")
	ErrCrookedShip    = errors.New("Корабль должен находиться на одной прямой!")
	ErrShipLength     = errors.New("Недопустимая длина корабля!")
	ErrEngage         = errors.New("Позиция занята другим кораблем!")
	ErrNotEnoughtLen  = errors.New("Длина корабля слишком мала!")
	ErrUnknown        = errors.New("Неизвестная ошибка")
	Creator           = "Создатель"
	Gamer             = "Игрок"
	ShipsCount        = 0
	ShotCount         = 50
	GM                = map[bool]bool{true: false}
)

type intCoordinates [][11]int

func main() {
	fmt.Println(Hello)
	start()
}

func start() {
	fmt.Print(StartChoice, SentenceChoice)
	var res int
	fmt.Scanf("%d/n", &res)
	if res < 1 || res > 3 {
		fmt.Printf(ErrInput, 3)
		start()
	}
	switch res {
	case 1:
		startGame()
		main()
		return
	case 2:
		startTest()
		main()
		return
	case 3:
		fmt.Println(TnxFGame)
		return
	}
}

func startGame() {
	fmt.Printf("%s катеров: ", SelectQuantity)
	shipCount(&boatCount, 4)
	fmt.Printf("%s малых кораблей: ", SelectQuantity)
	shipCount(&minShipCount, 3)
	fmt.Printf("%s крейсеров: ", SelectQuantity)
	shipCount(&cruiserCount, 2)
	fmt.Printf("%s авианосцев: ", SelectQuantity)
	shipCount(&AirCarCount, 1)

	if boatCount == 0 && minShipCount == 0 && cruiserCount == 0 && AirCarCount == 0 {
		return
	}

	gameMap := createMap(sizeMap)

	fmt.Print(Arrangement)
	printMap(gameMap, Creator)
	replaseShip("boat", boatCount, &gameMap)
	fmt.Println()
	printMap(gameMap, Creator)
	replaseShip("minShip", minShipCount, &gameMap)
	fmt.Println()
	printMap(gameMap, Creator)
	replaseShip("cruiser", cruiserCount, &gameMap)
	fmt.Println()
	printMap(gameMap, Creator)
	replaseShip("AirCar", AirCarCount, &gameMap)
	fmt.Println()
	printMap(gameMap, Creator)
	fmt.Println(ShipPlaced)
	finlChoise(gameMap)
	fmt.Print(TnxFGame)
}

func startTest() int {
	return 1
}

func shipCount(SCount *int, maxShip int) {

	fmt.Scanf("%d/n", SCount)

	if *SCount == -1 {
		return
	}

	ShipsCount = ShipsCount + *SCount

	if *SCount < 0 || *SCount > maxShip {
		fmt.Printf("Выбирите число от 0 до %d!\n", maxShip)
		shipCount(SCount, maxShip)
	}
	return
}

func createMap(sizeMap int) map[int]map[bool]bool {
	gameMap := make(map[int]map[bool]bool, sizeMap)
	fmt.Println(gameMap)
	for i := 0; i < sizeMap; i++ {
		gameMap[i] = make(map[bool]bool, 1)
		gameMap[i][false] = false
		fmt.Println(gameMap[i][false])
	}
	return gameMap
}

func printMap(gameMap map[int]map[bool]bool, user string) {
	for i := 0; i <= sizeMap-1; i++ {
		if i == 0 {
			fmt.Print(Space)
			continue
		}
		if i/11 == 0 {
			fmt.Print(" ", string(NameLineMap[i]), " ")
		}
		if i%11 == 0 {
			if i/11 == 10 {
				fmt.Print(i / 11)
				continue
			}
			fmt.Print(i/11, " ")
			continue
		}
		if user == Creator {
			for isShip := range gameMap[i] {
				if isShip == true {
					fmt.Print(" + ")
				} else if i/11 != 0 {
					fmt.Print(" - ")
				}
			}
		} else {
			if i/11 != 0 {
				for isShip, isShoot := range gameMap[i] {
					if isShoot == true {
						if isShip == true {
							fmt.Print(" + ")
						} else {
							fmt.Print(" - ")
						}
					} else {
						fmt.Print("   ")
					}
				}
			}
		}
		if i%11 == 10 {
			fmt.Println()
			continue
		}
	}
	fmt.Println()
}

func replaseShip(shipName string, shipCount int, gameMap *map[int]map[bool]bool) {
	if shipCount == 0 {
		return
	}

	var (
		name       = ""
		endCell    = "и"
		numCell    = "первый и последний"
		example    = ""
		shipLength = 0
	)

	intCdt := intCoordinates([][11]int{})

	switch shipName {
	case "boat":
		name = "Катер"
		endCell = "у"
		numCell = "её"
		example = "А 1"
		shipLength = 1
	case "minShip":
		name = "Малый корабль"
		example = "А 1 А 2"
		shipLength = 2
	case "cruiser":
		name = "Крейсер"
		example = "А 1 А 3"
		shipLength = 3
	case "AirCar":
		name = "Авианосец"
		example = "А 1 А 4"
		shipLength = 4
	}

	fmt.Printf("%s занимает %d ячейк%s, введите %s номер, чтобы поставить %s (пример: %s):\n", name, shipLength, endCell, numCell, name, example)

	for i := 0; i < shipCount; i++ {
		fmt.Printf("%s #%d: ", name, i+1)
		var (
			letter1 = ""
			letter2 = ""
			num1    = -1
			num2    = -1
		)

		fmt.Scan(&letter1, &num1)
		if shipName != "boat" {
			fmt.Scan(&letter2, &num2)
		}

		if num1 == -1 {
			return
		}

		err := checkCoordinates(letter1, num1)
		if err != nil {
			fmt.Println(err)
			i--
			continue
		}

		if shipName != "boat" {
			err2 := checkCoordinates(letter2, num2)
			if err2 != nil {
				fmt.Println(err2)
				i--
				continue
			}
		}

		intCdt = convertCdt(letter1, letter2, num1, num2)

		err = checkCrookShip(intCdt, shipLength)
		if err != nil {
			fmt.Println(err)
			i--
			continue
		}

		*gameMap, err = addShip(intCdt, *gameMap, shipLength)

		if err != nil {
			fmt.Println(err)
			i--
		}
		//printMap(*gameMap, Creator)
	}
}

func checkCoordinates(letter string, num int) error {

	if len(letter) > 1 {
		return ErrCoordinatesLen
	}
	checkLett := false
	for _, dicLett := range NameLineMap {
		if letter == string(dicLett) {
			checkLett = true
			break
		}
	}
	if checkLett == false {
		return ErrCoordinatesLet
	}
	if num < 1 || num > 10 {
		return ErrCoordinatesNum
	}
	return nil
}

func checkCrookShip(cdt [][11]int, shipLength int) error {
	if shipLength == 1 {
		return nil
	}
	var lets, nums [2]int
	sideCheck := 0
	highOrLengh := false

	//fmt.Println(cdt)

	for i := 0; i < 2; i++ {
		for _, let := range cdt {
			for j, lett := range let {
				if lett != 0 {
					lets[i] = lett
					nums[i] = j
					//fmt.Println(cdt, let, lett, j)
					i++
				}
			}
		}
	}

	//fmt.Println(lets, nums)

	if nums[0] == nums[1] {
		sideCheck++
	}
	if lets[0] == lets[1] {
		sideCheck++
		highOrLengh = true
	}

	switch sideCheck {
	case 0:
		return ErrCrookedShip
	case 2:
		return ErrNotEnoughtLen
	}

	if highOrLengh == false {
		//fmt.Println(math.Abs(float64(lets[0] - lets[1])))
		//fmt.Println(float64(shipLength) - 1)
		if math.Abs(float64(lets[0]-lets[1])) != float64(shipLength)-1 {
			return ErrShipLength
		}
	} else {
		//fmt.Println(math.Abs(float64(nums[0] - nums[1])))
		//fmt.Println(float64(shipLength) * math.Sqrt(float64(sizeMap)))
		if (math.Abs(float64(nums[0]-nums[1]))+1)*math.Sqrt(float64(sizeMap)) != float64(shipLength)*math.Sqrt(float64(sizeMap)) {
			return ErrShipLength
		}
	}

	return nil
}

func convertCdt(let1, let2 string, num1, num2 int) [][11]int {
	intCdt := make([][11]int, 2)
	for i, dicLett := range NameLineMap {
		if let1 == string(dicLett) {
			intCdt[0][i] = num1
		}
		if let2 == string(dicLett) {
			intCdt[1][i] = num2
		}
	}
	return intCdt
}

func addShip(cdt [][11]int, gameMap map[int]map[bool]bool, shipLength int) (map[int]map[bool]bool, error) {
	var (
		convCdt, checkPos int
		lets, nums        [2]int
	)
	fnlShip := make([]int, 0)
	for i := 0; i < 2; i++ {
		for let, num := range cdt[i] {
			if num == 0 {
				continue
			}
			convCdt = num*int(math.Sqrt(float64(sizeMap))) + let
			for i := 0; i < 9; i++ {
				checkCdt := changeCdt(convCdt, i)
				for isShip := range gameMap[checkCdt] {
					if isShip == true {
						return gameMap, ErrEngage
					}
					//fmt.Println(let, num, convCdt)
				}
			}
			fnlShip = append(fnlShip, convCdt)
			checkPos++
			lets[i] = let
			nums[i] = num
		}
	}

	if shipLength > 2 {
		AdCdt := make([][11]int, shipLength-2)
		for i := 0; i < shipLength-2; i++ {
			if lets[0] == lets[1] {
				if nums[0] < nums[1] {
					AdCdt[i][lets[0]] = cdt[0][lets[0]] + 1 + i
				} else {
					AdCdt[i][lets[0]] = cdt[0][lets[0]] - 1 - i
				}
			} else {
				if lets[0] < lets[1] {
					AdCdt[i][lets[0]+1+i] = cdt[0][lets[0]]
				} else {
					AdCdt[i][lets[0]-1-i] = cdt[0][lets[0]]
				}
			}
			//fmt.Println(lets, nums, cdt[i])
		}
		//fmt.Println(AdCdt)

		for _, AdsCdt := range AdCdt {
			for let, num := range AdsCdt {
				if num == 0 {
					continue
				}
				convCdt = num*int(math.Sqrt(float64(sizeMap))) + let
				for i := 0; i < 9; i++ {
					checkCdt := changeCdt(convCdt, i)
					for isShip := range gameMap[checkCdt] {
						if isShip == true {
							return gameMap, ErrEngage
						}
						//fmt.Println(let, num, convCdt)
					}
				}
				fnlShip = append(fnlShip, convCdt)
				checkPos++
				//fmt.Println(convCdt, num, let)
			}
		}

	}

	//fmt.Println(fnlShip)

	if checkPos == shipLength {
		for _, fnlCdt := range fnlShip {
			gameMap[fnlCdt] = GM
		}
	} else {
		return gameMap, ErrUnknown
	}
	return gameMap, nil
}

func startGamePlayer(gameMap map[int]map[bool]bool) {
	var boatEnd, minEnd, minShipEnd, cruiserEnd, airCarEnd string
	boatEnd, minEnd, minShipEnd, cruiserEnd, airCarEnd = shipTxtEnd()

	fmt.Printf(ShipCountText, boatCount, boatEnd, minShipCount, minEnd, minShipEnd, cruiserCount, cruiserEnd, AirCarCount, airCarEnd)
	fmt.Printf(Task, ShotCount)
	fmt.Scan()
	printMap(gameMap, Gamer)
	fmt.Print(GamerTutorial)
	fmt.Scan()
	fmt.Print(StartTxt)
	startFire(gameMap)
}

func startFire(gameMap map[int]map[bool]bool) {
	fmt.Print(GamerChoice, SentenceChoice)
	var input int
	fmt.Scan(&input)

	if input == -1 {
		start()
	}

	if input < 1 || input > 4 {
		fmt.Printf(ErrInput, 4)
		startFire(gameMap)
	}

	switch input {
	case 1:
		printMap(gameMap, Gamer)
		startFire(gameMap)
	case 2:
		fire(&gameMap)
		startFire(gameMap)
	case 3:
		fmt.Printf(RemainderShots, ShotCount)
		startFire(gameMap)
	case 4:
		return
	}
}

func fire(gameMap *map[int]map[bool]bool) {
	fmt.Print(FireTxt)
	var (
		let string
		num int
	)
	fmt.Scan(&let, &num)

	err := checkCoordinates(let, num)
	if err != nil {
		fmt.Println(err)
		fire(gameMap)
	}

	intCdt := intCoordinates([][11]int{})
	intCdt = convertCdt(let, "-1", num, -1)

	msg := checkFire(intCdt, *gameMap)
	fmt.Println(msg)
}

func checkFire(cdt [][11]int, gameMap map[int]map[bool]bool) string {
	for let, num := range cdt[0] {
		convCdt := num*int(math.Sqrt(float64(sizeMap))) + let
		for isShip := range gameMap[convCdt] {
			if isShip == true {
				fmt.Println(convCdt)
				gameMap[convCdt][true] = true
				return "Попадание!"
			} else {
				gameMap[convCdt][false] = true
				return "Промах!"
			}
		}
	}
	return "Error"
}

func changeCdt(j, convCdt int) int {
	switch j {
	case 0:
		return convCdt
	case 1:
		return convCdt + 1
	case 2:
		return convCdt - 1
	case 3:
		return convCdt + 11
	case 4:
		return convCdt - 11
	case 5:
		return convCdt + 12
	case 6:
		return convCdt - 12
	case 7:
		return convCdt + 10
	case 8:
		return convCdt - 10
	}
	return 0
}

func finlChoise(gameMap map[int]map[bool]bool) {
	fmt.Print(StartTheGame, SentenceChoice)
	var input int
	fmt.Scanf("%d\n", &input)
	switch input {
	case -1:
		return
	case 1:
		printMap(gameMap, Creator)
	case 2:
		startGamePlayer(gameMap)
	case 3:
		return
	default:
		fmt.Printf(ErrInput, 3)
		finlChoise(gameMap)
	}
}

func shipTxtEnd() (string, string, string, string, string) {
	var boatEnd, minEnd, minShipEnd, cruiserEnd, airCarEnd string
	switch boatCount {
	case 0:
		boatEnd = "ов"
	case 2, 3, 4:
		boatEnd = "а"
	}
	switch minShipCount {
	case 0:
		minEnd = "х"
		minShipEnd = "ей"
	case 1:
		minEnd = "ый"
		minShipEnd = "ь"
	case 2, 3:
		minEnd = "ыx"
		minShipEnd = "я"
	}
	switch cruiserCount {
	case 0:
		cruiserEnd = "ов"
	case 2:
		cruiserEnd = "а"
	}
	switch AirCarCount {
	case 0:
		airCarEnd = "цев"
	case 1:
		airCarEnd = "ец"
	}
	return boatEnd, minEnd, minShipEnd, cruiserEnd, airCarEnd
}
