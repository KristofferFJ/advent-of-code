package aoc2022.day18.part1

private data class Point(val x: Int, val y: Int, val z: Int)
private data class Side(val vectorStart: Point, val vectorEnd: Point)
private data class Cube(val x: Int, val y: Int, val z: Int) {
    fun getSides(): Set<Side> {
        return setOf(
            Side(Point(x, y, z), Point(x + 1, y + 1, z)),
            Side(Point(x, y, z), Point(x + 1, y, z + 1)),
            Side(Point(x, y, z), Point(x, y + 1, z + 1)),
            Side(Point(x + 1, y, z), Point(x + 1, y + 1, z + 1)),
            Side(Point(x, y + 1, z), Point(x + 1, y + 1, z + 1)),
            Side(Point(x, y, z + 1), Point(x + 1, y + 1, z + 1)),
        )
    }
}

fun main() {
    val cubes = INPUT.split("\n").map { it.split(",") }.map {
        Cube(
            it[0].toInt(),
            it[1].toInt(),
            it[2].toInt()
        )
    }

    val groupedSides = cubes.flatMap { it.getSides() }.groupBy { it }
    println(groupedSides.filter { it.value.size == 1 }.size)
}
