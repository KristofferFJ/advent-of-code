package aoc2020.day01

fun main() {
    fun partOne() {
        input.forEach { number ->
            input.find { it + number == 2020 }?.let { println(it.times(number)); return }
        }
    }

    fun partTwo() {
        input.forEach { firstNumber ->
            input.forEach { secondNumber ->
                input.find { it + firstNumber + secondNumber == 2020 }?.let { println(it.times(firstNumber).times(secondNumber)); return }
            }
        }
    }

    partOne()
    partTwo()
}

private val input = """""".trimIndent()
    .split("\n")
    .map { it.toInt() }
