#### -- Print Header -- ####
RESET="\x1b[0m"
BRIGHT="\x1b[1m"
RED="\x1b[31m"
GREEN="\x1b[32m"

printf "\E[H\E[2J" ## Clear screen
printf $BRIGHT
echo "Launching N-Puzzle Performance Test$RESET\n"

start=`date +%s`
## echo "Usage: '''go build''' to build the binary 'n-puzzle'. then ./performance_test.sh"


#### -- Config -- ####

MIN_SIZE=3
MAX_SIZE=3
TEST_CASES=5
UNSOLVABLE_TEST=1
SOLVABLE_TEST=1
UNIT_TEST=1
RANDOM_TEST=1
HEURISTIC="manhattan"


#### -- Print Config -- ####
echo "\t\x1b[4m-- Config --$RESET"
echo "Minimum size: \t\t $MIN_SIZE"
echo "Maximum size: \t\t $MAX_SIZE"
echo "Number of test cases: \t $TEST_CASES"
if [ "$UNSOLVABLE_TEST" != 0 ]
then
	echo "Unsolvable Tests: \t$GREEN on$RESET"
else	
	echo "Unsolvable Tests: \t$RED off$RESET"
fi
if [ "$SOLVABLE_TEST" != 0 ]
then
	echo "Solvable Tests: \t$GREEN on$RESET"
else	
	echo "Solvable Tests: \t$RED off$RESET"
fi
if [ "$UNIT_TEST" != 0 ]
then
	echo "Unit Tests: \t\t$GREEN on$RESET"
else	
	echo "Unit Tests: \t\t$RED off$RESET"
fi
if [ "$RANDOM_TEST" != 0 ]
then
	echo "Random Tests: \t\t$GREEN on$RESET"
else	
	echo "Random Tests: \t\t$RED off$RESET"
fi
echo "Heuristic: \t\t $HEURISTIC\n"



