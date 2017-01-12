package main

import "fmt"
//import "math"
import "math/rand"
import "time"

const numberofboards = 101
const boardsize = 101
const numberofparts = 101

var pos int= 1;
var turnsofdice int= 0;
var turnsofgame int= 0
var turnsofdiceforagame int= 0;
var states[101]int;
var numberofdiceturns int= 100000;
var turnsasmoney float64 = 100;
var turnsasmoneytakeninx float64 = 0;
var gameturnmultiple float64 = 1;
var gametotalmultiple float64 = 1;
var interest float64 = 0.006;
var prefinterest float64 = 0.0045;
var basispoint float64 = 100;
var ladderStart[11]int;
var ladderEnd[11]int;
var snakeStart[10]int;
var snakeEnd[10]int;
var ladderHit [12]int;
var snakeHit[11]int;
var lIndex[101]int;
var sIndex[101]int;
var boardstates[101][101]int;
var boardlIndex[101][101]int;
var boardsIndex[101][101]int;
var boardladderHit[101][12]int;
var boardsnakeHit[101][11]int;
var boardpos[101]int;
var boardturnsofdiceforagame[101]int;
var boardturnsofgame[101]int;
var boardgameturnmultiple[101][31]float64;
var boardgametotalmultiple[101][31]float64;
var boardsnakesshuffled int = 0;
var boardladdersshuffled int = 0;
var commonboardgametotalmultipleleft[101]float64;
var commonboardgametotalmultipleadded[101]float64;
var commonboardgametotalmultipletakenin[101]float64;
var prefcommonboardgametotalmultipleleft[101][31]float64;
var prefcommonboardgametotalmultipleadded[101][31]float64;
var prefcommonboardgametotalmultipletakenin[101][31]float64;

