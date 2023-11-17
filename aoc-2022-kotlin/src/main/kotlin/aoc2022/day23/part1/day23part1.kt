package aoc2022.day23.part1

import utils.Field
import utils.Grid
import utils.Point

val grid = Grid(INPUT)
val directions = listOf("N", "S", "W", "E")

data class Proposition(val from: Field, val to: Point)

fun main() {
    IntRange(0, 9).forEach { index ->
        val propositions = grid.getFields().filter { it.value == "#" }.mapNotNull {
            makeProposition(it, index)
        }
        val nonDuplicatePositions = propositions.groupBy { it.to }.filter { it.value.size == 1 }.map { it.value[0] }
        nonDuplicatePositions.forEach {
            grid.set(it.to.x, it.to.y, "#")
            it.from.value = "."
        }
        println(grid)
    }

    val maxX = grid.getFields().filter { it.value == "#" }.maxOfOrNull { it.x }!!
    val maxY = grid.getFields().filter { it.value == "#" }.maxOfOrNull { it.y }!!
    val minX = grid.getFields().filter { it.value == "#" }.minOfOrNull { it.x }!!
    val minY = grid.getFields().filter { it.value == "#" }.minOfOrNull { it.y }!!

    println((maxX - minX + 1) * (maxY - minY + 1) - grid.getFields().filter { it.value == "#" }.size)
}

fun makeProposition(field: Field, iteration: Int): Proposition? {
    val x = field.x
    val y = field.y
    fun checkNorth(): Boolean {
        return grid.get(x - 1, y - 1).value == "." && grid.get(x, y - 1).value == "."
                && grid.get(x + 1, y - 1).value == "."
    }

    fun checkSouth(): Boolean {
        return grid.get(x - 1, y + 1).value == "." && grid.get(x, y + 1).value == "."
                && grid.get(x + 1, y + 1).value == "."
    }

    fun checkWest(): Boolean {
        return grid.get(x - 1, y - 1).value == "." && grid.get(x - 1, y).value == "."
                && grid.get(x - 1, y + 1).value == "."
    }

    fun checkEast(): Boolean {
        return grid.get(x + 1, y - 1).value == "." && grid.get(x + 1, y).value == "."
                && grid.get(x + 1, y + 1).value == "."
    }

    if (checkNorth() && checkSouth() && checkWest() && checkEast()) return null

    IntRange(0, 3).forEach {
        when (directions[(it + iteration) % 4]) {
            "N" -> if (checkNorth()) return Proposition(field, Point(x, y - 1))
            "S" -> if (checkSouth()) return Proposition(field, Point(x, y + 1))
            "W" -> if (checkWest()) return Proposition(field, Point(x - 1, y))
            "E" -> if (checkEast()) return Proposition(field, Point(x + 1, y))
        }
    }

    return null
}
