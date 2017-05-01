#---------MAIN PATHS---------
class-dump=bin/
src-path=./
main-path=./com/theaigames/tictactoe/Tictactoe.java 
main-class=com.theaigames.tictactoe.Tictactoe

#---------BOTS---------
base-path=../bot_monte_base/
bot_monte_base="go run $(base-path)bot.go $(base-path)board.go $(base-path)fields.go $(base-path)monte.go" #< $(base-path)moves2.txt

ucb-path=../bot_monte_UCB/
bot_monte_UCB="go run $(ucb-path)bot.go $(ucb-path)board.go $(ucb-path)fields.go $(ucb-path)monte.go" #< $(ucb-path)moves2.txt

minimax-path=../bot_minimax/
bot_minimax="go run $(minimax-path)bot.go $(minimax-path)board.go $(minimax-path)fields.go" #< $(minimax-path)moves2.txt   #MINIMAX

bot_monte_bitvec-path=../bot_monte_bitvec/
bot_monte_bitvec="go run $(bot_monte_bitvec-path)bot.go $(bot_monte_bitvec-path)board.go $(bot_monte_bitvec-path)fields.go $(bot_monte_bitvec-path)monte.go" #< $(bot_monte_bitvec-path)moves2.txt   #MINIMAX

all: compile run

compile:
	javac -d $(class-dump) -cp $(src-path) $(main-path)

run:
	java -cp $(class-dump):$(src-path) $(main-class) $(bot_monte_base) $(bot_monte_bitvec) #2>err.txt 1>out.txt

bot:
	$(bot1-run)

