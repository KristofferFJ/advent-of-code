package aoc2022.day12.part1

import utils.LetterUtil.Companion.getLetterValue

fun main() {
    val grid: List<List<Char>> = INPUT.split("\n").map { it.toList() }
    val depth = grid.size - 1
    val length = grid.first().size - 1
    val distances: List<MutableList<Int>> = grid.map { it.map { (depth + 1) * (length + 1) }.toMutableList() }

    data class Position(val x: Int, val y: Int) {
        fun getHeight(): Int {
            val letter = grid[this.y][this.x]
            if (letter == 'S') {
                return 'a'.getLetterValue()
            }
            if (letter == 'E') {
                return 'z'.getLetterValue()
            }
            return letter.getLetterValue()
        }

        fun getDistance(): Int {
            return distances[this.y][this.x]
        }

        fun setDistance(distance: Int) {
            distances[this.y][this.x] = distance
        }
    }

    fun findEndpoint(): Position {
        grid.forEachIndexed { indexRow, chars ->
            chars.forEachIndexed { indexColumn, char ->
                if (char == 'E') return Position(indexColumn, indexRow)
            }
        }
        throw Exception("what")
    }

    fun findStart(): Position {
        grid.forEachIndexed { indexRow, chars ->
            chars.forEachIndexed { indexColumn, char ->
                if (char == 'S') return Position(indexColumn, indexRow)
            }
        }
        throw Exception("what")
    }

    fun getAdjacent(position: Position): List<Position> {
        val adjacentPositions = mutableListOf<Position>()
        if (position.x < length) {
            adjacentPositions.add(Position(position.x + 1, position.y))
        }
        if (position.x > 0) {
            adjacentPositions.add(Position(position.x - 1, position.y))
        }
        if (position.y > 0) {
            adjacentPositions.add(Position(position.x, position.y - 1))
        }
        if (position.y < depth) {
            adjacentPositions.add(Position(position.x, position.y + 1))
        }
        return adjacentPositions.filter {
            (it.getHeight() <= position.getHeight() + 1)
                    && (it.getDistance() > position.getDistance())
        }
    }

    fun createPaths(point: Position) {
        val adjacent = getAdjacent(point)
        adjacent.forEach {
            if (it.getDistance() > point.getDistance() + 1) {
                it.setDistance(point.getDistance() + 1)
                createPaths(it)
            }
        }
    }

    val start = findStart()
    start.setDistance(0)
    createPaths(start)
    println(distances.map {it.joinToString(",")}.joinToString("\n"))
    println(findEndpoint().getDistance())
}
