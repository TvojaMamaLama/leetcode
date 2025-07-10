package main

func canCompleteCircuit1(gas []int, cost []int) int {
	total, tank, start := 0, 0, 0

	for i := 0; i < len(cost); i++ {
		profit := gas[i] - cost[i]
		total += profit
		tank += profit

		if tank < 0 {
			start = i + 1
			tank = 0
		}

	}

	if total < 0 {
		return -1
	}

	return start
}
