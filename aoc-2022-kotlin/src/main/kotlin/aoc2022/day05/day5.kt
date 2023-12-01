package aoc2022.day05

import utils.CollectionUtils.Companion.removeSlice
import utils.CollectionUtils.Companion.transpose

val TOWERS = INPUT.split("\n\n").first()
val INSTRUCTIONS = INPUT.split("\n\n").last()

fun getTowers(): List<MutableList<String>> {
    val transposed = TOWERS.replace("[", "").replace("]", "").replace("    ", ",").replace(" ", ",").split("\n")
        .map { it.split(",") }.transpose()
    transposed.map { it.reverse() }
    transposed.forEach { it.removeAll { it.isBlank() } }
    return transposed
}

data class Instruction(val amount: Int, val fromTower: Int, val toTower: Int)

fun getInstructions(): List<Instruction> {
    return INSTRUCTIONS.split("\n").filter { it.isNotEmpty() }.map { it.substring(5).split(" from | to ".toRegex()) }
        .map { Instruction(it[0].toInt(), it[1].toInt() - 1, it[2].toInt() - 1) }
}

fun doInstruction(instruction: Instruction, towers: List<MutableList<String>>) {
    val removedSlice = towers[instruction.fromTower].removeSlice(instruction.amount)
    towers[instruction.toTower].addAll(removedSlice)
}

fun doInstructionMaintainingOrder(instruction: Instruction, towers: List<MutableList<String>>) {
    val removedSlice = towers[instruction.fromTower].removeSlice(instruction.amount)
    removedSlice.reverse()
    towers[instruction.toTower].addAll(removedSlice)
}

fun partOne(): String {
    val towers = getTowers()
    val instructions = getInstructions()
    instructions.forEach { doInstruction(it, towers) }
    return towers.joinToString("") { it.last() }
}

fun partTwo(): String {
    val towers = getTowers()
    val instructions = getInstructions()
    instructions.forEach { doInstructionMaintainingOrder(it, towers) }
    return towers.joinToString("") { it.last() }
}

fun main() {
    println(partOne())
    println(partTwo())
}
