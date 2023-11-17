import read_functions as read

list_of_strings = read.read_as_strings("input.txt")
split_into_elves = [elf.split(",") for elf in ",".join(list_of_strings).split(",,")]
elf_sums = [sum([int(food) for food in elf]) for elf in split_into_elves]

# del 1
print(max(elf_sums))

# del 2
elf_sums.sort(reverse=True)
print(sum(elf_sums[0:3])
      )
