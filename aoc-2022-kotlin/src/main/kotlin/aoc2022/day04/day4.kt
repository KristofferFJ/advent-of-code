package aoc2022.day04

val sections: List<Section> = INPUT.split("\n").map {
    it.split(",").map {
        it.split("-")
            .map { it.toInt() }
    }
}.map {
    Section(
        Assignment(it[0][0], it[0][1]),
        Assignment(it[1][0], it[1][1]),
    )
}


data class Assignment(val from: Int, val to: Int)
data class Section(val one: Assignment, val two: Assignment) {
    fun overlapsCompletely(): Boolean {
        if(one.from <= two.from && one.to >= two.to) return true
        if(two.from <= one.from && two.to >= one.to) return true
        return false
    }

    fun overlaps(): Boolean {
        if((one.to <= two.to && one.to >= two.from) || (one.from <= two.to && one.from >= two.from)) return true
        if((two.to <= one.to && two.to >= one.from) || (two.from <= one.to && two.from >= one.from)) return true
        return false
    }
}

fun partOne(): Int {
    return sections.filter { it.overlapsCompletely() }.size
}

fun partTwo(): Int {
    return sections.filter { it.overlaps() }.size
}

fun main() {
    println(partOne())
    println(partTwo())
}
