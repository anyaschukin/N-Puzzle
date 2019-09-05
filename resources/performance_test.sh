echo "\x1b[1mLaunching N-Puzzle performance test ...\x1B[0m\n"
## if fail
## echo "Usage: '''go build''' to build the binary 'n-puzzle'. then ./performance_test.sh"

# 3
echo "\x1b[1mSize - 3\x1B[0m"
case=10
count=1
u=0
while [ $count -lt $(expr $case + 1) ]
do	
	echo "\r-- $count/$case --\c"
	output=$(python generator.py -u 5 >> rm_me.txt; ../n-puzzle rm_me.txt)
##	echo "$output"
	check="This puzzle is unsolvable."
	if [ "$output"="$check" ]
	then
		u=$(($u + 1))
	fi	
	count=$(($count + 1))
done
echo "\nUnsolvable correctly identified: $u/10" ## out of 10
##echo "$u"	

#while [ $count -lt $(expr $case + 1) ]
#do
#	echo "\r-- $count/$case --\c"
#	echo "\n"
#	time=$(python generator.py -u 5 >> rm_me.txt; ../n-puzzle rm_me.txt)
#	count=$(($count + 1))
#done

echo "Unsolvable correctly identified: " ## out of 10
echo "Solvable correctly solved: " ## out of 10
echo "Average solve time: "
echo "Min solve time: "
echo "Max solve time: "
