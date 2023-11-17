package aoc2022.day01

fun getSums(): List<Int> {
    val groups: List<List<Int>> = INPUT.split("\n\n").map { it.split("\n").map { it.toInt() } }
    return groups.map { it.sum() }.sorted().reversed()
}

fun partOne() {
    println(getSums()[0])
}

fun partTwo() {
    println(getSums().subList(0, 3).sum())
}

fun main() {
    partOne()
    partTwo()
}
