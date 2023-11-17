package aoc2022.day13.part2

import utils.CollectionUtils.Companion.prod
import utils.StringUtils.Companion.isInt
import utils.StringUtils.Companion.toLists


fun main() {
    val packetPairs = INPUT.split("\n").filter { it != "" }.toMutableList()
    packetPairs.addAll(listOf("[[2]]", "[[6]]"))
    println(
        orderPackets(packetPairs).reversed()
            .mapIndexed { index, packetPair ->
                Pair(index, packetPair)
            }.filter {
                (it.second == "[[2]]" || it.second == "[[6]]")
            }.map {
                it.first + 1
            }.prod()
    )
}

fun compareValues(leftValues: List<String>, rightValues: List<String>): Boolean? {
    leftValues.forEachIndexed { index, left ->
        if (index > rightValues.size - 1) return false
        if (compareValues(left, rightValues[index]) != null) {
            return compareValues(left, rightValues[index])!!
        }
    }
    if (rightValues.size > leftValues.size) return true
    return null
}

fun compareValues(leftValue: String, rightValue: String): Boolean? {
    if (leftValue.isInt() && rightValue.isInt()) {
        if (leftValue.toInt() == rightValue.toInt()) return null
        return leftValue.toInt() < rightValue.toInt()
    }
    if (leftValue.isInt()) {
        return compareValues(listOf(leftValue), rightValue.toLists())
    }
    if (rightValue.isInt()) {
        return compareValues(leftValue.toLists(), listOf(rightValue))
    }
    return compareValues(leftValue.toLists(), rightValue.toLists())
}

fun orderPackets(packets: List<String>): List<String> {
    return packets.sortedWith { a, b ->
        val result = compareValues(a.toLists(), b.toLists())
        when {
            (result == true) -> 1
            (result == false) -> -1
            else -> 0
        }
    }
}
