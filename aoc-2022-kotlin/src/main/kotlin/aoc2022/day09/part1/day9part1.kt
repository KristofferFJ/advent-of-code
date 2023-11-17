package aoc2022.day09.part1

import kotlin.math.absoluteValue

fun main() {
    val instructions = INPUT.split("\n").map { Instruction(it.substring(0, 1), it.substring(2).toInt()) }

    fun partOne() {
        val grid = Grid()
        instructions.forEach {
            grid.doInstruction(it)
        }
        println(grid.visited.size)
    }

    partOne()
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
    val tail: Position = Position(0, 0),
    val visited: MutableSet<Position> = mutableSetOf(Position(0, 0))
) {
    fun doInstruction(instruction: Instruction) {
        for (move in 0 until instruction.moves) {
            moveHead(instruction.direction)
            moveTail()
            markTailPosition()
        }
    }

    private fun moveHead(direction: String) {
        if (direction == "U") head.y += 1
        if (direction == "D") head.y -= 1
        if (direction == "R") head.x += 1
        if (direction == "L") head.x -= 1
    }

    private fun moveTail() {
        if ((head.x - tail.x).absoluteValue == 2) {
            tail.approachX(head)
            if ((head.y - tail.y).absoluteValue > 0) tail.approachY(head)
        } else if ((head.y - tail.y).absoluteValue == 2) {
            tail.approachY(head)
            if ((head.x - tail.x).absoluteValue > 0) tail.approachX(head)
        }
    }

    private fun markTailPosition() {
        visited.add(this.tail.copy())
    }
}
