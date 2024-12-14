# Wardrobe Inventory Simulation

## How it Works

```
PARAMETERS

SHOWER_TIMES = [6, 17]
MEAN_LAUNDRY_PER_WEEK = 5
MEAN_LAUNDRY_TIME_HOUR = 24
```

 1. Everyday at 0, determine whether:
     a. Day is cold or not.
 2. Everyday at `SHOWER_TIMES`, change clothes and add used clothes to laundry 
    basket.
 3. Use exponential with `MEAN_LAUNDRY_PER_WEEK` to determine laundry times.
 4. Clothes are done laundrying `MEAN_LAUNDRY_TIME_HOUR` hours after laundry.

The aim is to determine the minimum amount of clothes I need to have.
