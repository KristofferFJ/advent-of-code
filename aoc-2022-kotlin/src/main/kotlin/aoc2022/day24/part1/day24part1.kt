package aoc2022.day24.part1

import utils.CollectionUtils.Companion.addReturn
import utils.CollectionUtils.Companion.isIn
import utils.CollectionUtils.Companion.removeDuplicates
import utils.Grid
import utils.Point


private val grid = Grid(INPUT)
private val length = grid.rows[0].size - 1
private val height = grid.rows.size - 1
private val endPoint = Point(length - 1, height)
private val moves = mutableListOf(mutableListOf(Point(1, 0)))
private fun main() {
    grid.getFields().filter { it.value.isIn(blizzardSymbols) }.forEach { it.value = "." }
    var finished = false
    while (!finished) {
        val currentMoves = moves.toList()
        moveBlizzards()
        currentMoves.forEach {
            if (it.doRound()) finished = true
            moves.removeAll { it.size < moves.maxBy { it.size }.size }
            moves.removeDuplicates({ it.last() })
        }
    }
    println(moves.filter { it.last() == endPoint })
    println(moves.first { it.last() == endPoint }.size - 1)
}

private fun List<Point>.doRound(): Boolean {
    val position = this.last()
    if (position.addX(1).isClear()) moves.add(this.toMutableList().addReturn(position.addX(1)))
    if (position.addX(-1).isClear()) moves.add(this.toMutableList().addReturn(position.addX(-1)))
    if (position.addY(1).isClear()) moves.add(this.toMutableList().addReturn(position.addY(1)))
    if (position.addY(-1).isClear()) moves.add(this.toMutableList().addReturn(position.addY(-1)))
    if (position.isClear()) moves.add(this.toMutableList().addReturn(position))
    moves.remove(this)

    return moves.any { it.last() == endPoint }
}

private fun Point.isClear(): Boolean {
    if (this.y == -1) return false
    return grid.at(this) == "." && blizzards.none { it.x == this.x && it.y == this.y }
}

private data class Blizzard(var x: Int, var y: Int, var value: String)

private val blizzardSymbols = listOf("<", ">", "^", "v")
private val blizzards = grid.getFields().filter { it.value.isIn(blizzardSymbols) }.map { Blizzard(it.x, it.y, it.value) }
private fun moveBlizzards() {
    blizzards.forEach { blizzard ->
        when (blizzard.value) {
            "<" -> if (blizzard.x == 1) blizzard.x = length - 1 else blizzard.x -= 1
            ">" -> if (blizzard.x == length - 1) blizzard.x = 1 else blizzard.x += 1
            "v" -> if (blizzard.y == height - 1) blizzard.y = 1 else blizzard.y += 1
            "^" -> if (blizzard.y == 1) blizzard.y = height - 1 else blizzard.y -= 1
        }
    }
}
