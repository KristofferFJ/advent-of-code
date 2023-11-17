package aoc2022.day10.part2

fun main() {
    val input = INPUT.split("\n")

    fun partTwo() {
        var spritePosition = 0
        var cycle = 0
        var row = 0
        val grid: List<MutableList<String>> = IntRange(1, 6).map {
            IntRange(1, 40).map { "." }.toMutableList()
        }

        fun drawPosition() {
            if(cycle == 40) {
                row += 1
                cycle = 0
            }
            if(cycle >= spritePosition && cycle <= spritePosition + 2) {
                grid[row][cycle] = "X"
            }
        }

        fun handleInstruction(value: Int, instructionCycle: Int) {
            if(instructionCycle == 0) {
                drawPosition()
                cycle += 1
                return handleInstruction(value, 1)
            } else {
                drawPosition()
                cycle += 1
                spritePosition += value
            }
        }

        input.forEach {
            if(it.startsWith("noop")) {
                drawPosition()
                cycle += 1
            }
            if(it.startsWith("addx")) {
                handleInstruction(it.substring(5).toInt(), 0)
            }
        }
        println(grid.map { it.joinToString("") }.joinToString("\n"))
    }

    partTwo()
}
