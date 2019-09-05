echo "\x1b[1mLaunching N-Puzzle performance test \x1B[0m...\n"
## if fail
## echo "Usage: '''go build''' to build the binary 'n-puzzle'. then ./performance_test.sh"

# 3
echo "\x1b[1mSize - 3\x1B[0m"

####  -- Unsolvable --  ####
case=10
count=0
u=0
while [ $count -lt $case ]
do
	echo ".\c"
	unsolvable=$(python generator.py -u 3 >> rm_me.txt; ../n-puzzle rm_me.txt)
	if [ "$unsolvable"="This puzzle is unsolvable." ]
	then
		u=$(($u + 1))
	fi	
	count=$(($count + 1))
	$(rm rm_me.txt)
done
if [ "$u" -lt "$count" ]
then
	echo "\x1b[31m"
else
	echo "\x1b[32m"
fi
echo "Unsolvable correctly identified: $u/$count\x1b[0m"

####  -- Solvable --  ####
case=10
count=0
solved=0
#tmin=60
#tmax=0
tcumulative=0
while [ $count -lt $case ]
do
#	echo ".\c"
	solvable=$(python generator.py -s 3 >> rm_me.txt; ../n-puzzle rm_me.txt)
	end=$( echo "$solvable" | tail -n -1)
	if [ "$end"=" You've finished n-puzzle!" ]
	then
		solved=$(($solved + 1))
	fi
	time=$( echo "$solvable" | tail -n -2 | head -n 1 | cut -d " " -f 3)
	echo "time: $time"
	

## s ms or µs
#	tcumulative=$(($tcumulative + $time))
#	echo "$tcumulative"

	count=$(($count + 1))
	$(rm rm_me.txt)
done

#echo "$time"


if [ "$u" -lt "$count" ]
then
	echo "\x1b[31m"
else
	echo "\x1b[32m"
fi
echo "Solvable correctly solved: $solved/$count\x1b[0m"
echo "Average solve time: "
echo "Min solve time: "
echo "Max solve time: "
