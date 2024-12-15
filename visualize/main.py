import matplotlib.pyplot as plt
import numpy as np
import json
from typing import Dict, List, Tuple

# Define constants for color mapping
COLOR_CLEAN = '#00A9E0'
COLOR_DIRTY = '#7A4B3B'
COLOR_BEING_CLEANED = '#A4D7D3'

def load_data(file_path: str) -> List[Dict]:
    """ Load the data from a JSON file and return it. """
    try:
        with open(file_path, "r") as f:
            return json.load(f)
    except (FileNotFoundError, json.JSONDecodeError) as e:
        print(f"Error loading file {file_path}: {e}")
        return []

def extract_data(data: List[Dict]) -> Tuple[List[int], List[int], List[int], List[int]]:
    """ Extract the hours and states from the data. """
    hours = [item["Hour"] for item in data]
    state_0 = [item["NumOfOutfitsByState"].get("0", 0) for item in data]
    state_1 = [item["NumOfOutfitsByState"].get("1", 0) for item in data]
    state_2 = [item["NumOfOutfitsByState"].get("2", 0) for item in data]
    return hours, state_0, state_1, state_2

def plot_data(hours: List[int], state_0: List[int], state_1: List[int], state_2: List[int], units_per_x: int = 4) -> None:
    """ Create the stacked line chart with the provided data and save it. """
    # Compute cumulative sums for stacked line
    stacked_state_1 = np.add(state_0, state_1)
    stacked_state_2 = np.add(stacked_state_1, state_2)

    # Plot the stacked line chart with filled areas
    fig, ax = plt.subplots()

    # Adjust the figure width based on the number of hours
    fig_width = len(hours) / units_per_x
    fig.set_figwidth(fig_width)

    ax.set_xlim(left=0, right=hours[-1])
    ax.set_ylim(bottom=0, top=stacked_state_2[-1])

    # Fill areas under the lines
    ax.fill_between(hours, 0, state_0, color=COLOR_CLEAN, label='Clean', alpha=0.3)
    ax.fill_between(hours, 0, state_1, color=COLOR_DIRTY, label='Dirty', alpha=0.3)
    ax.fill_between(hours, 0, state_2, color=COLOR_BEING_CLEANED, label='Being Cleaned', alpha=0.3)

    # Adding labels and title
    ax.set_xlabel('Hour')
    ax.set_ylabel('Number of Outfits')
    ax.set_title('Outfits by State Over Time')
    ax.legend()

    # Save the plot to a file
    plt.savefig('outfits_by_state_over_time.png', bbox_inches='tight')

    # Close the plot to avoid displaying it
    plt.close()

def main(file_path: str) -> None:
    """ Main function to load data, extract and plot. """
    data = load_data(file_path)
    if data:
        hours, state_0, state_1, state_2 = extract_data(data)
        plot_data(hours, state_0, state_1, state_2)

# Run the script with the given file path
if __name__ == "__main__":
    file_path = "../logs/2024-12-15T10:11:06+07:00.json"
    main(file_path)
