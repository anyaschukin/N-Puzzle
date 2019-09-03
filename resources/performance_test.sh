echo "\x1b[1mLaunching N-Puzzle performance test ...\x1B[0m\n"
## if fail
## echo "Usage: '''go build''' to build the binary 'n-puzzle'. then ./performance_test.sh"

#3
echo "\x1b[1mSize - 3\x1B[0m"
case=10
count=1
sum=0
while [ $count -lt $(expr $case + 1) ]
do
	echo "\r-- $count/$case --\c"
	arg1=$(python generator.py -s 5 >> rm_me.txt; ../n-puzzle rm_me.txt)
	echo arg1
	count=$(($count + 1))
done

echo "Unsolvable correctly identified: " ## out of 10
echo "Solvable correctly solved: " ## out of 10
echo "Average solve time: "
echo "Min solve time: "
echo "Max solve time: "

######################################

echo "\x1B[32mavg moves diff: $(($sum / $(($count - 1))))\x1B[0m"

#time of one
echo "$(time ./lem-in -t < 1 | tail -n 1)"


	arg1=$(./n-puzzle --flow-one > 1 && cat 1 | grep "#Here" | cut -d ' ' -f 8 | head -n 1)
	arg2=$(./lem-in -t < 1 | tail -n 2 | head -n 1 | cut -d ' ' -f 2)
	diff=$(($arg2 - $arg1))
	sum=$(($sum + $diff))
	count=$(($count + 1))
done
