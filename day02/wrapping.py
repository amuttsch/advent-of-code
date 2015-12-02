import re
import heapq
from functools import reduce

def get_line():
    with open('input.txt', 'r') as file:
        for line in file:
            yield list(map(lambda x: int(x), re.split('x', line)))
    
total_wrapping_paper = 0        
total_ribbon = 0
for box in get_line():
    surface = [box[0]*box[1], box[1]*box[2], box[2]*box[0]]
    total_wrapping_paper += min(surface) # Add slack
    total_wrapping_paper += sum(map(lambda x: x*2, surface))
    total_ribbon += sum(map(lambda x: x+x, heapq.nsmallest(2, box))) # Perimeter
    total_ribbon += reduce(lambda x,y: x*y, box) # bow
    
print('Total wrapping paper:', total_wrapping_paper)
print('Total ribbon:', total_ribbon)
    