#### -- Test Function -- ####
# unit_test()
# {
# 	echo hi!
# }

	# 	if [ "$TEST_CASES" -lt 10 ]
	# 	then
	# 		case=$TEST_CASES
	# 	else
	# 		case=10
	# 	fi
	# 	u=0
	# 	count=0
	# 	best=42
	# 	worst=0
	# 	tcumulative=0
	# 	count=0
	# 	while [ $count -lt $case ]
	# 	do
	# 		count=$(($count + 1))
	# 		test_num=$(($test_num + 1))
	# 		unit=$(echo "Boards/Unsolvable/$size/$size""u$count.txt")
	# 		output=$(../n-puzzle -h=$HEURISTIC $unit)
	# 		unsolvable=$(echo "$output" | tail -n -2 | head -n 1)
	# 		if [ "$unsolvable" = "This puzzle is unsolvable." ]
	# 		then
	# 			u=$(($u + 1))
	# 			echo "$GREEN.$RESET\c"
	# 		else	
	# 			echo "$RED.$RESET\c"
	# 			continue
	# 		fi

	# 		time=$(echo "$output" | tail -n -1 | cut -d " " -f 3)
	# 		prefix=$(echo "$time" | rev | cut -c-1-2 | rev | cut -c-1-1)
	# 		if [ "$prefix" = "m" ]
	# 		then
	# 			time=$(echo "$time" | rev | cut -c3-42 | rev)
	# 			time=$(echo "scale = 9; ($time / 1000)" | bc)	
	# 			tcumulative=$(echo "scale = 9; $tcumulative + $time" | bc)
	# 			time_up=$(echo "scale = 0; $time * 1000000000" | bc | cut -d "." -f 1)
	# 			worst_up=$(echo "scale = 0; $worst * 1000000000" | bc | cut -d "." -f 1)
	# 			best_up=$(echo "scale = 0; $best * 1000000000" | bc | cut -d "." -f 1)
	# 			if [ "$time_up" -gt "$worst_up" ]
	# 			then
	# 				worst=$time
	# 			fi
	# 			if [ "$time_up" -lt "$best_up" ]
	# 			then
	# 				best=$time
	# 			fi
	# 		elif [ "$prefix" = "µ" ]
	# 		then
	# 			time=$(echo "$time" | rev | cut -c3-42 | rev)
	# 			time=$(echo "scale = 9; ($time / 1000000)" | bc)	
	# 			tcumulative=$(echo "scale = 9; $tcumulative + $time" | bc)	
	# 			time_up=$(echo "scale = 0; $time * 1000000000" | bc | cut -d "." -f 1)
	# 			worst_up=$(echo "scale = 0; $worst * 1000000000" | bc | cut -d "." -f 1)
	# 			best_up=$(echo "scale = 0; $best * 1000000000" | bc | cut -d "." -f 1)
	# 			if [ "$time_up" -gt "$worst_up" ]
	# 			then
	# 				worst=$time
	# 			fi
	# 			if [ "$time_up" -lt "$best_up" ]
	# 			then
	# 				best=$time
	# 			fi
	# 		else
	# 			time=$(echo "$time" | rev | cut -c2-42 | rev)
	# 			tcumulative=$(echo "$tcumulative + $time" | bc)
	# 			time_up=$(echo "scale = 0; $time * 1000000000" | bc | cut -d "." -f 1)
	# 			worst_up=$(echo "scale = 0; $worst * 1000000000" | bc | cut -d "." -f 1)
	# 			best_up=$(echo "scale = 0; $best * 1000000000" | bc | cut -d "." -f 1)
	# 			if [ "$time_up" -gt "$worst_up" ]
	# 			then
	# 				worst=$time
	# 			fi
	# 			if [ "$time_up" -lt "$best_up" ]
	# 			then
	# 				best=$time
	# 			fi
	# 		fi
	# 	done
	# 	if [ "$u" != 0 ]
	# 	then
	# 		mean=$(echo "scale = 9; $tcumulative / $u" | bc)
	# 	else
	# 		mean="$RED Failed$RESET"
	# 	fi
	# 	if [ "$worst" = 0 ]
	# 	then
	# 		worst="$RED Failed$RESET"
	# 	fi
	# 	if [ "$best" = 42 ]
	# 	then
	# 		best="$RED Failed$RESET"
	# 	fi

	# 	if [ "$solved" = 0 ]
	# 	then
	# 		echo "$RED"
	# 	elif [ "$u" -lt "$count" ]
	# 	then
	# 		echo "\x1b[33m"
	# 	else
	# 		echo "$GREEN"
	# 	fi
	# 	echo "Unsolvable unit tests correctly identified: \t$u/$count$RESET"
	# 	echo "Solve time in seconds:\t\t\tMean: \t$mean"
	# 	echo "\t\t\t\t\tWorst: \t$worst"
	# 	echo "\t\t\t\t\tBest: \t$best"
	# fi




# unit_test


#### -- Test -- ####
size=$MIN_SIZE
test_num=0
if [ -f "rm_me.txt" ]
then
	$(rm rm_me.txt)
fi
while [ $size -lt $(expr $MAX_SIZE + 1) ]
do
	echo $BRIGHT
	echo "Size - $size$RESET"

#### -- Unsolvable Unit Tests -- ####
	if [ "$UNIT_TEST" != 0 -a "$UNSOLVABLE_TEST" != 0 -a "$size" -gt 2 -a "$size" -lt 10 ]
	then
		if [ "$TEST_CASES" -lt 10 ]
		then
			case=$TEST_CASES
		else
			case=10
		fi
		u=0
		count=0
		best=42
		worst=0
		tcumulative=0
		count=0
		while [ $count -lt $case ]
		do
			count=$(($count + 1))
			test_num=$(($test_num + 1))
			unit=$(echo "Boards/Unsolvable/$size/$size""u$count.txt")
			output=$(../n-puzzle -h=$HEURISTIC $unit)
			unsolvable=$(echo "$output" | tail -n -2 | head -n 1)
			if [ "$unsolvable" = "This puzzle is unsolvable." ]
			then
				u=$(($u + 1))
				echo "$GREEN.$RESET\c"
			else	
				echo "$RED.$RESET\c"
				continue
			fi

			time=$(echo "$output" | tail -n -1 | cut -d " " -f 3)
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
		done
		if [ "$u" != 0 ]
		then
			mean=$(echo "scale = 9; $tcumulative / $u" | bc)
		else
			mean="$RED Failed$RESET"
		fi
		if [ "$worst" = 0 ]
		then
			worst="$RED Failed$RESET"
		fi
		if [ "$best" = 42 ]
		then
			best="$RED Failed$RESET"
		fi

		if [ "$solved" = 0 ]
		then
			echo "$RED"
		elif [ "$u" -lt "$count" ]
		then
			echo "\x1b[33m"
		else
			echo "$GREEN"
		fi
		echo "Unsolvable unit tests correctly identified: \t$u/$count$RESET"
		echo "Solve time in seconds:\t\t\tMean: \t$mean"
		echo "\t\t\t\t\tWorst: \t$worst"
		echo "\t\t\t\t\tBest: \t$best"
	fi

