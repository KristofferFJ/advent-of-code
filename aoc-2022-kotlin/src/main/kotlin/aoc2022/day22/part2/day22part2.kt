package aoc2022.day22.part2

private data class Point(var x: Int, var y: Int) {
    fun moveDirection(direction: String): Point {
        if (direction == "U") this.y -= 1
        if (direction == "D") this.y += 1
        if (direction == "R") this.x += 1
        if (direction == "L") this.x -= 1
        return this
    }
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
        val startPosition = this.position.copy()
        val startDirection = this.direction
        val checkPosition = this.position.copy()
        if (atEnd(position, direction) || seePoint(checkPosition.moveDirection(direction)) == " ") {
            wrapAround()
            if (seePoint(position) != ".") {
                position = startPosition
                direction = startDirection
            }
            return
        }
        val newPosition = startPosition.copy()
        if (seePoint(newPosition.moveDirection(direction)) == ".") this.position = newPosition
    }

    private fun wrapAround() {
        when(direction) {
            "D" -> wrapDown()
            "U" -> wrapUp()
            "L" -> wrapLeft()
            "R" -> wrapRight()
        }
    }

    private fun wrapDown() {
        if (IntRange(0, 49).contains(position.x)) {
            position.x += 100
            position.y = 0
            return
        }
        if (IntRange(50, 99).contains(position.x)) {
            position.y = position.x + 100
            position.x = 49
            direction = "L"
            return
        }
        if (IntRange(100, 149).contains(position.x)) {
            position.y = position.x - 50
            position.x = 99
            direction = "L"
            return
        }
    }

    private fun wrapUp() {
        if (IntRange(0, 49).contains(position.x)) {
            position.y = position.x + 50
            position.x = 50
            direction = "R"
            return
        }
        if (IntRange(50, 99).contains(position.x)) {
            position.y = position.x + 100
            position.x = 0
            direction = "R"
            return
        }
        if (IntRange(100, 149).contains(position.x)) {
            position.x -= 100
            position.y = 199
            return
        }
    }

    private fun wrapRight() {
        if (IntRange(0, 49).contains(position.y)) {
            position.y = 149 - position.y
            position.x = 99
            direction = "L"
            return
        }
        if (IntRange(50, 99).contains(position.y)) {
            position.x = position.y + 50
            position.y = 49
            direction = "U"
            return
        }
        if (IntRange(100, 149).contains(position.y)) {
            position.y = 149 - position.y
            position.x = 149
            direction = "L"
            return
        }
        if (IntRange(150, 199).contains(position.y)) {
            position.x = position.y - 100
            position.y = 149
            direction = "U"
        }
    }

    private fun wrapLeft() {
        if (IntRange(0, 49).contains(position.y)) {
            position.y = 149 - position.y
            position.x = 0
            direction = "R"
            return
        }
        if (IntRange(50, 99).contains(position.y)) {
            position.x = position.y - 50
            position.y = 100
            direction = "D"
            return
        }
        if (IntRange(100, 149).contains(position.y)) {
            position.y = 149 - position.y
            position.x = 50
            direction = "R"
            return
        }
        if (IntRange(150, 199).contains(position.y)) {
            position.x = position.y - 100
            position.y = 0
            direction = "D"
        }
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
