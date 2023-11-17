package aoc2022.day13.part1

import utils.CollectionUtils.Companion.partitionByElement
import utils.StringUtils.Companion.isInt
import utils.StringUtils.Companion.toLists

data class PacketPair(val left: String, val right: String, var valid: Boolean = true) {
    fun getLeftValues(): List<String> {
        return left.toLists()
    }

    fun getRightValues(): List<String> {
        return right.toLists()
    }
}

fun main() {
    val packetPairs = INPUT.split("\n").partitionByElement("").map {
        PacketPair(
            it[0],
            it[1]
        )
    }

    arePacketsValid(packetPairs)

    println(
        packetPairs.mapIndexed { index, packetPair ->
            Pair(index, packetPair)
        }.filter {
            it.second.valid
        }.sumOf {
            it.first + 1
        }
    )
}

fun arePacketsValid(packets: List<PacketPair>) {
    packets.forEach {
        it.valid = compareValues(it.getLeftValues(), it.getRightValues())!!
    }
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
