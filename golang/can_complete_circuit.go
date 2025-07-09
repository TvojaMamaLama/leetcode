package main

func canCompleteCircuit(gas []int, cost []int) int {
	start, total, tank := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		gain := gas[i] - cost[i]
		total += gain
		tank += gain

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
