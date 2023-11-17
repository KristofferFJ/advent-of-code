package aoc2022.day17.part2

import utils.Grid
import utils.Point
import kotlin.math.max

var jetIndex = 0
val chamber = Grid(".........")
fun main() {
    var highestPoint = 0
    var atIndex = 0
    var shapes = 0
    LongRange(1, 1757 + 2 * 1755 + 1403).forEachIndexed { index, i ->
        if(chamber.rows.size > 1000) {
            chamber.rows = chamber.rows.subList(250, chamber.rows.size)
        }
        if(jetIndex % INPUT.length == 0 && index % 5 == 0) {
            println(chamber)
        }
        //did the rest by hand. After the first iteration of jets, tower was 2722 high and used 1757 shapes.
        // Then for every iteration it would use 1755 shapes for 2747 height. Leaving remainder of 1403 shapes
        // which adds height of 10399 - 8216.
        if(jetIndex % INPUT.length == 0) {
            println("height=" + (highestPoint - atIndex))
            println("shapes=" + (index - shapes))
            println("highestPointNow=$highestPoint")
            println("")
            atIndex = highestPoint
            shapes = index
        }
        val shape = spawnShape(index, highestPoint)
        highestPoint = max(fallRock(shape), highestPoint)
    }
    print(highestPoint)
}

fun fallRock(shape: List<Point>): Int {
    var movingShape = shape
    while (true) {
        val direction = if (INPUT[jetIndex % INPUT.length] == '<') -1 else 1
        movingShape = useJets(movingShape, direction)
        jetIndex += 1
        if (ableToDrop(movingShape)) {
            movingShape = movingShape.map { it.addY(-1) }
            continue
        }
        turnToStone(movingShape)
        return movingShape.maxOf { it.y } + 1
    }
}

private fun turnToStone(shape: List<Point>) {
    shape.forEach { chamber.set(it.x, it.y, "#") }
}

fun ableToDrop(shape: List<Point>): Boolean {
    return shape.all { chamber.at(it.addY(-1)) == "." } && shape.all { it.y > 0 }
}

val edges = listOf(0, 8)
fun useJets(shape: List<Point>, direction: Int): List<Point> {
    if (shape.any { it.x + direction in edges || chamber.at(it.addX(direction)) == "#" }) return shape
    return shape.map { it.addX(direction) }
}

fun spawnShape(shapeNumber: Int, height: Int): List<Point> {
    return when (shapeNumber % 5) {
        0 -> mutableListOf(Point(3, 3 + height), Point(4, 3 + height), Point(5, 3 + height), Point(6, 3 + height))
        1 -> mutableListOf(
            Point(4, 3 + height),
            Point(3, 4 + height), Point(4, 4 + height), Point(5, 4 + height),
            Point(4, 5 + height),
        )

        2 -> mutableListOf(
            Point(5, 5 + height),
            Point(5, 4 + height),
            Point(5, 3 + height), Point(4, 3 + height), Point(3, 3 + height)
        )

        3 -> mutableListOf(
            Point(3, 6 + height),
            Point(3, 5 + height),
            Point(3, 4 + height),
            Point(3, 3 + height)
        )

        4 -> mutableListOf(
            Point(3, 4 + height), Point(4, 4 + height),
            Point(3, 3 + height), Point(4, 3 + height)
        )

        else -> throw Exception("")
    }
}
