package aoc2022.day21.part2

data class Operation(val operation: ((Long, Long) -> Long), val symbol: String)
data class SecretNumber(val code: String, var value: Long? = null)
data class Monkey(
    val name: String,
    val secretNumber1: SecretNumber? = null,
    val secretNumber2: SecretNumber? = null,
    val operation: Operation? = null,
    var result: Long? = null
) {
    fun setSecretNumber(code: String, number: Long) {
        if (this.secretNumber1!!.code == code) this.secretNumber1.value = number
        if (this.secretNumber2!!.code == code) this.secretNumber2.value = number
        if (this.secretNumber1.value != null && this.secretNumber2.value != null) {
            this.result = operation!!.operation.invoke(this.secretNumber1.value!!, this.secretNumber2.value!!)
        }
    }

    fun getUnsolvedSecretCodes(): List<String> {
        return listOf(secretNumber1, secretNumber2).filter { it!!.value == null }.map { it!!.code }
    }

    fun printMonkey(monkeys: List<Monkey>): String {
        if (name == "humn") return "humn"
        if (result != null) {
            return result.toString()
        }
        val firstPart = if (secretNumber1!!.value != null) secretNumber1.value.toString() else
            "(" + monkeys.find { it.name == secretNumber1.code }!!.printMonkey(monkeys) + ")"
        val secondPart = if (name == "root") "=" else operation!!.symbol
        val lastPart = if (secretNumber2!!.value != null) secretNumber2.value.toString() else
            "(" + monkeys.find { it.name == secretNumber2.code }!!.printMonkey(monkeys) + ")"
        return firstPart + secondPart + lastPart
    }
}

private fun List<Monkey>.getSolvedNumbers(): List<SecretNumber> {
    return this.filter { it.result != null }.map { SecretNumber(it.name, it.result) }
}

fun main() {
    val monkeys = INPUT.split("\n").map { it.split(": | ".toRegex()) }
        .map {
            if (it[0] == "humn") Monkey("humn")
            else if (it.size == 2) Monkey(name = it[0], result = it[1].toLong())
            else Monkey(
                name = it[0],
                secretNumber1 = SecretNumber(it[1]),
                operation = when (it[2]) {
                    "+" -> Operation({ a: Long, b: Long -> a + b }, "+")
                    "*" -> Operation({ a: Long, b: Long -> a * b }, "*")
                    "/" -> Operation({ a: Long, b: Long -> a / b }, "/")
                    "-" -> Operation({ a: Long, b: Long -> a - b }, "-")
                    else -> throw Exception("")
                },
                secretNumber2 = SecretNumber(it[3])
            )
        }
    var unsolvedMonkeys = monkeys.filter { it.result == null }.toMutableList()
    do {
        var anyUpdated = false
        val solvedNumbers = monkeys.getSolvedNumbers()
        val solvedNumberCodes = solvedNumbers.map { it.code }
        unsolvedMonkeys.forEach { unsolvedMonkey ->
            if (unsolvedMonkey.name == "humn") return@forEach
            unsolvedMonkey.getUnsolvedSecretCodes().forEach { unsolvedCode ->
                if (solvedNumberCodes.contains(unsolvedCode)) {
                    val solved = solvedNumbers.find { it.code == unsolvedCode }!!
                    unsolvedMonkey.setSecretNumber(solved.code, solved.value!!)
                    anyUpdated = true
                }
            }
        }
        unsolvedMonkeys = monkeys.filter { it.result == null }.toMutableList()
    } while (anyUpdated)
    val rootMonkey = monkeys.find { it.name == "root" }!!
    println(rootMonkey.printMonkey(monkeys))
}
