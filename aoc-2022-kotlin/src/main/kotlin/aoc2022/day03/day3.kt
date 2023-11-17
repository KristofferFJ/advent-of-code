package aoc2022.day03

import utils.LetterUtil.Companion.getLetterValue
import utils.CollectionUtils.Companion.partitionListsInGroupsOfSize

class Rucksack(val left: List<Char>, val right: List<Char>) {
    fun getIntersection(): Char {
        val intersection = left.intersect(right.toSet())
        check(intersection.size == 1)
        return intersection.first()
    }
}

class Group(val rucksack1: Rucksack, val rucksack2: Rucksack, val rucksack3: Rucksack) {
    fun getIntersection(): Char {
        val intersection = (rucksack1.left + rucksack1.right)
            .intersect((rucksack2.left + rucksack2.right).toSet())
            .intersect((rucksack3.left + rucksack3.right).toSet())
        check(intersection.size == 1)
        return intersection.first()
    }
}

fun getRucksacks(): List<Rucksack> {
    return INPUT.split("\n").map {
        val size = it.length
        Rucksack(it.substring(0, size / 2).toList(), it.substring(size / 2).toList())
    }
}

fun getGroups(): List<Group> {
    return getRucksacks().partitionListsInGroupsOfSize(3).map { Group(it[0], it[1], it[2]) }
}


fun partOne() {
    println(getRucksacks().sumOf { it.getIntersection().getLetterValue() })
}

fun partTwo() {
    println(getGroups().sumOf { it.getIntersection().getLetterValue() })
}

fun main() {
    partOne()
    partTwo()
}
