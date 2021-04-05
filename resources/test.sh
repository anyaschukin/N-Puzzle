#### -- Print Header -- ####
RESET="\x1b[0m"
BRIGHT="\x1b[1m"
RED="\x1b[31m"
GREEN="\x1b[32m"
YELLOW="\x1b[33m"

printf "\E[H\E[2J" ## Clear screen
printf $BRIGHT
echo "Launching N-Puzzle Performance Test$RESET\n"

start=`date +%s`
## echo "Usage: '''go build''' to build the binary 'n-puzzle'. then ./performance_test.sh"


#### -- Config -- ####

MIN_SIZE=3			# 3 min
MAX_SIZE=4			# 5
TEST_CASES=2		# 5
UNSOLVABLE_TEST=1	# 0 = off, 1 = on
SOLVABLE_TEST=1		# 0 = off, 1 = on
UNIT_TEST=1			# 0 = off, 1 = on
RANDOM_TEST=1		# 0 = off, 1 = on
HEURISTIC="manhattan"


#### -- Print Config -- ####
echo "\t\x1b[4m-- Config --$RESET"
echo "Minimum size:\t\t$MIN_SIZE"
echo "Maximum size:\t\t$MAX_SIZE"
echo "Number of test cases:\t$TEST_CASES"
if [ "$UNSOLVABLE_TEST" != 0 ]
then
	echo "Unsolvable Tests:$GREEN\ton$RESET"
else	
	echo "Unsolvable Tests:$RED\toff$RESET"
fi
if [ "$SOLVABLE_TEST" != 0 ]
then
	echo "Solvable Tests:$GREEN\t\ton$RESET"
else	
	echo "Solvable Tests:$RED\t\toff$RESET"
fi
if [ "$UNIT_TEST" != 0 ]
then
	echo "Unit Tests:$GREEN\t\ton$RESET"
else	
	echo "Unit Tests:$RED\t\toff$RESET"
fi
if [ "$RANDOM_TEST" != 0 ]
then
	echo "Random Tests:$GREEN\t\ton$RESET"
else	
	echo "Random Tests:$RED\t\toff$RESET"
fi
echo "Heuristic:\t\t$HEURISTIC\n"


#### -- Test Function -- ####
unit_test()
{
	SOLVABLE=$1
	UNIT=$2
	case=$TEST_CASES
	solved=0
	count=0
	best=42
	worst=0
	tcumulative=0
	count=0

	## Loop cases
	while [ $count -lt $case ]
	do
		count=$(($count + 1))
		test_num=$(($test_num + 1))

		## Run
		if [ "$SOLVABLE" == "Unsolvable" ]
		then ## Unsolvable
			if [ "$UNIT" == "Unit" ]
			then ## Unit
				unit=$(echo "Boards/Unsolvable/$size/$size""u$count.txt")
				output=$(../n-puzzle -h=$HEURISTIC $unit)
			else ## Random
				output=$(python generator.py -u $size >> rm_me.txt; ../n-puzzle -h=$HEURISTIC rm_me.txt)
			fi
			unsolvable=$(echo "$output" | tail -n -2 | head -n 1)
			if [ "$unsolvable" = "This puzzle is unsolvable." ]
			then
				solved=$(($solved + 1))
				echo "$GREEN.$RESET\c"
			else	
				echo "$RED.$RESET\c"
				if [ -f "rm_me.txt" ]
				then
					rm rm_me.txt
				fi
				continue
			fi
			time=$(echo "$output" | tail -n -1 | cut -d " " -f 3)
		else ## Solvable
			if [ "$UNIT" == "Unit" ]
			then ## Unit
				unit=$(echo "Boards/Solvable/$size/$size""s$count.txt")
				output=$(../n-puzzle -h=$HEURISTIC $unit)
			else ## Random
				output=$(python generator.py -s $size >> rm_me.txt; ../n-puzzle -h=$HEURISTIC rm_me.txt)
			fi
			end=$(echo "$output" | tail -n -1)
			if [ "$end" != "You've finished n-puzzle!" ]
			then
				echo "$RED.$RESET\c"
				continue
			else
				solved=$(($solved + 1))
				echo "$GREEN.$RESET\c"
			fi
			time=$(echo "$output" | tail -n -2 | head -n 1 | cut -d " " -f 3)
		fi

		## Time
		prefix=$(echo "$time" | rev | cut -c-1-2 | rev | cut -c-1-1)
		if [ "$prefix" = "m" ]
		then
			time=$(echo "$time" | rev | cut -c3-42 | rev)
			time=$(echo "scale = 9; ($time / 1000)" | bc)	
		elif [ "$prefix" = "µ" ]
		then
			time=$(echo "$time" | rev | cut -c3-42 | rev)
			time=$(echo "scale = 9; ($time / 1000000)" | bc)

		elif [ "$prefix" = "n" ]
		then
			time=$(echo "$time" | rev | cut -c3-42 | rev)
			time=$(echo "scale = 9; ($time / 1000000000)" | bc)
		else
			time=$(echo "$time" | rev | cut -c2-42 | rev)
		fi
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

		if [ -f "rm_me.txt" ]
		then
			rm rm_me.txt
		fi
	done

	## Print results
	if [ "$solved" != 0 ]
	then
		mean=$(echo "scale = 9; $tcumulative / $solved" | bc)
	else
		mean="$RED Failed$RESET"
	fi
	if [ "$worst" = 0 ]
	then
		worst="$RED""Failed$RESET"
	fi
	if [ "$best" = 42 ]
	then
		best="$RED""Failed$RESET"
	fi

	if [ "$solved" = 0 ]
	then
		echo $RED
	elif [ "$solved" -lt "$count" ]
	then
		echo $YELLOW
	else
		echo $GREEN
	fi
	printf "%s %s: \t %s/%s$RESET\n" $SOLVABLE $UNIT $solved $count 
	echo "Solve time in seconds"
	echo "\t\t Worst: $worst"
	echo "\t\t Mean:  $mean"
	echo "\t\t Best:  $best"
}


#### -- Test -- ####
size=$MIN_SIZE
test_num=0
if [ -f "rm_me.txt" ]
then
	rm rm_me.txt
fi
## Loop size
while [ $size -lt $(expr $MAX_SIZE + 1) ]
do
	echo $BRIGHT
	echo "Size - $size$RESET"

	if [ "$UNSOLVABLE_TEST" != 0 -a "$UNIT_TEST" != 0 -a "$size" -gt 2 -a "$size" -lt 10 ]
	then
		unit_test Unsolvable Unit
	fi
	if [ "$UNSOLVABLE_TEST" != 0  -a "$RANDOM_TEST" != 0 ]
	then
		unit_test Unsolvable Random
	fi
	if [ "$SOLVABLE_TEST" != 0 -a "$UNIT_TEST" != 0 -a "$size" -gt 2 -a "$size" -lt 10 ]
	then
		unit_test Solvable Unit
	fi
	if [ "$SOLVABLE_TEST" != 0 -a "$RANDOM_TEST" != 0 ]
	then
		unit_test Solvable Random
	fi

	echo
	size=$(($size + 1))
done

end=`date +%s`
runtime=$((end-start))
echo $BRIGHT
echo "N-Puzzle performance test finished, $test_num tests run in $runtime seconds.\n"
