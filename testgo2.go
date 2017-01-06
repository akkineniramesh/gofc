package main

import "fmt"
//import "math"
import "math/rand"
import "time"

var pos int= 1;
var turnsofdice int= 0;
var turnsofgame int= 0
var turnsofdiceforagame int= 0;
var states[101]int;
var numberofdiceturns int= 100000;
var turnsasmoney float64 = 100;
//turnsasmoneytakeninx:= 0;
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
	runGame();
	printGameStats();
	fmt.Println("hello ramesh 1\n")
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
	fillLadders();
	fillSnakes();
	//fillLaddersrand();
	//fillSnakesrand();
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
	turnsasmoney = turnsasmoney*(((10000 - basispoint) / 10000) + (basispoint / (float64(turnsofdiceforagame) * 350)));
	//logmultiple();
	turnsofdiceforagame = 0;
	//if(turnsofgame%35==0)numberofdiceturns=numberofdiceturns+100;
}
func printGameStats(){
	fmt.Printf("turnsasmoney= %f \n", turnsasmoney);
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