#### -- Config -- ####

min_size=3
max_size=5
test_cases=10
unsolvable_test=1
solvable_test=1
unit_test=1
random_test=1


#### -- Print Header -- ####
start=`date +%s`
printf "\E[H\E[2J"
echo "\x1b[1mLaunching N-Puzzle Performance Test\x1B[0m\n"
## echo "Usage: '''go build''' to build the binary 'n-puzzle'. then ./performance_test.sh"

echo "\t\x1b[4m-- Config --\x1b[0m"
echo "Minimum size: \t\t$min_size"
echo "Maximum size: \t\t$max_size"
echo "Number of test cases: \t$test_cases"
if [ "$unsolvable_test" != 0 ]
then
	echo "Unsolvable Tests: \t\x1b[32mon\x1b[0m"
else	
	echo "Unsolvable Tests: \t\x1b[31moff\x1b[0m"
fi
if [ "$solvable_test" != 0 ]
then
	echo "Solvable Tests: \t\x1b[32mon\x1b[0m"
else	
	echo "Solvable Tests: \t\x1b[31moff\x1b[0m"
fi
if [ "$unit_test" != 0 ]
then
	echo "Unit Tests: \t\t\x1b[32mon\x1b[0m"
else	
	echo "Unit Tests: \t\t\x1b[31moff\x1b[0m"
fi
if [ "$random_test" != 0 ]
then
	echo "Random Tests: \t\t\x1b[32mon\x1b[0m\n"
else	
	echo "Random Tests: \t\t\x1b[31moff\x1b[0m\n"
fi

#### -- Test -- ####
size=$min_size
test_num=0
if [ -f "rm_me.txt" ]
then
	$(rm rm_me.txt)
fi
while [ $size -lt $(expr $max_size + 1) ]
do
	echo "\x1b[1mSize - $size\x1B[0m"

#### -- Unsolvable Unit Tests -- ####
	if [ "$unit_test" != 0 -a "$unsolvable_test" != 0 -a "$size" -gt 2 -a "$size" -lt 10 ]
	then
		if [ "$test_cases" -lt 10 ]
		then
			case=$test_cases
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
			unit=$(echo "Boards/Unsolvable/$size/$size""u$count.txt")
			output=$(../n-puzzle $unit)
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

			unsolvable=$(echo "$output" | tail -n -2 | head -n 1)
			if [ "$unsolvable" = "This puzzle is unsolvable." ]
			then
				u=$(($u + 1))
				echo "\x1b[32m.\x1b[0m\c"
			else	
				echo "\x1b[31m.\x1b[0m\c"
			fi
			test_num=$(($test_num + 1))
		done
		if [ "$u" != 0 ]
		then
			mean=$(echo "scale = 9; $tcumulative / $u" | bc)
		else
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

		if [ "$solved" = 0 ]
		then
			echo "\x1b[31m"
		elif [ "$u" -lt "$count" ]
		then
			echo "\x1b[33m"
		else
			echo "\x1b[32m"
		fi
		echo "Unsolvable unit tests correctly identified: \t$u/$count\x1b[0m"
		echo "Solve time in seconds:\t\t\tMean: \t$mean"
		echo "\t\t\t\t\tWorst: \t$worst"
		echo "\t\t\t\t\tBest: \t$best"
	fi

