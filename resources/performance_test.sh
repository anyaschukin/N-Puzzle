echo "\x1b[1mLaunching N-Puzzle performance test \x1B[0m...\n"
## echo "Usage: '''go build''' to build the binary 'n-puzzle'. then ./performance_test.sh"

####  -- Config --  ####
test_cases=10
min_size=3
max_size=5

echo "\x1b[4m-- Config --\x1b[0m"
echo "Number of test cases = $test_cases"
echo "Minimum size = $min_size"
echo "Maximum size = $max_size\n"

####  -- Test --  ####
size=$min_size
if [ -f "rm_me.txt" ]
then
	$(rm rm_me.txt)
fi
while [ $size -lt $(expr $max_size + 1) ]
do
	echo "\x1b[1mSize - $size\x1B[0m"

	####  -- Unsolvable --  ####
	case=$test_cases
	count=0
	u=0
	while [ $count -lt $case ]
	do
		echo ".\c"
		count=$(($count + 1))
		unsolvable=$(python generator.py -u $size >> rm_me.txt; ../n-puzzle rm_me.txt)
		if [ "$unsolvable"="This puzzle is unsolvable." ]
		then
			u=$(($u + 1))
		fi	
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
	case=$test_cases
	count=0
	solved=0
	best=42
	worst=0
	tcumulative=0
	while [ $count -lt $case ]
	do
		echo ".\c"
		count=$(($count + 1))
		solvable=$(python generator.py -s $size >> rm_me.txt; ../n-puzzle rm_me.txt)
#		echo "$solvable"
		end=$(echo "$solvable" | tail -n -1)
		if [ "$end" != " You've finished n-puzzle!" ]
		then
			$(rm rm_me.txt)
			continue
		else
			solved=$(($solved + 1))
		fi
		time=$(echo "$solvable" | tail -n -2 | head -n 1 | cut -d " " -f 3)
	#	echo "time: $time"
		prefix=$(echo "$time" | rev | cut -c-1-2 | rev | cut -c-1-1)
		if [ "$prefix" = "m" ]
		then
			time=$(echo "$time" | rev | cut -c3-42 | rev)
			time=$(echo "scale = 9; ($time / 1000)" | bc)	
			tcumulative=$(echo "scale = 9; $tcumulative + $time" | bc)
			time_up=$(echo "scale = 0; $time * 1000000000" | bc | cut -d "." -f 1)
			worst_up=$(echo "scale = 0; $worst * 1000000000" | bc | cut -d "." -f 1)
			best_up=$(echo "scale = 0; $best * 1000000000" | bc | cut -d "." -f 1)
			if [ "$time_up" -gt "$worst_up" ]
			then
				worst=$time
			fi
			if [ "$time_up" -lt "$best_up" ]
			then
				best=$time
			fi
		elif [ "$prefix" = "µ" ]
		then
			time=$(echo "$time" | rev | cut -c3-42 | rev)
			time=$(echo "scale = 9; ($time / 1000000)" | bc)	
			tcumulative=$(echo "scale = 9; $tcumulative + $time" | bc)	
			time_up=$(echo "scale = 0; $time * 1000000000" | bc | cut -d "." -f 1)
			worst_up=$(echo "scale = 0; $worst * 1000000000" | bc | cut -d "." -f 1)
			best_up=$(echo "scale = 0; $best * 1000000000" | bc | cut -d "." -f 1)
			if [ "$time_up" -gt "$worst_up" ]
			then
				worst=$time
			fi
			if [ "$time_up" -lt "$best_up" ]
			then
				best=$time
			fi
		else
			time=$(echo "$time" | rev | cut -c2-42 | rev)
			tcumulative=$(echo "$tcumulative + $time" | bc)
			time_up=$(echo "scale = 0; $time * 1000000000" | bc | cut -d "." -f 1)
			worst_up=$(echo "scale = 0; $worst * 1000000000" | bc | cut -d "." -f 1)
			best_up=$(echo "scale = 0; $best * 1000000000" | bc | cut -d "." -f 1)
			if [ "$time_up" -gt "$worst_up" ]
			then
				worst=$time
			fi
			if [ "$time_up" -lt "$best_up" ]
			then
				best=$time
			fi
		fi
	#	echo "time_up: $time_up"
	#	echo "worst_up: $worst_up"
	#	echo "best_up: $best_up"
	#	echo "tcumulative: $tcumulative"
	#	echo "worst: $worst"
	#	echo "best: $best"
		$(rm rm_me.txt)
	done

	mean=$(echo "scale = 9; $tcumulative / $solved" | bc)
	if [ "$mean" = "" ]
	then
		mean="\x1b[31mFailed\x1b[0m"
	fi
	if [ "$worst" = 0 ]
	then
		worst="\x1b[31mFailed\x1b[0m"
	fi
	if [ "$best" = 42 ]
	then
		best="\x1b[31mFailed\x1b[0m"
	fi

	#echo "tcumulative: $tcumulative"
	#echo "count: $count"
	#echo "mean: $mean"

	if [ "$solved" -lt "$count" ]
	then
		echo "\x1b[31m"
	else
		echo "\x1b[32m"
	fi
	echo "Solvable correctly solved: $solved/$count\x1b[0m"
	echo "Solve time in seconds:"
	echo "\tMean: \t$mean"
	echo "\tWorst: \t$worst"
	echo "\tBest: \t$best\n"
	size=$(($size + 1))
done
echo "\n\x1b[5mN-Puzzle performance test finished. Have a nice day!\x1B[0m"
