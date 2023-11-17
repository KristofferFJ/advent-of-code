package aoc2022.day14.part2

import utils.IntUtils.Companion.toward
import utils.CollectionUtils.Companion.printGrid

fun main() {
    val wallInstructions = INPUT.split("\n").map {
        it.split(" -> ")
    }
    val floorY = wallInstructions.flatMap { it.map { it.split(",")[1].toInt() } }.max() + 2
    val grid: MutableList<MutableList<String>> = IntRange(0, floorY).map {
        IntRange(0, 1000).map {
            " "
        }.toMutableList()
    }.toMutableList()
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
            if(y == floorY - 1) {
                grid[y][x] = "O"
                droppedSands += 1
                x = 500; y = 0
            }
            if(grid[y + 1][x] == " ") { y += 1; continue }
            if(grid[y + 1][x - 1] == " ") { y += 1; x -= 1; continue}
            if(grid[y + 1][x + 1] == " ") { y += 1; x += 1; continue} else if(x == 500 && y == 0) {
                grid[y][x] = "O"
                droppedSands += 1
                break
            }
            grid[y][x] = "O"
            droppedSands += 1
            x = 500; y = 0
        }
        return droppedSands
    }

    println(dropSand())
    println(grid.printGrid())
}