func main() {
	ladderStart=[11]int{ 8,19,21,28,36,43,50,54,61,62,66 }
	ladderEnd=[11]int{ 26,38,82,53,57,77,91,88,99,96,87 }
	snakeStart=[10]int{ 98,93,83,69,68,64,59,52,48,46 }
	snakeEnd=[10]int{ 13,37,22,33,2,24,18,11,9,15 }
	var dice int;
	for dice = 0; dice <= 6; dice = dice + 1{fmt.Printf("%d",dice)}
	for dice = 0; dice <= 6; dice = dice + 1{
		fmt.Printf("%d", dice)
	}
	var dice1 int;
	dice1 = 10;
	for dice1 >= 1{
		fmt.Printf("%d", dice1);
		dice1 = dice1 - 1;
	}
	var ndice [10]int;
	for dice1 = 0; dice1<10; dice1++{
		ndice[dice1] = dice1 * 2;
		fmt.Printf("%d", ndice[dice1]);
	}
	if dice1 > dice{fmt.Printf(" %d %d", dice1, dice)}
	fmt.Printf(" %d", multiply(5, 3));
	rand.Seed(time.Now().UnixNano());
	//var pos int =0;
	//pos =0;
	for dice1 = 0; dice1<10; dice1++{
		pos = Dice(pos);
		fmt.Printf("%d ", pos);
	}
	fmt.Printf("\n");
	//runGame();
	//runGames();
	//boardrunGames();
	commonboardrunGames();
	//printGameStats();
	fmt.Printf("hello ramesh \n")
}
func multiply(m1 int, m2 int) int{
	fmt.Printf(" %d %d", m1, m2);
	return m1*m2;
}
func Dice(p int) int{
	//var pos int;
	p = p + rand.Intn(6) + 1;
	//printf(" %d", pos);
	return p;
}
func runGame(){
	for i:= 1; i<101; i++{
		states[i] = 0;
		lIndex[i] = 0;
		sIndex[i] = 0;
	}
	for i:= 1; i < 12; i++{ladderHit[i] = 0;}
	for i:= 1; i < 11; i++{snakeHit[i] = 0;}
	//getSnakesLadders(args);
	//System.out.println("turnsasmoneyatstart=" + turnsasmoney);
	//fillLadders();
	//fillSnakes();
	fillLaddersrand();
	fillSnakesrand();
	for turnsofdice<numberofdiceturns{
		pos = Dice(pos);
		TurnsofDice();
		if (pos>100){numberofGames();}
		pos = SnakesLadders();
	}
}
func TurnsofDice(){
	turnsofdice++;
	turnsofdiceforagame++;
}
func SnakesLadders()int{
	if (states[pos] != 0){
		if (lIndex[pos] != 0){ladderHit[lIndex[pos]]++;}
		if (sIndex[pos] != 0){snakeHit[sIndex[pos]]++;}
		//if(pos>states[pos])printf("snakehit %d",pos);
		//if(pos<states[pos])printf("ladderhit %d",pos);
		pos = states[pos];
	}
	return pos;
}
func fillLadders(){
	for i := 0; i<11; i++{
		//if (lIndex[ladderStart[i]] != 0)
		//{
		states[ladderStart[i]] = ladderEnd[i];
		lIndex[ladderStart[i]] = i + 1;
		//}
	}
}
func fillSnakes(){
	for i := 0; i<10; i++{
		//if (sIndex[snakeStart[i]] != 0)
		//{
		states[snakeStart[i]] = snakeEnd[i];
		sIndex[snakeStart[i]] = i + 1;
		//}
	}
}
func numberofGames(){
	pos = pos - 100;
	//System.out.println("turnsofgame="+turnsofgame);
	//System.out.println("position="+pos);
	turnsofgame++;
	//turnsasmoney = turnsasmoney*(((10000 - basispoint) / 10000) + (basispoint / (float64(turnsofdiceforagame) * 350)));
	logmultiple();
	turnsofdiceforagame = 0;
	//if(turnsofgame%35==0)numberofdiceturns=numberofdiceturns+100;
}
func printGameStats(){
	//fmt.Printf("turnsasmoney= %f \n", turnsasmoney);
	fmt.Printf(" gametotalmultiple %f \n", gametotalmultiple);
}
func logmultiple(){
	gameturnmultiple = ((10000 - basispoint) / 10000)*(turnsasmoneytakeninx + 1)- turnsasmoneytakeninx*(1 + interest)+ (basispoint / (float64(turnsofdiceforagame) * 350))*(turnsasmoneytakeninx + 1);
	//gameturnmultiple = ((10000 - basispoint) / 10000) + ((basispoint / (turnsofdiceforagame * 350)));
	//printf(" gameturnmultiple %f ", gameturnmultiple);
	gametotalmultiple = gametotalmultiple * gameturnmultiple;
	gameturnmultiple = 1;
	//loggameturnmultiple = log(((10000 - basispoint) / 10000) + ((basispoint / (turnsofdiceforagame * 350))));
	//printf(" loggameturnmultiple %e ", loggameturnmultiple);
	//loggametotalmultiple = loggametotalmultiple + loggameturnmultiple;
	//loggameturnmultiple = 0;
	//printf(" loggametotalmultiple %e ", loggametotalmultiple);
}
func runGames(){
	for i:= 0; i < 3; i++{
		turnsasmoneytakeninx = float64(i);
		//printf(" turnsasmoneytakeninx %f ", turnsasmoneytakeninx);
		fmt.Printf(" %f ", turnsasmoneytakeninx);
		runGame();
		printGameStats();
		turnsofdice = 0;
		turnsofgame = 0;
	}
}
func fillLaddersrand(){
	for i:= 0; i<11; i++{
		if (rand.Intn(6) > 2){
			states[ladderStart[i]] = ladderEnd[i];
			lIndex[ladderStart[i]] = i + 1;
		}
	}
}
func fillSnakesrand(){
	for i:= 0; i<10; i++{
		if (rand.Intn(6) > 2){
			states[snakeStart[i]] = snakeEnd[i];
			sIndex[snakeStart[i]] = i + 1;
		}
	}
}
func boardrunGame(){
	gameturnmultiple = 1;
	turnsasmoneytakeninx = 0;
	for h:= 1; h<101; h++{
		boardpos[h] = 1;
		boardturnsofdiceforagame[h] = 0;
		boardturnsofgame[h] = 0;
		//boardgameturnmultiple[h] = 1;
		//boardgametotalmultiple[h] = 1;
		//commonboardgametotalmultipletakenin[h] = 0;
		for i:= 1; i < 31; i++{
			//boardgameturnmultiple[h][i] = 1;
			boardgametotalmultiple[h][i] = 1;
			//if (h<2)commonboardgametotalmultipleleft[i] = 1;
			//if (h<2)commonboardgametotalmultipleadded[i] = 1;
		}

		for i:= 1; i<101; i++{
			boardstates[h][i] = 0;
			boardlIndex[h][i] = 0;
			boardsIndex[h][i] = 0;
		}

		for i:= 1; i < 12; i++{boardladderHit[h][i] = 0;}
		for i:= 1; i < 11; i++{boardsnakeHit[h][i] = 0;}
		boardfillLaddersrand(h);
		boardfillSnakesrand(h);
	}
	for turnsofdice<numberofdiceturns{
		boardDice();
		boardTurnsofDice();
		for h := 1; h<101; h++{
			if (boardpos[h] > 100){boardnumberofGames1(h);}
		}
		boardSnakesLadders();
		boardshuffleSnakesLaddersrand();
	}
	//commonboardlogtotalmultiple();
}
func boardfillLaddersrand(h int){
	for i:= 0; i<11; i++{
		if (rand.Intn(6) > 2){
			boardstates[h][ladderStart[i]] = ladderEnd[i];
			boardlIndex[h][ladderStart[i]] = i + 1;
		}
	}	
}
func boardfillSnakesrand(h int){
	for i:= 0; i<10; i++{
		if (rand.Intn(6) > 2){
			boardstates[h][snakeStart[i]] = snakeEnd[i];
			boardsIndex[h][snakeStart[i]] = i + 1;
		}
	}
}
func boardDice(){
	d:= rand.Intn(6) + 1;
	for h := 1; h<101; h++{
		boardpos[h] = boardpos[h] + d;
		//printf(" %d", pos);
	}
}
func boardTurnsofDice(){
	turnsofdice++;
	for h := 1; h<101; h++{
		boardturnsofdiceforagame[h]++;
	}
}
func boardSnakesLadders(){
	for h := 1; h<101; h++{
		//if(boardpos[h]>100){fmt.Printf("error");}
		if (boardstates[h][boardpos[h]] != 0){
			if (boardlIndex[h][boardpos[h]] != 0){boardladderHit[h][boardlIndex[h][boardpos[h]]]++;}
			if (boardsIndex[h][boardpos[h]] != 0){boardsnakeHit[h][boardsIndex[h][boardpos[h]]]++;}
			//if(pos>states[pos])printf("snakehit %d",pos);
			//if(pos<states[pos])printf("ladderhit %d",pos);
			boardpos[h] = boardstates[h][boardpos[h]];
		}
	}
}
func boardlogmultiple(h int){
	for i:= 1; i < 11; i++{
		turnsasmoneytakeninx = float64(i - 1);
		gameturnmultiple = ((10000 - basispoint) / 10000)*(turnsasmoneytakeninx + 1)- turnsasmoneytakeninx*(1 + interest)+ (basispoint / (float64(boardturnsofdiceforagame[h]) * 350))*(turnsasmoneytakeninx + 1);
		boardgametotalmultiple[h][i] = boardgametotalmultiple[h][i] * gameturnmultiple;
		gameturnmultiple = 1;
	}
	//boardgameturnmultiple[h] = ((10000 - basispoint) / 10000)*(turnsasmoneytakeninx + 1)
	//- turnsasmoneytakeninx*(1 + interest)
	//+ ((basispoint / (boardturnsofdiceforagame[h] * 350)))*(turnsasmoneytakeninx + 1);
	//boardgametotalmultiple[h] = boardgametotalmultiple[h] * boardgameturnmultiple[h];
	//boardgameturnmultiple[h] = 1;
}
func boardnumberofGames1(h int){
	boardpos[h] = boardpos[h] - 100;
	//System.out.println("turnsofgame="+turnsofgame);
	//System.out.println("position="+pos);
	boardturnsofgame[h]++;
	//turnsasmoney = turnsasmoney*(((10000 - basispoint) / 10000) + (basispoint / (turnsofdiceforagame * 350)));
	boardlogmultiple(h);
	//commonboardlogmultiple(h);
	//prefcommonboardlogmultiple(h);
	//prefcommonboardlogmultiple1(h);
	boardturnsofdiceforagame[h] = 0;
	//if(turnsofgame%35==0)numberofdiceturns=numberofdiceturns+100;
}
func boardprintGameStats(){
	for h:= 1; h<101; h++{
		for i:= 1; i < 11; i++{
			fmt.Printf(" %.1f ", boardgametotalmultiple[h][i]);
		}
		fmt.Printf("\n");
		//}
		//for (int h = 1; h < 101; h++)
		//{
		//for i:= 1; i < 11; i++{fmt.Printf("L %d = %d \n", i, boardladderHit[h][i]);}
		//for i:= 1; i <= 10; i++{fmt.Printf("S %d = %d \n", i, boardsnakeHit[h][i]);}
	}
}
func boardrunGames(){
	for i:= 0; i < 3; i++{
		boardrunGame();
		turnsofdice = 0;
		boardprintGameStats();
		//prefcommonboardprintGameStats1();
	}
}
func boardshuffleSnakesLaddersrand(){
	for h:= 1; h<101; h++{
		if ((h * 1) == (rand.Intn(100) + 1)){
			//printf("shuffling snakes %d ", h);
			boardsnakesshuffled++;
			boardladdersshuffled++;
			for i:= 1; i<101; i++{
				boardstates[h][i] = 0;
				boardlIndex[h][i] = 0;
				boardsIndex[h][i] = 0;
			}
			boardfillSnakesrand(h);
			boardfillLaddersrand(h);
		}
	}
}
func commonboardrunGame(){
	gameturnmultiple = 1;
	turnsasmoneytakeninx = 0;
	for h:= 1; h<101; h++{
		boardpos[h] = 1;
		boardturnsofdiceforagame[h] = 0;
		boardturnsofgame[h] = 0;
		//boardgameturnmultiple[h] = 1;
		//boardgametotalmultiple[h] = 1;
		commonboardgametotalmultipletakenin[h] = 0;
		for i:= 1; i < 31; i++{
			//boardgameturnmultiple[h][i] = 1;
			boardgametotalmultiple[h][i] = 0.01;
			if (h<2){commonboardgametotalmultipleleft[i] = 1;}
			if (h<2){commonboardgametotalmultipleadded[i] = 1;}
		}

		for i:= 1; i<101; i++{
			boardstates[h][i] = 0;
			boardlIndex[h][i] = 0;
			boardsIndex[h][i] = 0;
		}

		for i:= 1; i < 12; i++{boardladderHit[h][i] = 0;}
		for i:= 1; i < 11; i++{boardsnakeHit[h][i] = 0;}
		boardfillLaddersrand(h);
		boardfillSnakesrand(h);
	}
	for turnsofdice<numberofdiceturns{
		boardDice();
		boardTurnsofDice();
		for h := 1; h<101; h++{
			if (boardpos[h] > 100){commonboardnumberofGames(h);}
		}
		boardSnakesLadders();
		boardshuffleSnakesLaddersrand();
	}
	//commonboardlogtotalmultiple();
}
func commonboardnumberofGames(h int){
	boardpos[h] = boardpos[h] - 100;
	//System.out.println("turnsofgame="+turnsofgame);
	//System.out.println("position="+pos);
	boardturnsofgame[h]++;
	//turnsasmoney = turnsasmoney*(((10000 - basispoint) / 10000) + (basispoint / (turnsofdiceforagame * 350)));
	//boardlogmultiple(h);
	commonboardlogmultiple(h);
	//prefcommonboardlogmultiple(h);
	//prefcommonboardlogmultiple1(h);
	boardturnsofdiceforagame[h] = 0;
	//if(turnsofgame%35==0)numberofdiceturns=numberofdiceturns+100;
}
func commonboardlogmultiple(h int){
	for i:= 1; i < 11; i++{
		commonboardgametotalmultipleadded[i] = commonboardgametotalmultipleadded[i] - boardgametotalmultiple[h][i];
		turnsasmoneytakeninx = float64(i - 1);
		gameturnmultiple = ((10000 - basispoint) / 10000)*(turnsasmoneytakeninx + 1)- turnsasmoneytakeninx*(1 + interest)+ (basispoint / (float64(boardturnsofdiceforagame[h]) * 350))*(turnsasmoneytakeninx + 1);
		boardgametotalmultiple[h][i] = boardgametotalmultiple[h][i] * gameturnmultiple;
		commonboardallocatemultiple(h, i);
		//commonboardgametotalmultiple[i] = commonboardgametotalmultiple[i] + boardgametotalmultiple[h][i];
		//boardgametotalmultiple[h][i] = commonboardgametotalmultiple[i] / 10000;
		//commonboardgametotalmultiple[i] = commonboardgametotalmultiple[i] - boardgametotalmultiple[h][i];
		gameturnmultiple = 1;
	}
}
func commonboardprintGameStats(){
	for i:= 1; i < 11; i++{
			fmt.Printf(" %.1f ", commonboardgametotalmultipleadded[i]);
	}
	fmt.Printf("\n");
	for i:= 1; i < 11; i++{
			fmt.Printf(" %.1f ", commonboardgametotalmultipleleft[i]);
	}
	fmt.Printf("\n");
	for h:= 1; h<101; h++{
		fmt.Printf(" %.1f ", boardgametotalmultiple[h][1]);
	}
	fmt.Printf("\n");
}
func commonboardrunGames(){
	for i:= 0; i < 3; i++{
		//commonboardrunGame();
		commonboardrunGame2();
		turnsofdice = 0;
		//commonboardprintGameStats();
		prefcommonboardprintGameStats1();
	}
}
func commonboardallocatemultiple(h int, i int){
	//if (boardgametotalmultiple[h][i] > (commonboardgametotalmultipleleft[i]))
	if (boardgametotalmultiple[h][i] > (commonboardgametotalmultipleadded[i] / 3)){
		commonboardgametotalmultipleleft[i] = commonboardgametotalmultipleleft[i] + (boardgametotalmultiple[h][i] / 10);
		boardgametotalmultiple[h][i] = boardgametotalmultiple[h][i] * 0.9;
	}else if (boardgametotalmultiple[h][i] > (commonboardgametotalmultipleadded[i] / 50) && (boardgametotalmultiple[h][i] < commonboardgametotalmultipleleft[i] * 4)){
		commonboardgametotalmultipleleft[i] = commonboardgametotalmultipleleft[i] - (boardgametotalmultiple[h][i] / 4);
		boardgametotalmultiple[h][i] = boardgametotalmultiple[h][i] * 1.25;
	}else if (boardgametotalmultiple[h][i] < (commonboardgametotalmultipleadded[i] / 3000)){
		commonboardgametotalmultipleleft[i] = commonboardgametotalmultipleleft[i] + boardgametotalmultiple[h][i];
		boardgametotalmultiple[h][i] = commonboardgametotalmultipleleft[i] / 1000;
		commonboardgametotalmultipleleft[i] = commonboardgametotalmultipleleft[i] - boardgametotalmultiple[h][i];
	}
	commonboardgametotalmultipleadded[i] = commonboardgametotalmultipleadded[i] + boardgametotalmultiple[h][i];
}
func commonboardrunGame2(){
	gameturnmultiple = 1;
	turnsasmoneytakeninx = 0;
	for h:= 1; h<101; h++{
		boardpos[h] = 1;
		boardturnsofdiceforagame[h] = 0;
		boardturnsofgame[h] = 0;
		//boardgameturnmultiple[h] = 1;
		//boardgametotalmultiple[h] = 1;
		//commonboardgametotalmultipletakenin[h] = 0;
		//prefcommonboardgametotalmultipleleft[h] = 1;
		//prefcommonboardgametotalmultipleadded[h] = 1;
		for i:= 1; i < 31; i++{
			//boardgameturnmultiple[h][i] = 1;
			boardgametotalmultiple[h][i] = 0.01;
			prefcommonboardgametotalmultipleleft[h][i] = 1;
			prefcommonboardgametotalmultipleadded[h][i] = 1;
			prefcommonboardgametotalmultipletakenin[h][i] = 0;
			//if (h<2){commonboardgametotalmultipleleft[i] = 1;}
			//if (h<2){commonboardgametotalmultipleadded[i] = 1;}
		}

		for i:= 1; i<101; i++{
			boardstates[h][i] = 0;
			boardlIndex[h][i] = 0;
			boardsIndex[h][i] = 0;
		}

		for i:= 1; i < 12; i++{boardladderHit[h][i] = 0;}
		for i:= 1; i < 11; i++{boardsnakeHit[h][i] = 0;}
		boardfillLaddersrand(h);
		boardfillSnakesrand(h);
	}
	for turnsofdice<numberofdiceturns{
		boardDice();
		boardTurnsofDice();
		for h := 1; h<101; h++{
			if (boardpos[h] > 100){commonboardnumberofGames2(h);}
		}
		boardSnakesLadders();
		boardshuffleSnakesLaddersrand();
	}
	//commonboardlogtotalmultiple();
}
func commonboardnumberofGames2(h int){
	boardpos[h] = boardpos[h] - 100;
	//System.out.println("turnsofgame="+turnsofgame);
	//System.out.println("position="+pos);
	boardturnsofgame[h]++;
	//turnsasmoney = turnsasmoney*(((10000 - basispoint) / 10000) + (basispoint / (turnsofdiceforagame * 350)));
	//boardlogmultiple(h);
	//commonboardlogmultiple(h);
	//prefcommonboardlogmultiple(h);
	prefcommonboardlogmultiple1(h);
	boardturnsofdiceforagame[h] = 0;
	//if(turnsofgame%35==0)numberofdiceturns=numberofdiceturns+100;
}
func prefcommonboardallocatemultiple1(h int, i int){
	//if (boardgametotalmultiple[h][i] > (commonboardgametotalmultipleadded[i]))
	if (boardgametotalmultiple[h][i] > (prefcommonboardgametotalmultipleleft[h][i]) * 50){
		prefcommonboardgametotalmultipleleft[h][i] = prefcommonboardgametotalmultipleleft[h][i] + (boardgametotalmultiple[h][i] / 10);
		boardgametotalmultiple[h][i] = boardgametotalmultiple[h][i] * 0.9;
	}else if (boardgametotalmultiple[h][i] < prefcommonboardgametotalmultipleleft[h][i] * 10){
		prefcommonboardgametotalmultipletakenin[h][i] = prefcommonboardgametotalmultipletakenin[h][i] + 0.25;
		//boardgametotalmultiple[h][i] = boardgametotalmultiple[h][i] * 1.25;
	}else if (boardgametotalmultiple[h][i] < (prefcommonboardgametotalmultipleleft[h][i] / 10)){
		prefcommonboardgametotalmultipleleft[h][i] = prefcommonboardgametotalmultipleleft[h][i] + boardgametotalmultiple[h][i];
		boardgametotalmultiple[h][i] = prefcommonboardgametotalmultipleleft[h][i] / 10;
		prefcommonboardgametotalmultipleleft[h][i] = prefcommonboardgametotalmultipleleft[h][i] - boardgametotalmultiple[h][i];
	}
	prefcommonboardgametotalmultipleadded[h][i] = prefcommonboardgametotalmultipleadded[h][i] + boardgametotalmultiple[h][i];
}
func prefcommonboardlogmultiple1(h int){
	for i:= 1; i < 11; i++{
		prefcommonboardgametotalmultipleadded[h][i] = prefcommonboardgametotalmultipleadded[h][i] - boardgametotalmultiple[h][i];
		turnsasmoneytakeninx = float64(i - 1)+prefcommonboardgametotalmultipletakenin[h][i];
		//turnsasmoneytakeninx = commonboardgametotalmultipletakenin[h];
		gameturnmultiple = ((10000 - basispoint) / 10000)*(turnsasmoneytakeninx + 1)- turnsasmoneytakeninx*(1 + prefinterest)+ (basispoint / (float64(boardturnsofdiceforagame[h]) * 350))*(turnsasmoneytakeninx + 1);
		boardgametotalmultiple[h][i] = boardgametotalmultiple[h][i] * gameturnmultiple;
		prefcommonboardallocatemultiple1(h, i);
		//commonboardgametotalmultiple[i] = commonboardgametotalmultiple[i] + boardgametotalmultiple[h][i];
		//boardgametotalmultiple[h][i] = commonboardgametotalmultiple[i] / 10000;
		//commonboardgametotalmultiple[i] = commonboardgametotalmultiple[i] - boardgametotalmultiple[h][i];
		prefcommonboardgametotalmultipletakenin[h][i] = 0;
		gameturnmultiple = 1;
	}
}
func prefcommonboardprintGameStats1(){
	//fmt.Printf(" %.1f ", logprefgamewholemultiple);
	//fmt.Printf(" %.1f ", commonboardgametotalmultipleadded[1]);
	//fmt.Printf("\n");
	//fmt.Printf(" %.1f ", commonboardgametotalmultipleleft[1]);
	for h:= 1; h<101; h++{
		fmt.Printf(" %.1f ", prefcommonboardgametotalmultipleadded[h][2]);
		fmt.Printf(" %.1f ", prefcommonboardgametotalmultipleleft[h][2]);
		fmt.Printf(" %.1f ", prefcommonboardgametotalmultipletakenin[h][2]);
		//fmt.Printf(" %.1f ", boardgametotalmultiple[h][1]);
	}
	fmt.Printf("\n");
	fmt.Printf("snakes shuffled %d ", boardsnakesshuffled);
	fmt.Printf("ladders shuffled %d \n", boardladdersshuffled);
	fmt.Printf("\n");
}

