package aoc2022.day20.part2

data class IndexedNumber(
    val index: Int,
    val number: Long
)

const val DECRYPTION_KEY = 811589153L

fun main() {
    val originalList = INPUT.split("\n").mapIndexed { index, number ->
        IndexedNumber(index, number.toLong().times(DECRYPTION_KEY))
    }.toMutableList()
    val length = originalList.size
    repeat(10) {
        IntRange(0, length - 1).forEach { index ->
            val indexedNumber = originalList.find { it.index == index }!!
            val startIndex = originalList.indexOf(indexedNumber)
            val endIndex = (startIndex + indexedNumber.number).mod(length - 1)
            originalList.removeAt(startIndex)
            originalList.add(endIndex, indexedNumber)
        }
    }
    val indexOfZero = originalList.indexOf(originalList.find { it.number == 0L })
    println(listOf(1000, 2000, 3000).sumOf {
        originalList[(it + indexOfZero) % length].number
    })
}
