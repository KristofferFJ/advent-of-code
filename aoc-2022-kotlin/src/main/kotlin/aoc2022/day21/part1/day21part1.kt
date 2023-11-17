package aoc2022.day21.part1

data class SecretNumber(val code: String, var value: Long? = null)
data class Monkey(
    val name: String,
    val secretNumber1: SecretNumber? = null,
    val secretNumber2: SecretNumber? = null,
    val operation: ((Long, Long) -> Long)? = null,
    var result: Long? = null
) {
    fun setSecretNumber(code: String, number: Long) {
        if (this.secretNumber1!!.code == code) this.secretNumber1.value = number
        if (this.secretNumber2!!.code == code) this.secretNumber2.value = number
        if (this.secretNumber1.value != null && this.secretNumber2.value != null) {
            this.result = operation!!.invoke(this.secretNumber1.value!!, this.secretNumber2.value!!)
        }
    }

    fun getUnsolvedSecretCodes(): List<String> {
        return listOf(secretNumber1, secretNumber2).filter { it!!.value == null }.map { it!!.code }
    }
}

private fun List<Monkey>.getSolvedNumbers(): List<SecretNumber> {
    return this.filter { it.result != null }.map { SecretNumber(it.name, it.result) }
}

fun main() {
    val monkeys = INPUT.split("\n").map { it.split(": | ".toRegex()) }
        .map {
            if (it.size == 2) Monkey(name = it[0], result = it[1].toLong())
            else Monkey(
                name = it[0],
                secretNumber1 = SecretNumber(it[1]),
                operation = when (it[2]) {
                    "+" -> { a: Long, b: Long -> a + b }
                    "*" -> { a: Long, b: Long -> a * b }
                    "/" -> { a: Long, b: Long -> a / b }
                    "-" -> { a: Long, b: Long -> a - b }
                    else -> throw Exception("")
                },
                secretNumber2 = SecretNumber(it[3])
            )
        }
    var unsolvedMonkeys = monkeys.filter { it.result == null }.toMutableList()
    while (unsolvedMonkeys.isNotEmpty()) {
        val solvedNumbers = monkeys.getSolvedNumbers()
        val solvedNumberCodes = solvedNumbers.map { it.code }
        unsolvedMonkeys.forEach { unsolvedMonkey ->
            unsolvedMonkey.getUnsolvedSecretCodes().forEach { unsolvedCode ->
                if (solvedNumberCodes.contains(unsolvedCode)) {
                    val solved = solvedNumbers.find { it.code == unsolvedCode }!!
                    unsolvedMonkey.setSecretNumber(solved.code, solved.value!!)
                }
            }
        }
        unsolvedMonkeys = monkeys.filter { it.result == null }.toMutableList()
    }
    val rootMonkey = monkeys.find { it.name == "root" }!!
    println(rootMonkey.result)
}
