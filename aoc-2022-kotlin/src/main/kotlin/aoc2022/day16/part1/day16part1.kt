package aoc2022.day16.part1


val valves = createValves()
fun main() {
    val startPath = mutableListOf<Action>(valves.first { it.name == "AA" })
    repeat(30) {
        val newPaths = mutableListOf(mutableListOf<Action>())
        createRoutes(startPath.toMutableList(), startPath.size, newPaths)
        val bestPaths = newPaths.sortedBy { it.pressure() }.reversed().filter { it.size > 0 }
        val top20BestPaths = bestPaths.subList(0, minOf(20, bestPaths.size - 1))
        val extendedPaths = top20BestPaths.flatMap {
            val newNewPaths = mutableListOf(mutableListOf<Action>())
            createRoutes(it, it.size, newNewPaths)
            newNewPaths
        }.sortedBy { it.pressure() }.reversed().filter { it.size > 0 }
        if(extendedPaths.isEmpty()) {
            startPath.add(bestPaths.first()[startPath.size])
            return@repeat
        }
        val top20ExtendPaths = extendedPaths.subList(0, minOf(20, extendedPaths.size - 1))
        val furtherExtendedPaths = top20ExtendPaths.flatMap {
            val newNewPaths = mutableListOf(mutableListOf<Action>())
            createRoutes(it, it.size, newNewPaths)
            newNewPaths
        }.sortedBy { it.pressure() }.reversed().first()

        startPath.add(furtherExtendedPaths[startPath.size])
        println("Record: ${startPath.pressure()}")
    }
    println(startPath.map {
        when (it) {
            is TurnOn -> "On"; is Valve -> it.name; else -> ""
        }
    })
    println(startPath.pressure())
}

fun createRoutes(path: MutableList<Action>, startActions: Int, paths: MutableList<MutableList<Action>>) {
    if (path.size == startActions + 10 || path.size == 31) {
        paths.add(path)
        return
    }
    if (path.last() is TurnOn) {
        return (path[path.size - 2] as Valve).getConnections().forEach {
            createRoutes((path + it).toMutableList(), startActions, paths)
        }
    }
    if ((path.last() as Valve).flowRate > 0 && path.isNotTurnedOn((path.last() as Valve).name)) {
        createRoutes((path + TurnOn()).toMutableList(), startActions, paths)
    }
    (path.last() as Valve).getConnections().forEach {
        createRoutes((path + it).toMutableList(), startActions, paths)
    }
}

fun List<Action>.pressure(): Int {
    return this.mapIndexed { index, actions ->
        when (actions) {
            is TurnOn -> (this[index - 1] as Valve).flowRate * (30 - index)
            else -> 0
        }
    }.sum()
}

fun List<Action>.isNotTurnedOn(name: String): Boolean {
    return this.filterIndexed { index, action ->
        index < this.size - 1 && action is Valve && action.name == name && this[index + 1] is TurnOn
    }.isEmpty()
}

fun List<Action>.description(): String {
    return this.joinToString("") {
        when (it) {
            is Valve -> it.name; is TurnOn -> "On"; else -> ""
        }
    }
}

interface Action
class TurnOn : Action
data class Valve(
    val name: String,
    val flowRate: Int,
    private val connections: List<String>,
    var turnedOn: Boolean = false
) : Action {
    fun getConnections(): List<Valve> {
        return valves.filter { it.name in connections }
    }
}

fun createValves(): List<Valve> {
    return INPUT.split("\n").map {
        it.replace("Valve ", "").replace(" has flow rate", "")
            .replace(" tunnels lead to valves ", "")
            .replace(" tunnel leads to valve ", "")
            .split("=|;".toRegex())
    }.map {
        Valve(it[0], it[1].toInt(), it[2].split(",").map { it.trim() })
    }
}
