#### -- N-Puzzle Performance Test -- ####
## Runs Unsolvable & Solvable Unit tests from boards/
## & Random tests using boards/generator.py
## Then plots solve time, moves, size & time complexity by heuristic
## To run: ./test.sh


#### -- Config -- ####
MIN_SIZE=3			# 3 min
MAX_SIZE=4			# 4 default
TEST_CASES=5		# 11 unit cases available for sizes 3 to 9, in boards/
UNSOLVABLE_TEST=0	# 0 = off, 1 = on
SOLVABLE_TEST=1		# 0 = off, 1 = on
UNIT_TEST=1			# 0 = off, 1 = on
RANDOM_TEST=1		# 0 = off, 1 = on
declare -a heuristics=("manhattan" "nilsson" "outRowCol" "hamming" "euclidean")


#### -- Print Header -- ####
RESET="\x1b[0m"
BRIGHT="\x1b[1m"
RED="\x1b[31m"
GREEN="\x1b[32m"
YELLOW="\x1b[33m"
CLEAR_LINE="\r"
UNDERLINE="\x1b[4m"

printf "\E[H\E[2J" # Clear screen
printf $BRIGHT
echo "Launching N-Puzzle Performance Test ...$RESET\n"
go build

start=`date +%s`


#### -- Print Config -- ####
echo "$BRIGHT$UNDERLINE""Configuration$RESET"
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
printf "Heuristics:\t\t"
for heuristic in "${heuristics[@]}"
do
   printf "$heuristic\n\t\t\t"
done
echo


#### -- Cleanup .csv -- ####
if [ -e solve_time.csv ]
then
    rm solve_time.csv
fi
if [ -e moves.csv ]
then
    rm moves.csv
fi
if [ -e size_complexity.csv ]
then
    rm size_complexity.csv
fi
if [ -e time_complexity.csv ]
then
    rm time_complexity.csv
fi


#### -- Test Function -- ####
unit_test()
{
	SOLVABLE=$1
	UNIT=$2
	HEURISTIC=$3
	case=$TEST_CASES
	solved=0
	count=0
	best=42
	worst=0
	tcumulative=0

	## Loop cases
	while [ $count -lt $case ]
	do
		test_num=$(($test_num + 1))
		count=$(($count + 1))
		if [ "$UNIT" == "Unit" -a "$count" -gt 10 ]
		then
			break
		fi

		## Print ...
		type=$(echo $SOLVABLE $UNIT)
		if [ "$solved" = 0 ]
		then
			printf "%s %s ...$CLEAR_LINE" $SOLVABLE $UNIT
		elif [ "$solved" -lt $(($count - 1)) ]
		then
			printf "$YELLOW%s %s%-4s\t%s/%s ...    $RESET $CLEAR_LINE" $SOLVABLE $UNIT ":" $solved $(($count - 1))
		else
			printf "$GREEN%s %s%-4s\t%s/%s OK ...    $RESET $CLEAR_LINE" $SOLVABLE $UNIT ":" $solved $(($count - 1))
		fi

		## Run
		if [ "$SOLVABLE" == "Unsolvable" ]
		then ## Unsolvable
			if [ "$UNIT" == "Unit" ]
			then ## Unit
				unit=$(echo "boards/Unsolvable/$size/$size""u$count.txt")
				output=$(./n-puzzle -h=$HEURISTIC $unit)
			else ## Random
				output=$(python boards/generator.py -u $size >> rm_me.txt; ./n-puzzle -h=$HEURISTIC rm_me.txt)
			fi
			unsolvable=$(echo "$output" | tail -n -2 | head -n 1)
			if [ "$unsolvable" = "This puzzle is unsolvable." ]
			then
				solved=$(($solved + 1))
				total_solved=$(($total_solved + 1))
			else	
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
				unit=$(echo "boards/Solvable/$size/$size""s$count.txt")
				output=$(./n-puzzle -h=$HEURISTIC $unit)
			else ## Random
				output=$(python boards/generator.py -s $size >> rm_me.txt; ./n-puzzle -h=$HEURISTIC rm_me.txt)
			fi
			end=$(echo "$output" | tail -n -1)
			if [ "$end" != "You've finished n-puzzle!" ]
			then
				continue
			else
				solved=$(($solved + 1))
				total_solved=$(($total_solved + 1))
			fi
			time=$(echo "$output" | tail -n -2 | head -n 1 | cut -d " " -f 3)
		fi

		## Time
		prefix=$(echo "$time" | rev | cut -c-1-2 | rev | cut -c-1-1)
		if [ "$prefix" = "m" ]		# Milliseconds
		then
			time=$(echo "$time" | rev | cut -c3-42 | rev)
			time=$(echo "scale = 9; ($time / 1000)" | bc)	
		elif [ "$prefix" = "µ" ]	# Microseconds
		then
			time=$(echo "$time" | rev | cut -c3-42 | rev)
			time=$(echo "scale = 9; ($time / 1000000)" | bc)

		elif [ "$prefix" = "n" ]	# Nanoseconds
		then
			time=$(echo "$time" | rev | cut -c3-42 | rev)
			time=$(echo "scale = 9; ($time / 1000000000)" | bc)
		elif [[ "$time" =~ "m" ]]	# Minutes
		then
			minutes=$(echo "$time" | cut -d "m" -f 1)
			seconds=$(echo "$time" | cut -d "m" -f 2 | rev | cut -c3-42 | rev)
			time=$(echo "scale = 9; $seconds + ($minutes * 60)" | bc)
		else						# Seconds
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

		## Write to .csv
		if [ "$SOLVABLE" != "Unsolvable" ]
		then
			moves=$(echo "$output" | grep moves | cut -d ":" -f 2)
			size_complexity=$(echo "$output" | grep Size | cut -d ":" -f 2)
			time_complexity=$(echo "$output" | grep Time | cut -d ":" -f 2)

			printf "%f" $time >> solve_time.csv
			printf "%d" $moves >> moves.csv
			printf "%d" $size_complexity >> size_complexity.csv
			printf "%d" $time_complexity >> time_complexity.csv
			if [ $count -lt $case -o "$UNIT" == "Unit" -a $RANDOM_TEST != "0" ]
			then
				printf ", " >> solve_time.csv
				printf ", " >> moves.csv
				printf ", " >> size_complexity.csv
				printf ", " >> time_complexity.csv
			else 
				printf "\n" >> solve_time.csv
				printf "\n" >> moves.csv
				printf "\n" >> size_complexity.csv
				printf "\n" >> time_complexity.csv
			fi
		fi


		## Print Solved
		type=$(echo $SOLVABLE $UNIT)
		if [ "$solved" = 0 ]
		then
			printf "$RED%s %s%-3s\t%s/%s ERROR    $RESET $CLEAR_LINE" $SOLVABLE $UNIT ":" $solved $count
		elif [ "$solved" -lt "$count" ]
		then
			printf "$YELLOW%s %s%-3s\t%s/%s          $RESET $CLEAR_LINE" $SOLVABLE $UNIT ":" $solved $count
		else
			printf "$GREEN%s %s%-3s\t%s/%s OK        $RESET $CLEAR_LINE" $SOLVABLE $UNIT ":" $solved $count
		fi
		
		## Cleanup
		if [ -f "rm_me.txt" ]
		then
			rm rm_me.txt
		fi
	done

	## Print Time
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
	echo
	echo "Solve time       Worst: $worst"
	echo "(Seconds)        Mean:  $mean"
	echo "                 Best:  $best"
	echo
}