/*
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <time.h>
#include <string.h>

#define numberofboards 201
#define boardsize 101
#define numberofparts 101
int multiply(int m1, int m2);
int Dice(int pos);
int pos = 1, turnsofdice = 0, turnsofgame = 0;
int states[101];
long numberofdiceturns = 100000;
void runGame(), TurnsofDice(), fillLadders(), fillSnakes(), numberofGames(), printGameStats(), logmultiple(), runGames();
void fillSnakesrand(), fillLaddersrand(), printGameStatsrand();
void boardrunGames(), boardfillLaddersrand(), boardfillSnakesrand(), boardDice();
void boardTurnsofDice(), boardnumberofGames1(int b), boardlogmultiple(int c), boardprintGameStats();
void boardSnakesLadders(), boardshuffleSnakesrand(), boardshuffleLaddersrand();
void commonboardlogmultiple(int c), commonboardprintGameStats(), commonboardrunGames();
void commonboardlogtotalmultiple(), commonboardallocatemultiple(int c, int d);
void prefcommonboardlogmultiple1(int c), prefcommonboardprintGameStats1(), prefcommonboardallocatemultiple1(int c, int d);
void prefcommonboardallocatemultiple2(int c, int d), prefboardrunGames1(), prefcommonboardrunGames1();
void prefcommonboardlogwholemultiple1(), prefcommonboardallocatemultiple3(int c);
void prefcommonboardlogmultiple2(int c), boardnumberofGames2(int b), prefcommonboardlogtotalmultiple();
void prefcommonboardprintGameStats2(), prefcommonboardrunGames2(), prefboardrunGames2();
void prefboardfillLaddersrand(int h), prefboardfillSnakesrand(int h);
void prefboardshuffleSnakesrand(), prefboardshuffleLaddersrand();
void prefcommonboardlogwholemultiple0();
void prefcommonboardlogwholemultiple2();
void prefcommonboardprintGameStats3();
void prefboardshuffleSnakesLaddersrand();
//struct gamestate *gamestatestruct();
int SnakesLadders();
int ladderStart[] = { 8,19,21,28,36,43,50,54,61,62,66 };
int ladderEnd[] = { 26,38,82,53,57,77,91,88,99,96,87 };
int snakeStart[] = { 98,93,83,69,68,64,59,52,48,46 };
int snakeEnd[] = { 13,37,22,33,2,24,18,11,9,15 };
int ladderHit[12];
int snakeHit[11];
int lIndex[101];
int sIndex[101];
int boardstates[numberofboards][boardsize];
int boardlIndex[numberofboards][boardsize];
int boardsIndex[numberofboards][boardsize];
int boardladderHit[numberofboards][12];
int boardsnakeHit[numberofboards][11];
int boardpos[numberofboards];
int boardturnsofdiceforagame[numberofboards];
int boardturnsofgame[numberofboards];
//float boardgameturnmultiple[101];
//float boardgametotalmultiple[101];
float boardgameturnmultiple[101][31];
float boardgametotalmultiple[101][31];
float commonboardgametotalmultipleleft[31];
float commonboardgametotalmultipleadded[31];
float commonboardgametotalmultipletakenin[101];
float logprefwholemultiple = 0;
float logprefwholemultipletaken = 0;
float logprefwholemultipleadded = 1;
float logprefwholemultipleleft = 1;
float prefboardgametotalmultiple[numberofboards];
int snakesshuffled = 0;
int laddersshuffled = 0;
int turnsofdiceforagame = 0;
float turnsasmoney = 100;
int turnsasmoneytakeninx = 0;
float interest = 0.006;
float prefinterest = 0.0045;
float basispoint = 100;
float gameturnmultiple = 1;
float gametotalmultiple = 1;
float logprefgamewholemultiple = 1;
double loggameturnmultiple = 0;
double loggametotalmultiple = 0;
struct gamestate
{
	float logprefgamewholemultiple;
	float logprefwholemultipletaken;
	float logprefwholemultipleadded;
	float logprefwholemultipleleft;
	float gameprefinterest;
	int numberofgameboards;
	int gameboardsize;
	int gamenumberofparts;
	int boardpos[numberofboards];
	int boardturnsofdiceforagame[numberofboards];
	int boardturnsofgame[numberofboards];
	int boardstates[numberofboards][boardsize];
	int boardlIndex[numberofboards][boardsize];
	int boardsIndex[numberofboards][boardsize];
	int boardladderHit[numberofboards][12];
	int boardsnakeHit[numberofboards][11];
	float prefboardgametotalmultiple[numberofboards];
};
struct gamestate prefgamestate, p;
struct gamestate *gamestatestruct();
void writegamestate(struct gamestate *p1);
//char *writegamestatestr(struct gamestate *p1);
void writegamestatestr(struct gamestate *p1);
void printtime();
*/