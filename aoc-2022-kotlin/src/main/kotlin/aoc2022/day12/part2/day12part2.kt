package aoc2022.day12.part2

import utils.LetterUtil.Companion.getLetterValue

fun main() {
    val grid: List<List<Char>> = INPUT.split("\n").map { it.toList() }
    val depth = grid.size - 1
    val length = grid.first().size - 1

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

        fun getDistance(distances: List<MutableList<Int>>): Int {
            return distances[this.y][this.x]
        }

        fun setDistance(distance: Int, distances: List<MutableList<Int>>) {
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

    fun findStartPoints(): List<Position> {
        return grid.flatMapIndexed { indexRow, chars ->
            chars.mapIndexedNotNull { indexColumn, char ->
                if (char == 'S' || char == 'a') Position(indexColumn, indexRow) else null
            }
        }
    }

    fun getAdjacent(position: Position, distances: List<MutableList<Int>>): List<Position> {
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
                    && (it.getDistance(distances) > position.getDistance(distances))
        }
    }

    fun createPaths(point: Position, distances: List<MutableList<Int>>): List<MutableList<Int>> {
        val adjacent = getAdjacent(point, distances)
        adjacent.forEach {
            if (it.getDistance(distances) > point.getDistance(distances) + 1) {
                it.setDistance(point.getDistance(distances) + 1, distances)
                createPaths(it, distances)
            }
        }
        return distances
    }

    fun createPaths() {
        val startPoints = findStartPoints()
        println(startPoints.map {
            val distances: List<MutableList<Int>> = grid.map { it.map { (depth + 1) * (length + 1) }.toMutableList() }
            it.setDistance(0, distances)
            findEndpoint().getDistance(createPaths(it, distances))
        }.sortedBy { it })
    }

    createPaths()
}
