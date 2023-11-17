package aoc2020.day02

fun main() {
    fun partOne() {
        println(
            input.filter {
                val (from, to, letter) = it[0].split("-| ".toRegex())
                it[1].count { it.toString() == letter } >= from.toInt() && it[1].count { it.toString() == letter } <= to.toInt()
            }.size
        )
    }

    fun partTwo() {
        println(
            input.filter {
                val (first, second, letter) = it[0].split("-| ".toRegex())
                (it[1][first.toInt() - 1].toString() == letter) != (it[1][second.toInt() - 1].toString() == letter)
            }.size
        )
    }
    partOne()
    partTwo()
}

private val input = """""".trimIndent()
    .split("\n")
    .map { it.split(": ") }
