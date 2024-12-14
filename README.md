# Wardrobe Inventory Simulation

## How it Works

```
PARAMETERS

SHOWER_TIMES
MEAN_LAUNDRY_PER_WEEK
STD_LAUNDRY_PER_WEEK
MEAN_LAUNDRY_TIME_HOUR
STD_LAUNDRY_TIME_HOUR
```

 1. Everyday at `SHOWER_TIMES`, change clothes and add used clothes to laundry 
    basket.
 2. Use normal distribution with `MEAN_LAUNDRY_PER_WEEK` to determine laundry times.
 3. Clothes are done laundrying `MEAN_LAUNDRY_TIME_HOUR` hours after laundry.

The aim is to determine the minimum amount of clothes I need to have.

## Result

Using this configuration

```go
var SIMULATION_DURATION_HOUR = 10 * 365 * 24

var SHOWER_TIME = []int{6, 17}

const MEAN_LAUNDRY_DURATION_HOUR = 24
const STD_LAUNDRY_DURATION_HOUR = 3

const MEAN_HOUR_UNTIL_NEXT_LAUNDRY = 2
const STD_HOUR_UNTIL_NEXT_LAUNDRY = 1
```

I got this result

```
WARDROBE INVENTORY SIMULATION (2024-12-14T22:49:20+07:00)

Simulated Time (Hours)  : 87600
Total Outfit            : 6
```
