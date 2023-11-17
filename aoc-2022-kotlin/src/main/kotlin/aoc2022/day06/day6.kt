package aoc2022.day06

fun main() {
    fun partOne() {
        INPUT.forEachIndexed { index, _ -> if (INPUT.substring(index, index.plus(4)).toSet().size == 4) { println(index + 4); return }}
    }

    fun partTwo() {
        INPUT.forEachIndexed { index, _ -> if (INPUT.substring(index, index.plus(14)).toSet().size == 14) { println(index + 14); return }}
    }

    partOne()
    partTwo()
}
