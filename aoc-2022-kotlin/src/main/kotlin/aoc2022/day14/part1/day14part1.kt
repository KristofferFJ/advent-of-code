package aoc2022.day14.part1

import utils.IntUtils.Companion.toward
import utils.CollectionUtils.Companion.printGrid

fun main() {
    val grid: MutableList<MutableList<String>> = IntRange(0, 1000).map {
        IntRange(0, 1000).map { " " }.toMutableList()
    }.toMutableList()
    val wallInstructions = INPUT.split("\n").map {
        it.split(" -> ")
    }
    for (walls in wallInstructions) {
        for (i in 0 until walls.size - 1) {
            val (fromX, fromY) = walls[i].split(",").map { it.toInt() }
            val (toX, toY) = walls[i + 1].split(",").map { it.toInt() }
            for (x in fromX toward toX) {
                grid[fromY][x] = "X"
            }
            for (y in fromY toward toY) {
                grid[y][fromX] = "X"
            }
        }
    }

    fun dropSand(): Int {
        var x = 500; var y = 0; var droppedSands = 0
        while (true) {
            if(y == 900) break
            if(grid[y + 1][x] == " ") { y += 1; continue }
            if(grid[y + 1][x - 1] == " ") { y += 1; x -= 1; continue}
            if(grid[y + 1][x + 1] == " ") { y += 1; x += 1; continue}
            grid[y][x] = "O"
            droppedSands += 1
            x = 500; y = 0
        }
        return droppedSands
    }

    println(dropSand())
    println(grid.printGrid())
}
