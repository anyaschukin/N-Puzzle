echo "\x1b[1mLaunching N-Puzzle performance test ...\x1B[0m\n"
## if fail
## echo "Usage: '''go build''' to build the binary 'n-puzzle'. then ./performance_test.sh"

# 3
echo "\x1b[1mSize - 3\x1B[0m"

# Unsolvable
case=10
count=0
u=0
while [ $count -lt $case ]
do	
	echo "\r-- $count/$case --\c"
	output=$(python generator.py -u 3 >> rm_me.txt; ../n-puzzle rm_me.txt)
	unsolvable="This puzzle is unsolvable."
	if [ "$output"="$unsolvable" ]
	then
		u=$(($u + 1))
	fi	
	count=$(($count + 1))
done
if [ "$u" -lt "$count" ]
then
	echo "\x1b[31m"
else
	echo "\x1b[32m"
fi
echo "\nUnsolvable correctly identified: $u/$count\x1b[0m"

# Solvable
case=10
count=0
solved=0
#while [ $count -lt $(expr $case + 1) ]
#do
#	echo "\r-- $count/$case --\c"
#	echo "\n"
#	time=$(python generator.py -u 5 >> rm_me.txt; ../n-puzzle rm_me.txt)
#	count=$(($count + 1))
#done

echo "Solvable correctly solved: " ## out of 10
echo "Average solve time: "
echo "Min solve time: "
echo "Max solve time: "