#### -- Unsolvable Random Boards -- ####
	if [ "$random_test" != 0 -a "$unsolvable_test" != 0 ]
	then
		case=$test_cases
		u=0
		count=0
		best=42
		worst=0
		tcumulative=0
		count=0
		while [ $count -lt $case ]
		do
			count=$(($count + 1))
			output=$(python generator.py -u $size >> rm_me.txt; ../n-puzzle rm_me.txt)
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

			unsolvable=$(echo "$output" | tail -n -2 | head -n 1)
			if [ "$unsolvable" = "This puzzle is unsolvable." ]
			then
				u=$(($u + 1))
				echo "\x1b[32m.\x1b[0m\c"
			else	
				echo "\x1b[31m.\x1b[0m\c"
			fi
			test_num=$(($test_num + 1))
			$(rm rm_me.txt)
		done
		if [ "$u" != 0 ]
		then
			mean=$(echo "scale = 9; $tcumulative / $u" | bc)
		else
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

		if [ "$solved" = 0 ]
		then
			echo "\x1b[31m"
		elif [ "$u" -lt "$count" ]
		then
			echo "\x1b[33m"
		else
			echo "\x1b[32m"
		fi
		echo "Unsolvable random tests correctly identified: \t$u/$count\x1b[0m"
		echo "Solve time in seconds:\t\t\tMean: \t$mean"
		echo "\t\t\t\t\tWorst: \t$worst"
		echo "\t\t\t\t\tBest: \t$best"
	fi

#### -- Solvable Unit Tests -- ####
	if [ "$unit_test" != 0 -a "$solvable_test" != 0 -a "$size" -gt 2 -a "$size" -lt 10 ]
	then
		if [ "$test_cases" -lt 10 ]
		then
			case=$test_cases
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
			solvable=$(../n-puzzle $unit)
			end=$(echo "$solvable" | tail -n -1)
			if [ "$end" != "You've finished n-puzzle!" ]
			then
				echo "\x1b[31m.\x1b[0m\c"
				continue
			else
				solved=$(($solved + 1))
				echo "\x1b[32m.\x1b[0m\c"
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
				echo "$time" ##########################
				minute=$(echo "$time" | rev | cut -d "." -f 2 | cut -c-2-2)
				if [ "$minute" = "m" ]
				then
					echo "$time" ##########################
					minutes=$(echo "$time" | cut -d "m" -f 1)
					seconds=$(echo "$time" | rev | cat -d "m" -f 1 | cut -c-2-42 | rev)
					time=$(echo "scale = 9; ($minutes * 60) + $seconds" | bc | cut -d "." -f 1)
					echo "minutes calc: $time" ##################
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
			fi
		done
		if [ "$solved" != 0 ]
		then
			mean=$(echo "scale = 9; $tcumulative / $solved" | bc)
		else
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

		if [ "$solved" = 0 ]
		then
			echo "\x1b[31m"
		elif [ "$solved" -lt "$count" ]
		then
			echo "\x1b[33m"
		else
			echo "\x1b[32m"
		fi
		echo "Solvable unit tests correctly solved: \t\t$solved/$count\x1b[0m"
		echo "Solve time in seconds:\t\t\tMean: \t$mean"
		echo "\t\t\t\t\tWorst: \t$worst"
		echo "\t\t\t\t\tBest: \t$best"
	fi

#### -- Solvable Random Boards -- ####
	if [ "$random_test" != 0 -a "$solvable_test" != 0 ]
	then
		case=$test_cases
		solved=0
		best=42
		worst=0
		tcumulative=0
		count=0
		while [ $count -lt $case ]
		do
			count=$(($count + 1))
			test_num=$(($test_num + 1))
			solvable=$(python generator.py -s $size >> rm_me.txt; ../n-puzzle rm_me.txt)
			end=$(echo "$solvable" | tail -n -1)
			if [ "$end" != "You've finished n-puzzle!" ]
			then
				$(rm rm_me.txt)
				echo "\x1b[31m.\x1b[0m\c"
				continue
			else
				solved=$(($solved + 1))
				echo "\x1b[32m.\x1b[0m\c"
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
		if [ "$solved" != 0 ]
		then
			mean=$(echo "scale = 9; $tcumulative / $solved" | bc)
		else
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

		if [ "$solved" = 0 ]
		then
			echo "\x1b[31m"
		elif [ "$solved" -lt "$count" ]
		then
			echo "\x1b[33m"
		else
			echo "\x1b[32m"
		fi
		echo "Solvable random tests correctly solved: \t$solved/$count\x1b[0m"
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
