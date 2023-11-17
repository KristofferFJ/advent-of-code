package aoc2022.day09.part2

import kotlin.math.absoluteValue

fun main() {
    val instructions = INPUT.split("\n").map { Instruction(it.substring(0, 1), it.substring(2).toInt()) }

    fun partTwo() {
        val grid = Grid()
        instructions.forEach {
            grid.doInstruction(it)
        }
        println(grid.visited.size)
    }

    partTwo()
}

data class Instruction(val direction: String, val moves: Int)
data class Position(var x: Int, var y: Int) {
    fun approachX(position: Position) {
        if (position.x > this.x) this.x += 1
        else this.x -= 1
    }

    fun approachY(position: Position) {
        if (position.y > this.y) this.y += 1
        else this.y -= 1
    }
}

data class Grid(
    val head: Position = Position(0, 0),
    val tail1: Position = Position(0, 0),
    val tail2: Position = Position(0, 0),
    val tail3: Position = Position(0, 0),
    val tail4: Position = Position(0, 0),
    val tail5: Position = Position(0, 0),
    val tail6: Position = Position(0, 0),
    val tail7: Position = Position(0, 0),
    val tail8: Position = Position(0, 0),
    val tail9: Position = Position(0, 0),
    val visited: MutableSet<Position> = mutableSetOf(Position(0, 0))
) {

    fun doInstruction(instruction: Instruction) {
        for (move in 0 until instruction.moves) {
            moveHead(instruction.direction)
            movePointTowardsPoint(tail1, head)
            movePointTowardsPoint(tail2, tail1)
            movePointTowardsPoint(tail3, tail2)
            movePointTowardsPoint(tail4, tail3)
            movePointTowardsPoint(tail5, tail4)
            movePointTowardsPoint(tail6, tail5)
            movePointTowardsPoint(tail7, tail6)
            movePointTowardsPoint(tail8, tail7)
            movePointTowardsPoint(tail9, tail8)
            markTailPosition()
        }
    }

    private fun moveHead(direction: String) {
        if (direction == "U") head.y += 1
        if (direction == "D") head.y -= 1
        if (direction == "R") head.x += 1
        if (direction == "L") head.x -= 1
    }

    private fun movePointTowardsPoint(pointToMove: Position, pointToMoveTo: Position) {
        if ((pointToMove.x - pointToMoveTo.x).absoluteValue == 2) {
            pointToMove.approachX(pointToMoveTo)
            if ((pointToMove.y - pointToMoveTo.y).absoluteValue > 0) pointToMove.approachY(pointToMoveTo)
        } else if ((pointToMoveTo.y - pointToMove.y).absoluteValue == 2) {
            pointToMove.approachY(pointToMoveTo)
            if ((pointToMoveTo.x - pointToMove.x).absoluteValue > 0) pointToMove.approachX(pointToMoveTo)
        }
    }

    private fun markTailPosition() {
        visited.add(this.tail9.copy())
    }
}
