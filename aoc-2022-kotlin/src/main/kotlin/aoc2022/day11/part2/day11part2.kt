package aoc2022.day11.part2

import utils.CollectionUtils.Companion.prod

const val testMonkeyDivisor = 23 * 19 * 13 * 17
const val monkeyDivisor = 2 * 7 * 13 * 5 * 3 * 19 * 11 * 17

fun main() {
    IntRange(1, 10000).forEach {
        monkeys.forEach {
            it.throwItems()
        }
    }
    println(monkeys.sortedBy { it.inspections }.reversed().subList(0, 2).map { it.inspections }.prod())
}

data class Monkey(
    val startingItems: MutableList<Long>,
    val operation: (item: Long) -> Long,
    val test: (item: Long) -> Boolean,
    val trueMonkey: Int,
    val falseMonkey: Int,
    var inspections: Long = 0
) {
    fun throwItems() {
        this.startingItems.forEach {
            inspections += 1
            val worryLevel: Long = operation(it) % monkeyDivisor
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
        { it.mod(2) == 0 },
        2, 6
    ),
    Monkey(
        mutableListOf(53, 94, 65, 81, 93, 73, 57, 92),
        { it * it },
        { it.mod(7) == 0 },
        0, 2
    ),
    Monkey(
        mutableListOf(62, 63),
        { it + 7 },
        { it.mod(13) == 0 },
        7, 6
    ),
    Monkey(
        mutableListOf(57, 92, 56),
        { it + 4 },
        { it.mod(5) == 0 },
        4, 5
    ),
    Monkey(
        mutableListOf(67),
        { it + 5 },
        { it.mod(3) == 0 },
        1, 5
    ),
    Monkey(
        mutableListOf(85, 56, 66, 72, 57, 99),
        { it + 6 },
        { it.mod(19) == 0 },
        1, 0
    ),
    Monkey(
        mutableListOf(86, 65, 98, 97, 69),
        { it * 13 },
        { it.mod(11) == 0 },
        3, 7
    ),
    Monkey(
        mutableListOf(87, 68, 92, 66, 91, 50, 68),
        { it + 2 },
        { it.mod(17) == 0 },
        4, 3
    ),
)

val testmonkeys = mutableListOf(
    Monkey(
        mutableListOf(79, 98),
        { it * 19 },
        { it.mod(23) == 0 },
        2, 3
    ),
    Monkey(
        mutableListOf(54, 65, 75, 74),
        { it + 6 },
        { it.mod(19) == 0 },
        2, 0
    ),
    Monkey(
        mutableListOf(79, 60, 97),
        { it * it },
        { it.mod(13) == 0 },
        1, 3
    ),
    Monkey(
        mutableListOf(74),
        { it + 3 },
        { it.mod(17) == 0 },
        0, 1
    ),
)