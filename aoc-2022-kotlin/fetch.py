import sys
import requests
import os
from decouple import config

year = sys.argv[1]
day = str(sys.argv[2]).zfill(2)
print(f'Fetching input for year={year}, day={day}')

path = f'src/main/kotlin/aoc{year}/day{day}'
cookies = {'session': f'{config("COOKIE")}'}
response = requests.get(f'https://adventofcode.com/{year}/day/{sys.argv[2]}/input', cookies=cookies)
if response.status_code != 200:
    print(f'AoC response_code: {response.status_code}')
else:
    for i in range(1, 3):
        if not os.path.exists(f"{path}/part{i}"):
            os.makedirs(f"{path}/part{i}")
        with open(f"{path}/part{i}/input.kt", 'w') as f:
            f.write(f"package aoc{year}.day{day}.part{i}\n\n")
            f.write('const val INPUT = """')
            f.write(response.text)
            f.write('"""')
