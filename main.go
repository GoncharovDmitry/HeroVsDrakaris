package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Players interface {
	setName(name string)
	getDamage(damage int)
}

type mainMenuText struct {
	welcome string
	start   string
	exit    string
}

type Hero struct {
	name   string
	health int
	weapon weapon
	damage float32
}

type Dragon struct {
	name   string
	health int
	weapon weapon
	damage float32
}

type weapon struct {
	name      string
	minDamage int
	maxDamage int
}

func (hero *Hero) getDamage(damage int) {
	hero.health -= damage
}

func (hero *Hero) setName(name string) {
	hero.name = name
}

func (dragon *Dragon) getDamage(damage int) {
	dragon.health -= damage
}

func (dragon *Dragon) setName(name string) {
	dragon.name = name
}

func setName(players Players, name string) {
	players.setName(name)
}

func hitEnemy(enemy Players, damage int) {
	enemy.getDamage(damage)
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func gameStart() {
	fmt.Println(menu.welcome)
	switch mainMenuSelected() {
	case menu.start:
		initPlayers()
	case menu.exit:
		fmt.Println("Пока")
		os.Exit(0)
	}
}
func mainMenuSelected() string {
	switch userInput(">>") {
	case "1":
		return menu.start
	case "2":
		return menu.exit
	}
	fmt.Println("Попробуй ещё раз, я в тебя верю")
	return mainMenuSelected()
}

func userInput(welcomeLabel string) string {
	fmt.Print(welcomeLabel)
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	input := scan.Text()
	return input
}

func initPlayers() {
	hero = Hero{
		name:   userInput("Введите имя игорка:"),
		weapon: choiceWeapon(initWeapon()),
		health: 100,
		damage: 1.0,
	}
	if len(hero.name) == 0 {
		fmt.Println("Тогда будешь Васей")
		setName(&hero, "Вася")
	}
	dragon = Dragon{
		name:   userInput("Введите имя дракона:"),
		weapon: choiceWeapon(initWeapon()),
		health: 100,
		damage: 1.0,
	}
	if len(dragon.name) == 0 {
		fmt.Println("Тогда будешь биться с Drakaris")
		setName(&dragon, "Drakaris")
	}
	fight()
}

func initWeapon() []weapon {
	weaponList := []weapon{
		{name: "Меч", minDamage: 10, maxDamage: 20},
		{name: "Лук", minDamage: 7, maxDamage: 27},
		{name: "Голые руки", minDamage: 5, maxDamage: 10},
		{name: "Эскалибур", minDamage: 15, maxDamage: 30},
		{name: "Перчатка Таноса", minDamage: 30, maxDamage: 50},
		{name: "Меч джедая", minDamage: 20, maxDamage: 35},
	}
	return weaponList
}

func randChoiceWeapon(weaponList []weapon) weapon {
	randIndexOfListWeapon := randInt(0, len(weaponList)-1)
	return weaponList[randIndexOfListWeapon]
}

func choiceWeapon(weaponList []weapon) weapon {
	fmt.Println("Выберите оружние")
	fmt.Println(weaponList)
	return weaponList[0]
}

func fight() {
	fmt.Println("Битва началась")
	var round = 0
	for dragon.health > 0 && hero.health > 0 {
		fmt.Println("Раунд: ", round)
		fmt.Printf("%v здоровье: %v\t%v здоровье: %v\n", hero.name, hero.health, dragon.name, dragon.health)
		hitEnemy(&hero, randInt(0, 10))
		hitEnemy(&dragon, randInt(0, 10))
		round++
	}
	fmt.Println("Конец, выиграл ")
}

var hero = Hero{}
var dragon = Dragon{}
var menu = mainMenuText{
	welcome: "Добро пожаловать в игру" +
		"\nВ этой игре вам необходимо победить дракона" +
		"\nВыберите действие" +
		"\n1 - начать игру 2 - выйти",
	start: "start",
	exit:  "exit",
}

func main() {
	gameStart()
}