#### -- Unsolvable Random Boards -- ####
	if [ "$RANDOM_TEST" != 0 -a "$UNSOLVABLE_TEST" != 0 ]
	then
		case=$TEST_CASES
		u=0
		count=0
		best=42
		worst=0
		tcumulative=0
		count=0
		while [ $count -lt $case ]
		do
			count=$(($count + 1))
			test_num=$(($test_num + 1))
			output=$(python generator.py -u $size >> rm_me.txt; ../n-puzzle -h=$HEURISTIC rm_me.txt)
			unsolvable=$(echo "$output" | tail -n -2 | head -n 1)
			if [ "$unsolvable" = "This puzzle is unsolvable." ]
			then
				u=$(($u + 1))
				echo "$GREEN.$RESET\c"
			else	
				echo "$RED.$RESET\c"
				$(rm rm_me.txt)
				continue
			fi

			time=$(echo "$output" | tail -n -1 | cut -d " " -f 3)
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
			$(rm rm_me.txt)
		done
		if [ "$u" != 0 ]
		then
			mean=$(echo "scale = 9; $tcumulative / $u" | bc)
		else
			mean="$RED Failed$RESET"
		fi
		if [ "$worst" = 0 ]
		then
			worst="$RED Failed$RESET"
		fi
		if [ "$best" = 42 ]
		then
			best="$RED Failed$RESET"
		fi

		if [ "$solved" = 0 ]
		then
			echo "$RED"
		elif [ "$u" -lt "$count" ]
		then
			echo "\x1b[33m"
		else
			echo "$GREEN"
		fi
		echo "Unsolvable random tests correctly identified: \t$u/$count$RESET"
		echo "Solve time in seconds:\t\t\tMean: \t$mean"
		echo "\t\t\t\t\tWorst: \t$worst"
		echo "\t\t\t\t\tBest: \t$best"
	fi

