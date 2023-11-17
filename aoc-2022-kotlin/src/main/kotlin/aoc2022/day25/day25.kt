package aoc2022.day25

import utils.CollectionUtils.Companion.fromValue
import utils.CollectionUtils.Companion.isNotIn
import utils.LongUtils.Companion.pow

private val conversion = mapOf("2" to 2, "1" to 1, "0" to 0, "-" to -1, "=" to -2)
private fun main() {
    val sum = INPUT.split("\n").sumOf { it.toDecimal() }
    println(sum.toSnafu())
}

private fun String.toDecimal(): Long {
    return this.toList().reversed().mapIndexed { index, char ->
        char.toDecimal().times(5L.pow(index))
    }.sumOf { it }
}

private fun Long.toSnafu(): String {
    var remaining = this
    var result = ""
    while (remaining > 0) {
        var part = remaining.mod(5)
        if (part.isNotIn(conversion.values)) {
            part -= 5
            remaining += 5
        }
        result += conversion.fromValue(part)
        remaining = remaining.div(5)
    }
    return result.reversed()
}

private fun Char.toDecimal(): Long {
    return conversion[this.toString()]!!.toLong()
}
