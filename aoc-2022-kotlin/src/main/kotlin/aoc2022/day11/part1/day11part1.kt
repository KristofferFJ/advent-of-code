package aoc2022.day11.part1

import utils.CollectionUtils.Companion.prod

fun main() {
    IntRange(1, 20).forEach {
        monkeys.forEach {
            it.throwItems()
        }
    }
    println(monkeys.sortedBy { it.inspections }.reversed().subList(0, 2).map { it.inspections }.prod())
}

data class Monkey(
    val startingItems: MutableList<Int>,
    val operation: (item: Int) -> Int,
    val test: (item: Int) -> Boolean,
    val trueMonkey: Int,
    val falseMonkey: Int,
    var inspections: Int = 0
) {
    fun throwItems() {
        this.startingItems.forEach {
            inspections += 1
            val worryLevel: Int = operation(it) / 3
            if (test(worryLevel)) {
                monkeys[trueMonkey].startingItems.add(worryLevel)
            } else {
                monkeys[falseMonkey].startingItems.add(worryLevel)
            }
        }
        startingItems.clear()
    }
}

val monkeys = mutableListOf(
    Monkey(
        mutableListOf(85, 79, 63, 72),
        { it * 17 },
        { it % 2 == 0 },
        2, 6
    ),
    Monkey(
        mutableListOf(53, 94, 65, 81, 93, 73, 57, 92),
        { it * it },
        { it % 7 == 0 },
        0, 2
    ),
    Monkey(
        mutableListOf(62, 63),
        { it + 7 },
        { it % 13 == 0 },
        7, 6
    ),
    Monkey(
        mutableListOf(57, 92, 56),
        { it + 4 },
        { it % 5 == 0 },
        4, 5
    ),
    Monkey(
        mutableListOf(67),
        { it + 5 },
        { it % 3 == 0 },
        1, 5
    ),
    Monkey(
        mutableListOf(85, 56, 66, 72, 57, 99),
        { it + 6 },
        { it % 19 == 0 },
        1, 0
    ),
    Monkey(
        mutableListOf(86, 65, 98, 97, 69),
        { it * 13 },
        { it % 11 == 0 },
        3, 7
    ),
    Monkey(
        mutableListOf(87, 68, 92, 66, 91, 50, 68),
        { it + 2 },
        { it % 17 == 0 },
        4, 3
    ),
)

val testmonkeys = mutableListOf(
    Monkey(
        mutableListOf(79, 98),
        { it * 19 },
        { it % 23 == 0 },
        2, 3
    ),
    Monkey(
        mutableListOf(54, 65, 75, 74),
        { it + 6 },
        { it % 19 == 0 },
        2, 0
    ),
    Monkey(
        mutableListOf(79, 60, 97),
        { it * it },
        { it % 13 == 0 },
        1, 3
    ),
    Monkey(
        mutableListOf(74),
        { it + 3 },
        { it % 17 == 0 },
        0, 1
    ),
)