#### -- Solvable Unit Tests -- ####
	if [ "$UNIT_TEST" != 0 -a "$SOLVABLE_TEST" != 0 -a "$size" -gt 2 -a "$size" -lt 10 ]
	then
		if [ "$TEST_CASES" -lt 10 ]
		then
			case=$TEST_CASES
		else
			case=10
		fi
		solved=0
		best=42
		worst=0
		tcumulative=0
		count=0
		while [ $count -lt $case ]
		do
			count=$(($count + 1))
			test_num=$(($test_num + 1))
			unit=$(echo "Boards/Solvable/$size/$size""s$count.txt")
			solvable=$(../n-puzzle -h=$HEURISTIC $unit)
			end=$(echo "$solvable" | tail -n -1)
			if [ "$end" != "You've finished n-puzzle!" ]
			then
				echo "$RED.$RESET\c"
				continue
			else
				solved=$(($solved + 1))
				echo "$GREEN.$RESET\c"
			fi
			time=$(echo "$solvable" | tail -n -2 | head -n 1 | cut -d " " -f 3)
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
				minute=$(echo "$time" | rev | cut -d "." -f 2 | cut -c 2-2)
				minute_alt=$(echo "$time" | rev | cut -d "." -f 2 | cut -c 3-3)
				if [ "$minute" = "m" -o "$minute_alt" = "m" ]
				then
					minutes=$(echo "$time" | cut -d "m" -f 1)
					seconds=$(echo "$time" | rev | cut -d "m" -f 1 | cut -c 2-42 | rev)
					time=$(echo "scale = 9; ($minutes * 60) + $seconds" | bc)
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
				else
					time=$(echo "$time" | rev | cut -c 2-42 | rev)
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
			fi
		done
		if [ "$solved" != 0 ]
		then
			mean=$(echo "scale = 9; $tcumulative / $solved" | bc)
		else
			mean="$RED Failed$RESET"
		fi
		if [ "$worst" = 0 ]
		then
			worst="$RED Failed$RESET"
		fi
		if [ "$best" = 42 ]
		then
			best="$RED Failed$RESET"
		fi

		if [ "$solved" = 0 ]
		then
			echo "$RED"
		elif [ "$solved" -lt "$count" ]
		then
			echo "\x1b[33m"
		else
			echo "$GREEN"
		fi
		echo "Solvable unit tests correctly solved: \t\t$solved/$count$RESET"
		echo "Solve time in seconds:\t\t\tMean: \t$mean"
		echo "\t\t\t\t\tWorst: \t$worst"
		echo "\t\t\t\t\tBest: \t$best"
	fi

#### -- Solvable Random Boards -- ####
	if [ "$RANDOM_TEST" != 0 -a "$SOLVABLE_TEST" != 0 ]
	then
		case=$TEST_CASES
		solved=0
		best=42
		worst=0
		tcumulative=0
		count=0
		while [ $count -lt $case ]
		do
			count=$(($count + 1))
			test_num=$(($test_num + 1))
			solvable=$(python generator.py -s $size >> rm_me.txt; ../n-puzzle -h=$HEURISTIC rm_me.txt)
			end=$(echo "$solvable" | tail -n -1)
			if [ "$end" != "You've finished n-puzzle!" ]
			then
				$(rm rm_me.txt)
				echo "$RED.$RESET\c"
				continue
			else
				solved=$(($solved + 1))
				echo "$GREEN.$RESET\c"
			fi
			time=$(echo "$solvable" | tail -n -2 | head -n 1 | cut -d " " -f 3)
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
				minute=$(echo "$time" | rev | cut -d "." -f 2 | cut -c 2-2)
				minute_alt=$(echo "$time" | rev | cut -d "." -f 2 | cut -c 3-3)
				if [ "$minute" = "m" -o "$minute_alt" = "m" ]
				then
					minutes=$(echo "$time" | cut -d "m" -f 1)
					seconds=$(echo "$time" | rev | cut -d "m" -f 1 | cut -c 2-42 | rev)
					time=$(echo "scale = 9; ($minutes * 60) + $seconds" | bc)
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
				else
					time=$(echo "$time" | rev | cut -c 2-42 | rev)
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
			fi
			$(rm rm_me.txt)
		done
		if [ "$solved" != 0 ]
		then
			mean=$(echo "scale = 9; $tcumulative / $solved" | bc)
		else
			mean="$RED Failed$RESET"
		fi
		if [ "$worst" = 0 ]
		then
			worst="$RED Failed$RESET"
		fi
		if [ "$best" = 42 ]
		then
			best="$RED Failed$RESET"
		fi

		if [ "$solved" = 0 ]
		then
			echo "$RED"
		elif [ "$solved" -lt "$count" ]
		then
			echo "\x1b[33m"
		else
			echo "$GREEN"
		fi
		echo "Solvable random tests correctly solved: \t$solved/$count$RESET"
		echo "Solve time in seconds:\t\t\tMean: \t$mean"
		echo "\t\t\t\t\tWorst: \t$worst"
		echo "\t\t\t\t\tBest: \t$best"
	fi
	echo " "
	size=$(($size + 1))
done
end=`date +%s`
runtime=$((end-start))
echo "N-Puzzle performance test finished, $test_num tests run in $runtime seconds."
echo "Have a nice day!"
