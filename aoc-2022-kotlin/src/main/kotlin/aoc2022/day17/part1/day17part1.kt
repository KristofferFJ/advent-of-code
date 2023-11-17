package aoc2022.day17.part1

import utils.CollectionUtils.Companion.printGrid

val chamber: MutableList<MutableList<String>> = IntRange(0, 3500).map {
    IntRange(0, 6).map { " " }.toMutableList()
}.toMutableList()
var jetIndex = 0

fun main() {
    IntRange(1, 2022).forEachIndexed { index, i ->
        val spawnHeight = findSpawnHeight()
        spawnShape(index, spawnHeight)
        fallRock()
    }
    println(chamber.reversed().printGrid())
    println(findSpawnHeight() - 3)
}

fun fallRock() {
    while (true) {
        val direction = if (INPUT[jetIndex % INPUT.length] == '<') -1 else 1
        useJets(direction)
        jetIndex += 1
        if (!drop()) {
            turnToStone()
            return
        }
    }
}

private fun turnToStone() {
    chamber.forEachIndexed { rowIndex, row ->
        row.forEachIndexed { columnIndex, stone ->
            if (stone == "@") chamber[rowIndex][columnIndex] = "#"
        }
    }
}

fun drop(): Boolean {
    chamber.forEachIndexed { rowIndex, row ->
        if (!row.any { it == "@" }) return@forEachIndexed
        if (rowIndex == 0) return false
        row.forEachIndexed { index, stone ->
            if (stone != "@") return@forEachIndexed
            if (chamber[rowIndex - 1][index] == "#") {
                return false
            }
        }
        row.forEachIndexed { columnIndex, stone ->
            if (stone == "@") {
                chamber[rowIndex - 1][columnIndex] = "@"
                chamber[rowIndex][columnIndex] = " "
            }
        }
    }
    return true
}

fun useJets(direction: Int) {
    chamber.forEachIndexed { rowIndex, row ->
        if (!row.any { it == "@" }) return@forEachIndexed
        row.forEachIndexed { index, stone ->
            if (stone != "@") return@forEachIndexed
            if (direction == 1 && (index == 6 || row[index + 1] == "#")) {
                return
            }
            if (direction == -1 && (index == 0 || row[index - 1] == "#")) {
                return
            }
        }
    }
    chamber.forEachIndexed { rowIndex, row ->
        if (direction == 1) {
            row.reversed().forEachIndexed { columnIndex, stone ->
                if (stone == "@") {
                    chamber[rowIndex][6 - columnIndex + 1] = "@"
                    chamber[rowIndex][6 - columnIndex] = " "
                }
            }
        }
        if (direction == -1) {
            row.forEachIndexed { columnIndex, stone ->
                if (stone == "@") {
                    chamber[rowIndex][columnIndex - 1] = "@"
                    chamber[rowIndex][columnIndex] = " "
                }
            }
        }
    }
}

fun spawnShape(shapeNumber: Int, spawnHeight: Int) {
    when (shapeNumber % 5) {
        0 -> chamber[spawnHeight] = mutableListOf(" ", " ", "@", "@", "@", "@", " ")
        1 -> {
            chamber[spawnHeight + 2] = mutableListOf(" ", " ", " ", "@", " ", " ", " ")
            chamber[spawnHeight + 1] = mutableListOf(" ", " ", "@", "@", "@", " ", " ")
            chamber[spawnHeight + 0] = mutableListOf(" ", " ", " ", "@", " ", " ", " ")
        }
        2 -> {
            chamber[spawnHeight + 2] = mutableListOf(" ", " ", " ", " ", "@", " ", " ")
            chamber[spawnHeight + 1] = mutableListOf(" ", " ", " ", " ", "@", " ", " ")
            chamber[spawnHeight + 0] = mutableListOf(" ", " ", "@", "@", "@", " ", " ")
        }
        3 -> {
            chamber[spawnHeight + 3] = mutableListOf(" ", " ", "@", " ", " ", " ", " ")
            chamber[spawnHeight + 2] = mutableListOf(" ", " ", "@", " ", " ", " ", " ")
            chamber[spawnHeight + 1] = mutableListOf(" ", " ", "@", " ", " ", " ", " ")
            chamber[spawnHeight + 0] = mutableListOf(" ", " ", "@", " ", " ", " ", " ")
        }
        4 -> {
            chamber[spawnHeight + 1] = mutableListOf(" ", " ", "@", "@", " ", " ", " ")
            chamber[spawnHeight + 0] = mutableListOf(" ", " ", "@", "@", " ", " ", " ")
        }
    }
}

fun findSpawnHeight(): Int {
    return chamber.indexOf(chamber.find { it.all { it == " " } }) + 3
}