#### -- Test -- ####
size=$MIN_SIZE
test_num=0
total_solved=0
if [ -f "rm_me.txt" ]
then
	rm rm_me.txt
fi

## Loop size
while [ $size -lt $(expr $MAX_SIZE + 1) ]
do
	if [ $size -gt 3 ]
	then
		declare -a heuristics=("manhattan" "nilsson")
	fi
	## Loop heuristic
	for heuristic in "${heuristics[@]}"
	do
		printf $BRIGHT
		printf $UNDERLINE
		echo "Size - $size,  Heuristic - $heuristic$RESET"
		echo
		if [ "$UNSOLVABLE_TEST" != 0 -a "$UNIT_TEST" != 0 -a "$size" -gt 2 -a "$size" -lt 10 ]
		then
			unit_test Unsolvable Unit $heuristic
		fi
		if [ "$UNSOLVABLE_TEST" != 0  -a "$RANDOM_TEST" != 0 ]
		then
			unit_test Unsolvable Random $heuristic
		fi
		if [ "$SOLVABLE_TEST" != 0 -a "$UNIT_TEST" != 0 -a "$size" -gt 2 -a "$size" -lt 10 ]
		then
			unit_test Solvable Unit $heuristic
		fi
		if [ "$SOLVABLE_TEST" != 0 -a "$RANDOM_TEST" != 0 ]
		then
			unit_test Solvable Random $heuristic
		fi
		echo
	done
	size=$(($size + 1))

	python3 test_plot.py
	## Cleanup .csv
	if [ -e solve_time.csv ]
	then
		rm solve_time.csv
	fi
	if [ -e moves.csv ]
	then
		rm moves.csv
	fi
	if [ -e size_complexity.csv ]
	then
		rm size_complexity.csv
	fi
	if [ -e time_complexity.csv ]
	then
		rm time_complexity.csv
	fi
done


## Print End
end=`date +%s`
runtime=$((end-start))
printf $BRIGHT
if [ "$runtime" == 1 ]
then
	echo "N-Puzzle performance test finished, total runtime $runtime second.\n"
else
	echo "N-Puzzle performance test finished, total runtime $runtime seconds.\n"
fi

if [ "$total_solved" == "$test_num" ]
then
	echo "Passed$GREEN $total_solved / $test_num$RESET$BRIGHT total tests\n"
elif [ "$total_solved" == "0" ]
then
	echo "Passed$RED $total_solved / $test_num$RESET$BRIGHT total tests\n"	
else
	echo "Passed$YELLOW $total_solved / $test_num$RESET$BRIGHT total tests\n"
fi

#### -- Cleanup -- ####
rm n-puzzle
