package aoc2022.day24.part2

import utils.CollectionUtils.Companion.isIn
import utils.CollectionUtils.Companion.removeDuplicates
import utils.Grid
import utils.Point

private val grid = Grid(INPUT)
private val length = grid.rows[0].size - 1
private val height = grid.rows.size - 1
private val startPoint = Point(1, 0)
private val endPoint = Point(length - 1, height)

private data class Moves(
    val moves: Int = 0,
    val position: Point = startPoint,
    var atEnd: Boolean = false,
    var atStart: Boolean = false
)

private val movesList = mutableListOf(Moves())
private fun main() {
    grid.getFields().filter { it.value.isIn(blizzardSymbols) }.forEach { it.value = "." }
    while (true) {
        val currentMoves = movesList.toList()
        moveBlizzards()
        currentMoves.forEach {
            it.doRound()
            movesList.removeDuplicates({ listOf(it.atStart, it.atEnd, it.position) })
        }

        val atEndpoint = movesList.filter { it.position == endPoint }
        if (atEndpoint.isNotEmpty()) {
            atEndpoint.forEach { it.atEnd = true }
            movesList.removeAll { !it.atEnd }
        }
        val atStart = movesList.filter { it.position == startPoint && it.atEnd }
        if (atStart.isNotEmpty()) {
            atStart.forEach { it.atStart = true }
            movesList.removeAll { !it.atStart }
        }
        if (movesList.any { it.atStart && it.position == endPoint }) {
            println(movesList.first { it.atEnd && it.position == startPoint })
            println(movesList.first { it.atEnd && it.position == startPoint }.moves)
            break
        }
    }

}

private fun Moves.doRound() {
    if (position.addX(1).isClear()) movesList.add(this.copy(moves = this.moves + 1, position = position.addX(1)))
    if (position.addX(-1).isClear()) movesList.add(this.copy(moves = this.moves + 1, position = position.addX(-1)))
    if (position.addY(1).isClear()) movesList.add(this.copy(moves = this.moves + 1, position = position.addY(1)))
    if (position.addY(-1).isClear()) movesList.add(this.copy(moves = this.moves + 1, position = position.addY(-1)))
    if (position.isClear()) movesList.add(this.copy(moves = this.moves + 1))
    movesList.remove(this)
}

private fun Point.isClear(): Boolean {
    if (this.y == -1) return false
    if (this.y == height + 1) return false
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
