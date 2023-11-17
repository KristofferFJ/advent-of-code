package aoc2022.day10.part1

fun main() {
    val input = INPUT.split("\n")

    fun partOne() {
        var startValue = 1
        var cycle = 1
        val savedValues = mutableMapOf<Int, Int>()

        fun saveImportantValue() {
            if ((cycle - 20) % 40 == 0 && cycle < 221) {
                savedValues[cycle] = cycle * startValue
            }
        }

        fun handleInstruction(value: Int, instructionCycle: Int) {
            if (instructionCycle == 0) {
                saveImportantValue()
                cycle += 1
                return handleInstruction(value, 1)
            } else {
                saveImportantValue()
                cycle += 1
                startValue += value
            }
        }

        input.forEach {
            if (it.startsWith("noop")) {
                saveImportantValue()
                cycle += 1
            }
            if (it.startsWith("addx")) {
                handleInstruction(it.substring(5).toInt(), 0)
            }
        }
        println(savedValues.values.sum())
    }

    partOne()
}
