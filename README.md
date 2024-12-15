# Wardrobe Inventory Simulation

## How it Works

 1. Everyday at `SHOWER_TIMES`, change clothes and add used clothes to laundry 
    basket.
 2. Use exponetial distribution with `MEAN_LAUNDRY_PER_WEEK` to determine 
    laundry times. This has a hard limit of `MAX_LAUNDRY_PER_WEEK` and
    `MAX_LAUNDRY_PER_WEEK`.
 3. After a laundry is done, schedule next laundry at 5 am tomorrow.

The aim is to determine the minimum amount of clothes I need to have.

## Result

Using this configuration

```go
var SIMULATION_DURATION_HOUR = 40 * 365 * 24

var SHOWER_TIME = []int{6, 17}

const MEAN_LAUNDRY_DURATION_HOUR = 24
const MIN_LAUNDRY_DURATION_HOUR = MEAN_LAUNDRY_DURATION_HOUR / 2
const MAX_LAUNDRY_DURATION_HOUR = MEAN_LAUNDRY_DURATION_HOUR * 3
```

I got this result

```
WARDROBE INVENTORY SIMULATION (2024-12-15T09:49:08+07:00)

Simulated Time (Hours)  : 350400
Total Outfit            : 15
Logs Path               : logs/2024-12-15T09:49:08+07:00.json
```
