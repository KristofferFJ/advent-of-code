package aoc2022.day08

fun main() {
    fun partOne() {
        var sum = 0
        val height = 99
        val columns = 99
        val grid = INPUT.split("\n").map { it.toList().map { it.toString().toInt() } }
        grid.forEachIndexed { row, rows ->
            rows.forEachIndexed { column, tree ->
                if (row == 0 || row == height - 1 || column == 0 || column == columns - 1) {
                    sum += 1
                    return@forEachIndexed
                }
                if (grid[row].subList(0, column).max() < tree) {
                    sum += 1
                    return@forEachIndexed
                }
                if (grid[row].subList(column + 1, columns).max() < tree) {
                    sum += 1
                    return@forEachIndexed
                }
                var maxFromTop = 0
                for (i in 0 until row) {
                    if (grid[i][column] > maxFromTop) {
                        maxFromTop = grid[i][column]
                    }
                }
                if (maxFromTop < tree) {
                    sum += 1
                    return@forEachIndexed
                }
                var maxFromBottom = 0
                for (i in row + 1 until height) {
                    if (grid[i][column] > maxFromBottom) {
                        maxFromBottom = grid[i][column]
                    }
                }
                if (maxFromBottom < tree) {
                    sum += 1
                    return@forEachIndexed
                }
            }
        }
        println(sum)
    }

    fun partTwo() {
        val scenicScores = mutableListOf<Int>()
        val grid = INPUT.split("\n").map { it.toList().map { it.toString().toInt() } }
        val height = grid.size
        val columns = grid.first().size
        grid.forEachIndexed { row, rows ->
            rows.forEachIndexed { column, tree ->
                var left = 0; var right = 0; var up = 0; var down = 0
                if (column == 0 || column == height - 1 || row == 0 || row == height - 1) return@forEachIndexed
                for (i in 1..height) {
                    if (row - i == 0) {
                        up += 1; break
                    }
                    if (grid[row - i][column] < tree) {
                        up += 1
                    } else {
                        up += 1
                        break
                    }
                }
                for (i in 1..height) {
                    if (row + i == height - 1) {
                        down += 1; break
                    }
                    if (grid[row + i][column] < tree) {
                        down += 1
                    } else {
                        down += 1
                        break
                    }
                }
                for (i in 1..columns) {
                    if (column + i == columns - 1) {
                        right += 1; break
                    }
                    if (grid[row][column + i] < tree) {
                        right += 1
                    } else {
                        right += 1
                        break
                    }
                }
                for (i in 1..columns) {
                    if (column - i == 0) {
                        left += 1; break
                    }
                    if (grid[row][column - i] < tree) {
                        left += 1
                    } else {
                        left += 1
                        break
                    }
                }
                scenicScores.add(left * right * up * down)
            }
        }
        println(scenicScores.max())
    }

    partOne()
    partTwo()
}
