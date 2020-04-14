package main


import (
	//"github.com/veandco/go-sdl2/sdl"
	"fmt"
)



type gameTable struct{
	left, right int

}

func newGameTable()  gameTable{
	left := -1
	right := -1

	return gameTable{
		left:left,
		right:right,
	}
}

var startPositionWidth float64= float64(width/2-dominoWidth/2)/(float64(0.5))
var startPositionHeight float64= float64(height/2-dominoHeight/2)/float64(0.5)

var x_left = startPositionWidth+dominoWidth
var y_left = startPositionHeight
var x_right = startPositionWidth+2*dominoWidth
var y_right = startPositionHeight

type addingOnTable int

const(
	onStartPosition addingOnTable = -2
	onLeft addingOnTable = 1
	onRight addingOnTable = 2
	onBoth addingOnTable = 0
	onNone addingOnTable = -1
)

//adding domino on left side of the deck
func addDominoOnLeft(table *gameTable, dom *domino) {
	
	tmpLeft := table.left

	if table.left == -1{
		table.left = dom.left
		table.right = dom.right
	}else if dom.left == table.left{
		table.left = dom.right
	}else{
		table.left = dom.left
	}

	
	//TODO

	dom.x = float64(x_left) 
	dom.y = float64(y_left)
	dom.assigned = 0 //on left
	x_left -= dominoWidth

	if (dom.right != tmpLeft){
		dom.rotation = 180
	}
	
} 

//adding domino on right side of the deck
func addDominoOnRight(table *gameTable, dom *domino){
	
	tmpRight := table.right
	
	if dom.left == table.right{
		table.right = dom.right
	}else{
		table.right = dom.left
	}

	//TODO 
	
	dom.x = float64(x_right)
	dom.y = float64(y_right)
	dom.assigned = 0 
	x_right += dominoWidth

	if (dom.left != tmpRight){
		dom.rotation = 180
	}
}

func addDominoOnStart(table *gameTable, dom *domino){
	
	
	if table.left == -1{
		table.left = dom.left
		table.right = dom.right
	}else if dom.left == table.left{
		table.left = dom.right
	}else{
		table.left = dom.left
	}

	
	//TODO

	dom.x = float64(x_left) 
	dom.y = float64(y_left)
	dom.assigned = 0 //on left
	x_left -= dominoWidth

}



func (table gameTable) numOnLeft() int{
	return table.left
}


func (table gameTable) numOnRight() int{
	return table.right
}

//if it's possible to add domino on left function returns 1, if on right then 2,
//if both then 0, if none then -1, if it's start position then -2
func canBeAdded(table *gameTable, dom *domino) addingOnTable{

	if table.left == -1 { //start position
		return onStartPosition
	} else if (dom.left == table.left || dom.left == table.right) && (dom.right == table.left || dom.right == table.right) && (dom.left != dom.right){  //TODO choosing side
		return onBoth
	}else if dom.left == table.left || dom.right == table.left{
		return onLeft
	}else if dom.left == table.right || dom.right == table.right{
		return onRight
	}else{
		return onNone
	}
}


func play(plr *Player, num int, table *gameTable) bool{

	tryAdd := canBeAdded(table, &plr.deck[num])

	if tryAdd == onStartPosition{
		tmpDom := plr.deck[num]
		addDominoOnStart(table, &tmpDom)
		plr.deck[num] = tmpDom //domino changes x and y

		return true
	}else if tryAdd == onLeft{

		tmpDom := plr.deck[num]
		addDominoOnLeft(table, &tmpDom)
		plr.deck[num] = tmpDom

		return true

	}else if tryAdd == onRight{
		tmpDom := plr.deck[num]
		addDominoOnRight(table, &tmpDom)
		plr.deck[num] = tmpDom

		return true

	}else if tryAdd == onBoth{  //TODO
		tmpDom := plr.deck[num]
		addDominoOnRight(table, &tmpDom)
		plr.deck[num] = tmpDom
		return true

	}else{
	/*	var dominoTmp domino
		for _, element := range dominoesMap {
			if element.assigned ==-1{
				dominoTmp = element
				break
			}
		}
		dominoTmp.assigned = plr.deck[num].assigned
		plr.deck = append(plr.deck, dominoTmp)
		if play(plr, len(plr.deck)-1, table){
			return true
		}
	*/	return false
	}
	return false
}

func computerPlay(plr *Player, table *gameTable) bool{

	//for num :=0; num< len(plr.deck); num++{

	for num :=0; num< len(player2.deck); num++{
		if player2.deck[num].assigned==2{
		tryAdd := canBeAdded(table, &player2.deck[num])

		if tryAdd == onStartPosition{
			tmpDom := plr.deck[num]
			addDominoOnStart(table, &tmpDom)
			plr.deck[num] = tmpDom //domino changes x and y
			return true
		}else if tryAdd == onNone{
			continue
		}else if tryAdd == onLeft{
			tmpDom := plr.deck[num]
			addDominoOnLeft(table, &tmpDom)
			plr.deck[num] = tmpDom
			return true

		}else if tryAdd == onRight{
			tmpDom := plr.deck[num]
			addDominoOnRight(table, &tmpDom)
			plr.deck[num] = tmpDom
			return true

		}else if tryAdd == onBoth{  //TODO
			tmpDom := plr.deck[num]
			addDominoOnRight(table, &tmpDom)
			plr.deck[num] = tmpDom
			return true

		}else{
			fmt.Println("Nema dominu :(")
			return true
		}
	}
	}
	return false
}


func isWon(plr *Player) bool {
	var forCheck int = 0
	if plr==&player1{
		forCheck = 1
	}
	if plr==&player2{
		forCheck=2
	}

	for _, dom := range plr.deck {
		if dom.assigned == forCheck{
			return false
		}
	}
	return true
}
