package aoc2022.day22.part1

private data class Point(var x: Int, var y: Int) {
    fun moveDirection(direction: String): Point {
        if (direction == "U") this.y -= 1
        if (direction == "D") this.y += 1
        if (direction == "R") this.x += 1
        if (direction == "L") this.x -= 1
        return this
    }
}

fun getReverseDirection(direction: String): String {
    if (direction == "D") return "U"
    if (direction == "U") return "D"
    if (direction == "L") return "R"
    return "L"
}

private data class Grid(
    val coordinates: List<MutableList<String>>,
    var position: Point = Point(0, 0),
    var direction: String = "R"
) {
    fun setupStart() {
        if (seePoint(position) != ".") {
            position.moveDirection(direction); setupStart()
        }
    }

    fun move() {
        var newPosition = this.position.copy()
        if (atEnd(position, direction) || seePoint(newPosition.moveDirection(direction)) == " ") {
            newPosition = wrapAround(newPosition, direction)
            if (seePoint(newPosition) == ".") this.position = newPosition
        }
        if (seePoint(newPosition) == ".") this.position = newPosition
    }

    private fun wrapAround(position: Point, direction: String): Point {
        val reverseDirection = getReverseDirection(direction)
        var finalPosition = position
        val newPositionCheck = position.copy()
        while (!atEnd(newPositionCheck, reverseDirection) &&
            seePoint(newPositionCheck.moveDirection(reverseDirection)) != " "
        ) {
            finalPosition = newPositionCheck.copy()
        }
        return finalPosition
    }

    private fun atEnd(newPositionCheck: Point, direction: String): Boolean {
        if (direction == "R") {
            return this.coordinates[0].size == newPositionCheck.x + 1
        }
        if (direction == "L") {
            return 0 == newPositionCheck.x
        }
        if (direction == "U") {
            return 0 == newPositionCheck.y
        }
        return this.coordinates.size == newPositionCheck.y + 1
    }

    fun seePoint(point: Point): String {
        return this.coordinates[point.y][point.x]
    }

    fun turn(leftOrRight: String) {
        this.direction = if (leftOrRight == "L") {
            when (direction) {
                "U" -> "L"; "L" -> "D"; "D" -> "R"; "R" -> "U"; else -> throw Exception()
            }
        } else {
            when (direction) {
                "U" -> "R"; "R" -> "D"; "D" -> "L"; "L" -> "U"; else -> throw Exception()
            }
        }
    }
}

fun main() {
    val (coordinatesInput, rulesInput) = INPUT.split("\n\n")
    val coordinates = coordinatesInput.split("\n").map { it.toList().map { it.toString() }.toMutableList() }
    val grid = Grid(coordinates)
    val turningRules = rulesInput.split("\\d+".toRegex()).filter { it != "" }
    val movingRules = rulesInput.split("[a-zA-Z]".toRegex()).map { it.toInt() }
    grid.setupStart()

    movingRules.forEachIndexed { index, moves ->
        repeat(moves) {
            grid.move()
        }
        if (turningRules.size > index) grid.turn(turningRules[index])
    }

    println(
        listOf(
            listOf("R", "D", "L", "U").indexOf(grid.direction), 4 * (grid.position.x + 1), 1000 * (grid.position.y + 1)
        ).sum()
    )
}
