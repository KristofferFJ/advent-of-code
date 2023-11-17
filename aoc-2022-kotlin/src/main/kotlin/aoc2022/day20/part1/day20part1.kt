package aoc2022.day20.part1

data class IndexedNumber(
    val index: Int,
    val number: Int
)

fun main() {
    val originalList = INPUT.split("\n").mapIndexed { index, number ->
        IndexedNumber(index, number.toInt())
    }.toMutableList()
    val length = originalList.size
    IntRange(0, length - 1).forEach { index ->
        val indexedNumber = originalList.find { it.index == index }!!
        val startIndex = originalList.indexOf(indexedNumber)
        val endIndex = (startIndex + indexedNumber.number).mod(length - 1)
        originalList.removeAt(startIndex)
        originalList.add(endIndex, indexedNumber)
    }
    val indexOfZero = originalList.indexOf(originalList.find { it.number == 0 })
    println(listOf(1000, 2000, 3000).sumOf {
        originalList[(it + indexOfZero) % length].number
    })
}
